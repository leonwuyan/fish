package wohuibao

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const (
	Parter    = "600202"
	SecretKey = "bfd8309c691e4bcf875119669ebaa5e9"
	API_URL   = "http://pay.wooopay.com/chargebank.aspx"
)

type PayType int

const (
	PAY_TYPE_ALIPAY PayType = 1006
	PAY_TYPE_WECHAT PayType = 1005
	PAY_TYPE_BANK   PayType = 2088
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
	signStr := "parter=" + params["parter"] + "&"
	signStr += "type=" + params["type"] + "&"
	signStr += "value=" + params["value"] + "&"
	signStr += "orderid=" + params["orderid"] + "&"
	signStr += "callbackurl=" + params["callbackurl"]
	signStr += SecretKey
	return create_md5(signStr)
}
func makeReceiveSign(params map[string]string) string {
	signStr := "orderid=" + params["orderid"] + "&"
	signStr += "opstate=" + params["opstate"] + "&"
	signStr += "ovalue=" + params["ovalue"]
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
	params["parter"] = Parter
	params["type"] = strconv.Itoa(int(payType))
	params["value"] = amount
	params["orderid"] = order
	params["callbackurl"] = notify
	params["sign"] = makeSendSign(params)
	return fmt.Sprintf("%s?%s", API_URL, genParams(params))
}
func NotifyResult(params map[string]string) string {
	sign := makeReceiveSign(params)
	if params["sign"] == sign {
		if params["opstate"] == "0" {
			return "SUCCESS"
		}
	}
	return ""
}
