package zongheng

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const (
	CustomerId = "16114"
	SecretKey  = "ece4a7bbb0c54c2e9e7a6a3be0afd999"
	API_URL    = "http://pay.zonghepay.com/bank/"
)

type PayType int

const (
	PAY_TYPE_ALIPAY PayType = 1006
	PAY_TYPE_WECHAT PayType = 1007
	PAY_TYPE_BANK   PayType = 1005
)

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
	params["parter"] = CustomerId
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
