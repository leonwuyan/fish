package sunapi

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	appId     = "syb524966717"
	secretKey = "f3144e9c485e4c2d9ff400f2de591861"
	apiUrl    = "https://wendan8.com/sdk/api/v1/trade/create"
)

type PayResult struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	TradeNo    string `json:"tradeNo"`
	OutTradeNo string `json:"outTradeNo"`
	Result     string `json:"result"`
}
type PayType string

const (
	PAY_TYPE_ALIPAY PayType = "2001"
	PAY_TYPE_WECHAT PayType = "1001"
)

type NotifyResult struct {
	Msg        string `json:"msg"`
	RealAmount string `json:"realAmount"`
	Amount     string `json:"amount"`
	Code       string `json:"code"`
	TradeNo    string `json:"tradeNo"`
	PayMethod  string `json:"payMethod"`
	OutTradeNo string `json:"outTradeNo"`
	AppId      string `json:"appId"`
	Sign       string `json:"sign"`
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
	sign_str = sign_str + "key=" + secretKey
	return create_md5(sign_str)
}
func create_md5(sign string) (s string) {
	data := []byte(sign)
	has := md5.Sum(data)
	s = fmt.Sprintf("%x", has)
	return strings.ToUpper(s)
}
func GetPayUrl(amount string, payType PayType, order, notify string) (result PayResult, err error) {
	params := make(map[string]string)
	params["appId"] = appId
	params["payMethod"] = string(payType)
	params["notifyUrl"] = notify
	params["returnUrl"] = notify
	params["outTradeNo"] = order
	params["signType"] = "MD5"
	params["amount"] = amount
	params["nonceStr"] = order
	params["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	params["sign"] = makeSign(params)
	return getPayUrlFromServer(params)
}
func getPayUrlFromServer(params map[string]string) (result PayResult, err error) {
	postData, _ := json.Marshal(params)
	postBody := bytes.NewBuffer(postData)
	response, err := http.Post(apiUrl, "application/json", postBody)
	if err != nil {
		logs.Error(err)
		return
	}
	req, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error(err)
	}
	err = json.Unmarshal(req, &result)
	if err != nil {
		logs.Error(err)
	}
	return
}
func Notify(params map[string]string) string {
	sign := makeSign(params)
	if params["sign"] == sign {
		if params["code"] == "10000" {
			return "success"
		}
	}
	return "fail"
}
