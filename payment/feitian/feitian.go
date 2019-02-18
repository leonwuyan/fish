package feitian

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const (
	tranCode   = "1101"       //交易码 N 4 固定 1101
	agtId      = "19020265"   //机构号 N 8
	merId      = "1902026573" //
	md5Key     = "Q9L6aSMwasACzmwSm4LVdOrH9xnWmYax"
	publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgHkH+XWGYeqVs7+ktU8Tfkvdv7qK9P9Tp6BH8zpNsQPvY7S0qZAX3RforjPYR+TEItuWVayV1vLiBNe46+zoIVr0yeHgmqpkZAp9m75+z9LlFsocnWVu6E80Zciv9qgfxwdFmAlpMaBAU9iv3DX8xHIOpVdLaSpYYaZcCCD4HzSSrZllD077suCk4bxn7olvUWwFPVAkSF2B/DoNIfmW+gfEDup4CJVuxbTeeN6dXvhdDO3yOy3rcRvhZ+/fuAhELOroeN42zxep0gZqHVQGwXK9xRnYuiNydPal8Yyc4qUsQLylXy3eYjFGlAlYehcMajNGclgSVNOTyrx5cfA9eQIDAQAB"
	privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCzfohc73Eckzed8x9sxIEAY56svdgMfA+JyCJ+PU+UhsGjNdfVET1Y57Gm/dH4i6hXufXPFIfn4x1x75l03Js2yfOr16J/4tyz5MFvw7LQyzko5FFHNIIpGllQE8a2xkaE8A7EOgdU1xTS90+GDq+gh4j2uFsukzSAcB+evv+NID1dWm/tgI9UlUabPxvlOLs9ls8dDLf+BFF24P7NSou+4X2HoUwm8sZkY7+Pi1pNvSDCDKOee+98l+ZIsRMOrxmsxn2FGJNQ/6z/nzOK7YxJeK8fk5QqDScVZM8bWS5PLimI8mSdG8aqCIA4UjxkhZdwvuXezYq2gGtQqppc0cQpAgMBAAECggEACnWmJpxBYCPoCsk/HIx64/XqKjDJ5ThZg3LpJ6gU+he8MvGRSTY+9+5QeOWzlaM1xR7uBwowPUAL2DbZH2k3lkLAc1QWgRNMPyNmbkfeJLKi+xIK9nZF23VQnQz5+G6nb1m50o97G+OIsB5/QNWcQGOFhLMLhTRIy+88uHhy1dthGB7M5CMmC0or0MfYAxiZu/DciYjxt4MwD7N7ZxXx //rP5d3gY7zcOuL4Y1kvdvm0Kr65OopFx0F8O3ZVibvFPHN0L7aiwieRqK5pzNLtq6gVrTssGHbgPx6hJLwkFpWr9iDEgbmdUgINdQJ+eqcO7OP3HcvUiLUwllsaiMp2nQKBgQDeWZ94jA71ZFGO0XQB0siUpEe3zx5gS4WXGO8iJsyVuE4p1mSv8Cpf7mdmFfSrK4S8nyM2kMrTXgweeGoSH45kYvZH/TfyXOT4EZT+ED/C0Ywyeonrx33SesdoCrElaLhqy6F331r+2UhBVQ77mwTlXjkdam3JUCCylYucb+pcrwKBgQDOqJJ6O9Xe9BIsXMDK6OdHvnj0a6lZNVmv6s/vplWrFmhvEDztnkPx1wlV5pz/bcFlyDkhWvKUecK3DwyvSHZg8hZqz51uBbZNro32yYc1sB0ynmlb/c+RVh18SkkxzxTYbTrC4SNzT9It7JVnL/Xl5HpjDJwyhBaRjAR0PwUSpwKBgBfat6AgLo6nF5Th5bc4XOqNgWA36UNddtSSCT25ueMqJcCZTq4Nbw5hrlgmbNxcS95MmGOHPytUAKrYMlEFb6pXb6KjiPgIfUwb3scvCRgmkQrtWJSiD16ga3o/A2hHBtn2RLlujf2fZAAWVwgpRmoxJKGCRIr/fLKMFrFO1G7jAoGABrqqelx0bFwaM2OZy4Gl9koXei3/R65bC0VkG/OYmoeSQvuOYFZk8/0Cis+FkTOrtnq8kX96oqcMVhWhXhvH7wQzAqtb/vckpobmjICnbHIdoUZTO/GZzqu1LhTlKUAK97kmzq2yD9ErkN9BRijjoua7rXn9pA91BKxAgOCEc8UCgYA56/N/BkHRu8+KALhW6kvC4oqCZKP+dDqcpFjdTGKVgrE178p8pgJenAt8IpqV0Vi59lxj5S4End26SbLsCx+YeR5X2MYpPQu84y+d6l510NqFJ2VIWUoMUkC9SV2zMzUP1HJ+q5S5ZETo459iLspbm8WVCtK5+j95TFDTB6xL5g=="
	apiUrl     = "http://pay.feitian2018.com:8343/webwt/pay/gateway.do"
	termIp     = "47.75.129.72"
)

type PayType string

const (
	PAY_TYPE_ALIPAY PayType = "ALIWAPPAY"
	PAY_TYPE_WECHAT PayType = "WXPAY"
	PAY_TYPE_BANK   PayType = "CPPAY"
)

type ReqHead struct {
	Sign string `json:"sign"`
}
type PostData struct {
	REQ_HEAD interface{} `json:"REQ_HEAD"`
	REQ_BODY interface{} `json:"REQ_BODY"`
}
type ReceiveData struct {
	REP_HEAD interface{} `json:"REP_HEAD"`
	REP_BODY interface{} `json:"REP_BODY"`
}

func genParams(params map[string]string) (paramsStr string) {
	for k, _ := range params {
		paramsStr += k + "=" + params[k] + "&"
	}
	return paramsStr[0 : len(paramsStr)-1]
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
	sign_str = sign_str + "key=" + md5Key
	return create_md5(sign_str)
}
func formatKey(raw, prefix, suffix string) (result []byte) {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var ll = 64
	var sl = len(raw)
	var c = sl / ll
	if sl%ll > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * ll
		var e = b + ll
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}
func rsaEncrypt(origData []byte) (out []byte, err error) {
	var block *pem.Block
	pk := formatKey(privateKey, "-----BEGIN RSA PRIVATE KEY-----", "-----END RSA PRIVATE KEY-----")
	block, _ = pem.Decode([]byte(pk))
	if block == nil {
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	h2 := sha256.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), crypto.SHA256, hashed) //签名
	return signature2, err
}
func RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
}
func create_md5(sign string) (s string) {
	data := []byte(sign)
	has := md5.Sum(data)
	s = fmt.Sprintf("%x", has)
	return strings.ToUpper(s)
}

func GetApiUrl(amount string, payType PayType, order, notify string) (result ReceiveData, err error) {
	params := make(map[string]string)
	params["tranCode"] = tranCode
	params["agtId"] = agtId
	params["merId"] = merId
	params["orderAmt"] = amount
	params["orderId"] = order
	params["goodsName"] = fmt.Sprintf("%x", "金币")
	params["notifyUrl"] = notify
	params["nonceStr"] = order
	params["stlType"] = "T1"
	params["termIp"] = termIp
	params["payChannel"] = string(payType)
	sign := makeSign(params)
	params["sign"] = sign
	//rsaData, _ := RsaEncrypt(sign)
	b, _ := rsaEncrypt([]byte(sign))
	postData := PostData{
		REQ_HEAD: ReqHead{base64.StdEncoding.EncodeToString(b)},
		REQ_BODY: params,
	}
	return getPayUrlFromServer(postData)
}
func NotifyResult(rsaSign string, params map[string]string) string {
	sign := makeSign(params)
	if RsaVerySignWithSha1Base64(sign, rsaSign, publicKey) == nil {
		if params["orderState"] == "01" {
			return "SUCCESS"
		}
	}
	return ""
}

func getPayUrlFromServer(data PostData) (result ReceiveData, err error) {
	postData, _ := json.Marshal(data)
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
