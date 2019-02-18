package controllers

import (
	"bytes"
	"encoding/json"
	"fish/configs"
	"fish/enums"
	"fish/managers"
	"fish/payment"
	"fish/payment/feitian"
	"fish/payment/hongjia"
	"fish/payment/huiyi"
	"fish/payment/sunapi"
	"fish/payment/wohuibao"
	"fish/payment/yijia"
	"fish/payment/zongheng"
	"fmt"
	"github.com/astaxie/beego/logs"
	"image/jpeg"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type MainController struct {
	baseController
}

func (c *MainController) Prepare() {
	c.baseController.Prepare()
	managers.SystemInstance.PageVisitor(c.Ctx.Input, nil)
}
func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Advertise() {
	temp_id, _ := c.GetInt("template_id")
	agentId, _ := c.GetInt("agentId")
	switch temp_id {
	default:
		c.TplName = "advertise.tpl"
		break
	case 1:
		c.TplName = "advertise.tpl"
		break
	}
	c.Data["agentId"] = agentId
}

func (c *MainController) Recharge() {
	userId, _ := c.GetInt("userId")
	channel, _ := c.GetInt("channel")
	amount, _ := c.GetFloat("pay_amt")
	payType, _ := c.GetInt("pay_type")
	if configs.PayAmountFloor {
		amount = math.Floor(amount)
	}
	//PaymentAction := make(map[enums.PaymentChannel]func(userId, channel, pay_type int, amount float64))
	//PaymentAction[enums.PAY_CHANNEL_HUIYI] = c.Recharge_Hui_Yi
	//PaymentAction[enums.PAY_CHANNEL_WOHUIBAO] = c.Recharge_Wo_Hui_Bao
	//PaymentAction[enums.PAY_CHANNEL_HONGJIA] = c.Recharge_Hong_Jia
	//PaymentAction[enums.PAY_CHANNEL_ZONGHENG] = c.Recharge_Zong_Heng
	//PaymentAction[enums.PAY_CHANNEL_YIJIA] = c.Recharge_Yi_Jia
	//PaymentAction[enums.PAY_CHANNEL_SUNAPI] = c.Recharge_Sun_Api
	var configChannel enums.PaymentChannel
	switch payType {
	//支付宝
	case 22:
		i_channel, _ := strconv.Atoi(configs.PaymentGate["alipay"])
		configChannel = enums.PaymentChannel(i_channel)
		break
		//网银
	case 23:
		i_channel, _ := strconv.Atoi(configs.PaymentGate["bank"])
		configChannel = enums.PaymentChannel(i_channel)
		//微信
	case 30:
		i_channel, _ := strconv.Atoi(configs.PaymentGate["wechat"])
		configChannel = enums.PaymentChannel(i_channel)
	}
	c.RechargeHandler(configChannel, userId, channel, payType, amount)
	//PaymentAction[configChannel](userId, channel, pay_type, amount)
}
func (c *MainController) RechargeHandler(payChannel enums.PaymentChannel, userId, channel, payType int, amount float64) {
	pay_order := payment.Create_order()
	err := managers.SystemInstance.PreRecharge(payChannel, userId, channel, payType, int(amount*100), pay_order)
	if err != nil {
		c.Data["json"] = c.jsonData(enums.DB_ACTION_ERROR)
		c.Abort("10001")
	}
	switch payChannel {
	case enums.PAY_CHANNEL_HUIYI:
		c.Recharge_Hui_Yi(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_WOHUIBAO:
		c.Recharge_Wo_Hui_Bao(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_HONGJIA:
		c.Recharge_Hong_Jia(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_ZONGHENG:
		c.Recharge_Zong_Heng(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_YIJIA:
		c.Recharge_Yi_Jia(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_ALIPAY:
		//c.Recharge_Hui_Yi(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_SUNAPI:
		c.Recharge_Sun_Api(userId, channel, payType, amount, pay_order)
		break
	case enums.PAY_CHANNEL_FEITIAN:
		c.Recharge_Fei_Tian(userId, channel, payType, amount, pay_order)
		break
	}
}
func (c *MainController) Recharge_Hui_Yi(userId, channel, payType int, amount float64, payOrder string) {
	c.TplName = "payPost.tpl"
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		c.Data["html"] = huiyi.PostPay(int(amount*100), huiYiType(payType), payOrder, configs.Domain["domain"]+"notify/hui_yi")
		break
	}
}
func (c *MainController) Recharge_Wo_Hui_Bao(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		url := wohuibao.GetApiUrl(strconv.FormatFloat(amount, 'f', 2, 64), woHuiBaoType(payType), payOrder, configs.Domain["domain"]+"notify/wo_hui_bao")
		c.Redirect(url, 302)
		break
	}
}
func (c *MainController) Recharge_Hong_Jia(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		url := hongjia.GetApiUrl(strconv.FormatFloat(amount, 'f', 2, 64), hongJiaType(payType), payOrder, configs.Domain["domain"]+"notify/hong_jia")
		//println(url)
		c.Redirect(url, 302)
		break
	}
}
func (c *MainController) Recharge_Zong_Heng(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		url := zongheng.GetApiUrl(strconv.FormatFloat(amount, 'f', 2, 64), zongHengType(payType), payOrder, configs.Domain["domain"]+"notify/zong_heng")
		//println(url)
		c.Redirect(url, 302)
		break
	}
}
func (c *MainController) Recharge_Yi_Jia(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		c.Data["html"] = yijia.PostPay(amount, yiJiaType(payType), payOrder, configs.Domain["domain"]+"notify/sun_api")
		break
	}
}
func (c *MainController) Recharge_Sun_Api(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		result, err := sunapi.GetPayUrl(strconv.FormatFloat(amount, 'f', 2, 64), sunApiType(payType), payOrder, configs.Domain["domain"]+"notify/sun_api")
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("发生错误：%s", err.Error()))
		}
		if result.Code == "10000" {
			c.Redirect(result.Result, 302)
		} else {
			c.Ctx.WriteString(result.Msg)
		}
		break
	}
}
func (c *MainController) Recharge_Fei_Tian(userId, channel, payType int, amount float64, payOrder string) {
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		result, err := feitian.GetApiUrl(strconv.Itoa(int(amount*100)), feiTianType(payType), payOrder, configs.Domain["domain"]+"notify/fei_tian")
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("发生错误：%s", err.Error()))
		}
		//c.Ctx.WriteString(fmt.Sprintf("%v", result))
		resultMap := result.REP_BODY.(map[string]interface{})
		if resultMap["rspcode"].(string) == "000000" {
			c.Redirect(resultMap["codeUrl"].(string), 302)
		} else {
			var dst []byte
			fmt.Sscanf(resultMap["rspmsg"].(string), "%X", &dst)
			c.Ctx.WriteString(string(dst))
		}
		break
	}
}

func (c *MainController) Notify_Hui_Yi() {
	params := make(map[string]string)
	params["mer_id"] = c.GetString("mer_id")
	params["trade_no"] = c.GetString("trade_no")
	params["mer_trade_no"] = c.GetString("mer_trade_no")
	params["trade_amt"] = c.GetString("trade_amt")
	params["timestamp"] = c.GetString("timestamp")
	params["code"] = c.GetString("code")
	params["sign"] = c.GetString("sign")
	params["sign_method"] = c.GetString("sign_method")
	params["extend"] = c.GetString("extend")
	result := huiyi.NotifyResult(params)
	logs.Info(params)
	if result == "SUCCESS" {
		if err := managers.SystemInstance.FinishRecharge(params["trade_no"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		}
	}
	c.Ctx.WriteString(result)
}
func (c *MainController) Notify_Wo_Hui_Bao() {
	params := make(map[string]string)
	params["orderid"] = c.GetString("orderid")
	params["opstate"] = c.GetString("opstate")
	params["ovalue"] = c.GetString("ovalue")
	params["sign"] = c.GetString("sign")
	params["sysorderid"] = c.GetString("sysorderid")
	params["systime"] = c.GetString("systime")
	params["attach"] = c.GetString("attach")
	params["msg"] = c.GetString("msg")
	result := wohuibao.NotifyResult(params)
	logs.Info(params)
	if result == "SUCCESS" {
		if err := managers.SystemInstance.FinishRecharge(params["orderid"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		}
	}
	c.Ctx.WriteString(result)
}
func (c *MainController) Notify_Hong_Jia() {
	params := make(map[string]string)
	params["status"] = c.GetString("status")
	params["customerid"] = c.GetString("customerid")
	params["sdpayno"] = c.GetString("sdpayno")
	params["sdorderno"] = c.GetString("sdorderno")
	params["total_fee"] = c.GetString("total_fee")
	params["paytype"] = c.GetString("paytype")
	params["remark"] = c.GetString("remark")
	params["sign"] = c.GetString("sign")
	logs.Info(params)
	result := hongjia.NotifyResult(params)
	if result == "SUCCESS" {
		if err := managers.SystemInstance.FinishRecharge(params["sdorderno"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		}
	}
	c.Ctx.WriteString(strings.ToLower(result))
}
func (c *MainController) Notify_Zong_Heng() {
	params := make(map[string]string)
	params["orderid"] = c.GetString("orderid")
	params["opstate"] = c.GetString("opstate")
	params["ovalue"] = c.GetString("ovalue")
	params["sign"] = c.GetString("sign")
	logs.Info(params)
	result := hongjia.NotifyResult(params)
	if result == "SUCCESS" {
		if err := managers.SystemInstance.FinishRecharge(params["sdorderno"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		}
	}
	c.Ctx.WriteString(strings.ToLower(result))
}
func (c *MainController) Notify_Yi_Jia() {
	params := make(map[string]string)
	params["memberid"] = c.GetString("memberid")
	params["orderid"] = c.GetString("orderid")
	params["amount"] = c.GetString("amount")
	params["transaction_id"] = c.GetString("transaction_id")
	params["datetime"] = c.GetString("datetime")
	params["returncode"] = c.GetString("returncode")
	params["sign"] = c.GetString("sign")
	logs.Info(params)
	result := yijia.NotifyResult(params)
	if result == "SUCCESS" {
		if err := managers.SystemInstance.FinishRecharge(params["orderid"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		} else {
			result = "OK"
		}
	}
	c.Ctx.WriteString(result)
}
func (c *MainController) Notify_Sun_Api() {
	jsonData, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	params := make(map[string]string)
	err := json.Unmarshal(jsonData, &params)
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("json parse fail")
		return
	}
	logs.Info(params)
	result := sunapi.Notify(params)
	if result == "success" {
		managers.SystemInstance.FinishRecharge(params["outTradeNo"])
	}
	c.Ctx.WriteString(result)
}
func (c *MainController) Notify_Fei_Tian() {
	jsonData, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	//jsonData := []byte("{\"REP_HEAD\":{\"sign\":\"VCV0IlJLYxixKNmHnMfX9LkxqVBdElYE0cFOmZLJBzYgG14ra7h4EyeeDf9MYX5RFwmfNl37NbOVRuGIgk1uvDEhdM/J3HZZO0sBiDVfi/b24QIUuVeEadoscZVFiyHEPFQSWFecAx/e1cbzYAQj0OBsg094kdhWaNJD1dDcAFNtLf0sP6DbPhfH2v009dT/kHwkCxjAmPweCxHQ2DtHSOcdyrq8NfkFVVPxe1zwegShhVfhMzNxTAHNGOTqni8zdyRUvIdn0AhBPwfvAZcYYKr9HOqaRPInogb8MUZ1NS1Jw9v69LK/9JPShhV/PQtPtSfgmMvxmaZ9TtDd9pSEWA==\"},\"REP_BODY\":{\"orderState\":\"01\",\"tranSeqId\":\"190202651219021814434563579O0Obe\",\"payTime\":\"20190218144859\",\"orderAmt\":\"1000\",\"orderId\":\"20190218144345540\"}}")
	repData := make(map[string]interface{})
	err := json.Unmarshal(jsonData, &repData)
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("json parse fail")
		return
	}
	headData := repData["REP_HEAD"].(map[string]interface{})
	bodyData := repData["REP_BODY"].(map[string]interface{})
	params := make(map[string]string)
	for k, v := range bodyData {
		params[k] = v.(string)
	}
	result := feitian.NotifyResult(headData["sign"].(string), params)
	if result == "SUCCESS" {
		managers.SystemInstance.FinishRecharge(params["orderId"])
	}
	c.Ctx.WriteString(result)
}

func (c *MainController) GetPostForms() {
	c.Layout = ""
	c.TplName = "modal.tpl"
	action := c.GetString("a")
	var formParams []map[string]interface{}
	formParams = append(formParams, map[string]interface{}{"id": "url", "value": c.Ctx.Request.Referer()})
	switch action {
	case "add_agent":
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "添加代理"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "add_agent", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "agent_name", "name": "用 户 名", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "default_pass", "name": "设置密码", "value": "", "type": "password", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "mobile", "name": "手机号码", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "nickname", "name": "代理昵称", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "rate", "name": "返现比例", "value": "0", "type": "slider", "min": "0", "max": "70", "step": "1", "can-min": "0"})
		break
	case "add_agent_child":
		max_rate := c.GetString("rate")
		agent_id := c.GetString("id")
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "添加代理"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "add_agent_child", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "agent_id", "value": agent_id, "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "agent_name", "name": "用 户 名", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "default_pass", "name": "设置密码", "value": "", "type": "password", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "mobile", "name": "手机号码", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "nickname", "name": "代理昵称", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "rate", "name": "返现比例", "value": "0", "type": "slider", "min": "0", "max": max_rate, "step": "1", "can-min": "0"})
		break
	case "change_rate":
		max_rate := c.GetString("rate")
		agent_id := c.GetString("id")
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "修改提成"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "change_rate", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "agent_id", "value": agent_id, "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "rate", "name": "返现比例", "value": "0", "type": "slider", "min": "0", "max": max_rate, "step": "1", "can-min": "0"})
		break
	case "add_admin":
		var values = configs.AdminPower
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "添加管理员"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "add_admin", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_name", "name": "用 户 名", "value": "", "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_pass", "name": "设置密码", "value": "", "type": "password", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_powers", "name": "权限设置", "value": values, "checked": c.GetString("p"), "type": "checkbox", "required": "true"})
		break
	case "edit_admin":
		id, _ := c.GetInt("id")
		var values = configs.AdminPower
		admin, _ := managers.AdminInstance.GetAdminById(id)
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "编辑管理员"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "edit_admin", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_id", "value": c.GetString("id"), "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_name", "name": "用 户 名", "value": admin.Name, "type": "text", "required": "true", "readonly": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "admin_powers", "name": "权限设置", "value": values, "checked": admin.Permissions, "type": "checkbox", "required": "true"})
		break
	case "add_service_msg":
		formParams = append(formParams, map[string]interface{}{"id": "title", "value": "回复消息"})
		formParams = append(formParams, map[string]interface{}{"id": "action", "value": "add_service_msg", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "msg_id", "value": c.GetString("id"), "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "msg_type", "value": "1", "type": "hidden"})
		formParams = append(formParams, map[string]interface{}{"id": "user_id", "name": "用户ID", "value": c.GetString("uid"), "type": "text", "required": "true"})
		formParams = append(formParams, map[string]interface{}{"id": "msg_content", "name": "消息内容", "value": "", "type": "text", "required": "true"})
		break
	default:
		break
	}
	createPostForm(&c.baseController, formParams)
	return
}
func (c *MainController) GeneralizeQr() {
	agentId := c.Ctx.Input.Param(":agent")
	templateId := c.Ctx.Input.Param(":id")
	adUrl := getDownUrl() + "?id=" + agentId
	img, _ := createQr(fmt.Sprintf("static/img/bg%s.png", templateId), adUrl)
	//img, _ := createQr("static/img/bg.png", adUrl)
	var b bytes.Buffer
	jpeg.Encode(&b, img, &jpeg.Options{Quality: 75})
	c.Ctx.Output.ContentType("jpeg")
	c.Ctx.Output.Body(b.Bytes())
	//encoder := png.Encoder{CompressionLevel: png.BestCompression}
	//var b bytes.Buffer
	//encoder.Encode(&b, img)
	//c.Ctx.Output.ContentType("png")
	//c.Ctx.Output.Body(b.Bytes())
}

func huiYiType(payType int) huiyi.PayType {
	switch payType {
	case 22:
		return huiyi.ALIPAY_H5
	case 30:
		return huiyi.WECHAT_H5
	case 23:
		return huiyi.BANK_SC
	default:
		return huiyi.ALIPAY_H5
	}
}
func woHuiBaoType(payType int) wohuibao.PayType {
	switch payType {
	case 22:
		return wohuibao.PAY_TYPE_ALIPAY
	case 30:
		return wohuibao.PAY_TYPE_WECHAT
	case 23:
		return wohuibao.PAY_TYPE_BANK
	default:
		return wohuibao.PAY_TYPE_ALIPAY
	}
}
func hongJiaType(payType int) hongjia.PayType {
	switch payType {
	case 22:
		return hongjia.PAY_TYPE_ALIPAY
	case 30:
		return hongjia.PAY_TYPE_WECHAT
	case 23:
		return hongjia.PAY_TYPE_BANK
	default:
		return hongjia.PAY_TYPE_ALIPAY
	}
}
func zongHengType(payType int) zongheng.PayType {
	switch payType {
	case 22:
		return zongheng.PAY_TYPE_ALIPAY
	case 30:
		return zongheng.PAY_TYPE_WECHAT
	case 23:
		return zongheng.PAY_TYPE_BANK
	default:
		return zongheng.PAY_TYPE_ALIPAY
	}
}
func yiJiaType(payType int) yijia.PayType {
	switch payType {
	case 22:
		return yijia.PAY_TYPE_ALIPAY
	case 30:
		return yijia.PAY_TYPE_WECHAT
		//case 23:
		//	return yijia.PAY_TYPE_BANK
	default:
		return yijia.PAY_TYPE_ALIPAY
	}
}
func sunApiType(payType int) sunapi.PayType {
	switch payType {
	case 22:
		return sunapi.PAY_TYPE_ALIPAY
	case 30:
		return sunapi.PAY_TYPE_WECHAT
		//case 23:
		//	return yijia.PAY_TYPE_BANK
	default:
		return sunapi.PAY_TYPE_ALIPAY
	}
}
func feiTianType(payType int) feitian.PayType {
	switch payType {
	case 22:
		return feitian.PAY_TYPE_ALIPAY
	case 30:
		return feitian.PAY_TYPE_WECHAT
	case 23:
		return feitian.PAY_TYPE_BANK
	default:
		return feitian.PAY_TYPE_ALIPAY
	}
}
