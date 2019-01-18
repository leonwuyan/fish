package hongjia

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const (
	CustomerId = "12656"
	SecretKey  = "905e4b2709f4fa2ac29bf0f3dffec862ff180646"
	API_URL    = "http://www.hongjiapay.com/gateway"
)

type PayType string

const (
	PAY_TYPE_ALIPAY PayType = "alipaywap"
	PAY_TYPE_WECHAT PayType = "weixin"
	PAY_TYPE_BANK   PayType = "quickbank"
)

type ReceiveData struct {
	OrderId    string `json:"orderid"`
	OpState    string `json:"opstate"`
	OValue     string `json:"ovalue"`
	Sign       string `json:"sign"`
	SysOrderId string `json:"sysorderid"`
	SysTime    string `json:"systime"`
	Attach     string `json:"attach"`
	Msg        string `json:"msg"`
}

func genParams(params map[string]string) (paramsStr string) {
	for k, _ := range params {
		paramsStr += k + "=" + params[k] + "&"
	}
	return paramsStr[0 : len(paramsStr)-1]
}
func makeSendSign(params map[string]string) string {
	signStr := "version=" + params["version"] + "&"
	signStr += "customerid=" + params["customerid"] + "&"
	signStr += "total_fee=" + params["total_fee"] + "&"
	signStr += "sdorderno=" + params["sdorderno"] + "&"
	signStr += "notifyurl=" + params["notifyurl"] + "&"
	signStr += "returnurl=" + params["returnurl"] + "&"
	signStr += SecretKey
	return create_md5(signStr)
}
func makeReceiveSign(params map[string]string) string {
	signStr := "customerid=" + params["customerid"] + "&"
	signStr += "status=" + params["status"] + "&"
	signStr += "sdpayno=" + params["sdpayno"] + "&"
	signStr += "sdorderno=" + params["sdorderno"] + "&"
	signStr += "total_fee=" + params["total_fee"] + "&"
	signStr += "paytype=" + params["paytype"] + "&"
	signStr += SecretKey
	return create_md5(signStr)
}
func create_md5(sign string) (s string) {
	data := []byte(sign)
	has := md5.Sum(data)
	s = fmt.Sprintf("%x", has)
	return strings.ToLower(s)
}
func GetApiUrl(amount string, payType PayType, order, notify string) string {
	params := make(map[string]string)
	params["version"] = "1.0"
	params["customerid"] = CustomerId
	params["sdorderno"] = order
	params["total_fee"] = amount
	params["paytype"] = string(payType)
	params["notifyurl"] = notify
	params["returnurl"] = notify
	params["sign"] = makeSendSign(params)
	return fmt.Sprintf("%s?%s", API_URL, genParams(params))
}
func NotifyResult(params map[string]string) string {
	sign := makeReceiveSign(params)
	if params["sign"] == sign {
		if params["status"] == "1" {
			return "SUCCESS"
		}
	}
	return ""
}
