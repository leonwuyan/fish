package yijia

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	CustomerId = "190204646"
	SecretKey  = "ece4a7bbb0c54c2e9e7a6a3be0afd999"
	API_URL    = "http://www.yjiapay.com/Pay_Index.html"
)

type PayType int

const (
	PAY_TYPE_ALIPAY PayType = 904
	PAY_TYPE_WECHAT PayType = 917
	//PAY_TYPE_BANK   PayType = 1005
)

func createParams(amount int, payType PayType, order, notify string) (params map[string]string) {
	params = make(map[string]string)
	params["pay_memberid"] = CustomerId
	params["pay_orderid"] = order
	params["pay_applydate"] = time.Now().Format("2006-01-02 15:04:05")
	params["pay_bankcode"] = strconv.Itoa(int(payType))
	params["pay_notifyurl"] = notify
	params["pay_callbackurl"] = notify
	params["pay_amount"] = strconv.Itoa(amount)
	params["pay_md5sign"] = makeSign(params)
	params["pay_productname"] = "金币"
	return params
}
func makeSign(params map[string]string) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var sign_str string
	for _, key := range keys {
		if !(params[key] == "" || key == "sign") {
			sign_str += key + "=" + params[key] + "&"
		}
	}
	sign_str = sign_str + "key=" + SecretKey
	println(sign_str)
	return strings.ToUpper(create_md5(sign_str))
}
func create_md5(sign string) (s string) {
	data := []byte(sign)
	has := md5.Sum(data)
	s = fmt.Sprintf("%x", has)
	return strings.ToLower(s)
}
func PostPay(amount int, payType PayType, order, notify string) (form string) {
	params := createParams(amount, payType, order, notify)
	form = "<form action='" + API_URL + "' method='post' id='form_pay'>"
	for key, value := range params {
		form += "<input type='hidden' name='" + key + "' value='" + value + "'>"
	}
	form += "</form>"
	form += "<script>document.getElementById('form_pay').submit()</script>"
	return form
}
func NotifyResult(params map[string]string) string {
	sign := makeSign(params)
	if params["sign"] == sign {
		if params["returncode"] == "00" {
			return "SUCCESS"
		}
	}
	return ""
}
