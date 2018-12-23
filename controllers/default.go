package controllers

import (
	"fish/configs"
	"fish/enums"
	"fish/managers"
	"fish/payment"
	"fish/payment/huiyi"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	baseController
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
	//userId=167&channel=1&pay_type=22&pay_amt=68.14&deviceid=5878a7ab84fb43402106c575658472fa
	c.TplName = "payPost.tpl"
	userId, _ := c.GetInt("userId")
	channel, _ := c.GetInt("channel")
	amount, _ := c.GetFloat("pay_amt")
	pay_type, _ := c.GetInt("pay_type")
	//deviceid := c.GetString("deviceid")
	pay_order := payment.Create_order()
	err := managers.SystemInstanse.PreRecharge(userId, channel, pay_type, int(amount*100), pay_order)
	if err != nil {
		c.Data["json"] = c.jsonData(enums.DB_ACTION_ERROR)
		c.Abort("10001")
	}
	//var action interface{}
	switch channel {
	default:
		c.Abort("10002")
		break
	case 105:
		action := huiyi.PostPay
		c.Data["html"] = action(int(amount*100), huiYiType(pay_type), pay_order, configs.Domain["domain"]+"notify_hui_yi")
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
	logs.Info(c.Ctx.Request.URL)
	if result == "SUCCESS" {
		if err := managers.SystemInstanse.FinishRecharge(params["trade_no"]); err != nil {
			logs.Error(err)
			result = "Recharge fail"
		}
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
	default:
		break
	}
	createPostForm(&c.baseController, formParams)
	return
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
