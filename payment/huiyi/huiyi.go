package huiyi

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var merId = "1067"
var userId = "1067"
var secretKey = "8d1e408ee8575afd28349b5ae3cdea0a"
var currency = "CNY"
var http_pay = "https://gateway.66nou.com/v2/payment/trade"
//var http_query = "https://gateway.66nou.com/v2/payment/query"
var successCode = "000000"

type NotifyData struct {
	MerId      string `json:"mer_id"`
	TradeNo    string `json:"trade_no"`
	MerTradeNo string `json:"mer_trade_no"`
	TradeAmt   int    `json:"trade_amt"`
	Timestamp  string `json:"timestamp"`
	Code       string `json:"code"`
	Sign       string `json:"sign"`
	SignMethod string `json:"sign_method"`
	Extend     string `json:"extend"`
}
type PayType int

const (
	BANK_NET  PayType = 1
	BANK_SC   PayType = 2
	ALIPAY_QR PayType = 3
	ALIPAY_H5 PayType = 4
	WECHAT_QR PayType = 5
	WECHAT_H5 PayType = 6
)

func createParams(amount int, payType PayType, order, notify string) (params map[string]string) {
	params = make(map[string]string)
	params["mer_id"] = merId
	params["user_id"] = userId
	params["trade_amt"] = strconv.Itoa(amount)
	params["mer_trade_no"] = order
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["notify_url"] = notify
	//params["front_url"] = callback
	params["sign_method"] = "MD5"
	params["pay_type"] = strconv.Itoa(int(payType))
	params["product_name"] = ""
	params["product_desc"] = ""
	params["extend"] = ""
	params["currency"] = currency
	params["sign"] = makeSign(params, secretKey)
	return params
}
func makeSign(params map[string]string, secretKey string) string {
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
	sign_str = sign_str + "key=" + secretKey
	println(sign_str)
	return create_md5(sign_str)
}
func create_md5(sign string) (s string) {
	data := []byte(sign)
	has := md5.Sum(data)
	s = fmt.Sprintf("%x", has)
	return strings.ToUpper(s)
}
func PostPay(amount int, payType PayType, order, notify string) (form string) {
	params := createParams(amount, payType, order, notify)
	form = "<form action='" + http_pay + "' method='post' id='form_pay'>"
	for key, value := range params {
		form += "<input type='hidden' name='" + key + "' value='" + value + "'>"
	}
	form += "</form>"
	form += "<script>document.getElementById('form_pay').submit()</script>";
	return form
}
func NotifyResult(params map[string]string) string {
	sign := makeSign(params, secretKey)
	if params["sign"] == sign {
		if params["code"] == successCode {
			return "SUCCESS"
		}
	}
	return ""
}
