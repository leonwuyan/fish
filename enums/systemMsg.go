package enums

import (
)

var SystemMsg map[ReturnCode]string
var CashStateMsg map[CashState]string
var CashTypeMsg map[CashType]string

func init() {
	SystemMsg = make(map[ReturnCode]string)
	SystemMsg[SUCCESS] = "成功"
	SystemMsg[QUERY_DATA_ERROR] = "查询数据失败"
	SystemMsg[DB_ACTION_ERROR] = "数据库存储失败"
	SystemMsg[CAPTCHA_ERROR] = "验证码错误"
	SystemMsg[AGENT_NOT_LOGIN] = "没有登录"
	SystemMsg[AGENT_NOT_FOUND] = "账号不存在"
	SystemMsg[AGENT_UNUSED] = "不可用"
	SystemMsg[AGENT_PASSWORD_ERROR] = "密码错误"
	SystemMsg[AGENT_OLDPASSWORD_ERROR] = "旧密码错误"
	SystemMsg[AGENT_NEWPASSWORD_ERROR] = "新密码不一致"
	SystemMsg[PLAYER_NOT_FOUND] = "账号不存在"

	CashStateMsg = make(map[CashState]string)
	CashStateMsg[CASH_STATE_APPLY] = "申请中"
	CashStateMsg[CASH_STATE_WAIT] = "等待付款"
	CashStateMsg[CASH_STATE_REFUSED] = "拒绝兑换"
	CashStateMsg[CASH_STATE_SUCCESS] = "成功兑换"
	CashStateMsg[CASH_STATE_REFUND] = "退款"

	CashTypeMsg = make(map[CashType]string)
	CashTypeMsg[CASH_TYPE_ALIPAY] = "支付宝"
	CashTypeMsg[CASH_TYPE_BANKCARD] = "银行卡"
	CashTypeMsg[CASH_TYPE_AGENT] = "代理商"
}