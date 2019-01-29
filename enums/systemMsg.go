package enums

var SystemMsg map[ReturnCode]string
var CashStateMsg map[CashState]string
var CashTypeMsg map[CashType]string
var BankTypeName map[BankType]string
var PaymentChannelName map[PaymentChannel]string

func init() {
	SystemMsg = make(map[ReturnCode]string)
	SystemMsg[SUCCESS] = "成功"
	SystemMsg[QUERY_DATA_ERROR] = "查询数据失败"
	SystemMsg[DB_ACTION_ERROR] = "数据库存储失败"
	SystemMsg[CAPTCHA_ERROR] = "验证码错误"
	SystemMsg[AGENT_NOT_LOGIN] = "没有登录"
	SystemMsg[AGENT_NOT_FOUND] = "账号不存在"
	SystemMsg[NOT_ENOUGH_PERMISSION] = "权限不足，请与管理员联系"
	SystemMsg[AGENT_UNUSED] = "不可用"
	SystemMsg[AGENT_PASSWORD_ERROR] = "密码错误"
	SystemMsg[AGENT_OLDPASSWORD_ERROR] = "旧密码错误"
	SystemMsg[AGENT_NEWPASSWORD_ERROR] = "新密码不一致"
	SystemMsg[PLAYER_NOT_FOUND] = "账号不存在"
	SystemMsg[AMOUNT_MUST_100] = "提取金额必须是100的倍数，最低100"

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

	BankTypeName = make(map[BankType]string)
	BankTypeName[BANK_ICBC] = "工商银行"
	BankTypeName[BANK_ABC] = "农业银行"
	BankTypeName[BANK_BOC] = "中国银行"
	BankTypeName[BANK_CCB] = "建设银行"
	BankTypeName[BANK_CMB] = "招商银行"
	BankTypeName[BANK_COMM] = "交通银行"
	BankTypeName[BANK_CEB] = "光大银行"
	BankTypeName[BANK_HXB] = "华夏银行"
	BankTypeName[BANK_GDB] = "广发银行"
	BankTypeName[BANK_BCCB] = "北京银行"
	BankTypeName[BANK_BJRCB] = "北京农商行"
	BankTypeName[BANK_BOS] = "上海银行"
	BankTypeName[BANK_SHRCB] = "上海农商银行"
	BankTypeName[BANK_CBHB] = "渤海银行"
	BankTypeName[BANK_HCCB] = "杭州银行"
	BankTypeName[BANK_GZCB] = "广州市商业银行"
	BankTypeName[BANK_CITIC] = "中信银行"
	BankTypeName[BANK_PSBC] = "中国邮储银行"
	BankTypeName[BANK_CIB] = "兴业银行"
	BankTypeName[BANK_CMBC] = "民生银行"
	BankTypeName[BANK_SZPAB] = "平安银行"
	BankTypeName[BANK_SPDB] = "浦发银行"
	BankTypeName[BANK_URCB] = "杭州联合银行"
	BankTypeName[BANK_NBCB] = "宁波银行"
	BankTypeName[BANK_NJCB] = "南京银行"
	BankTypeName[BANK_WZCB] = "温州市商业银行"
	BankTypeName[BANK_CSCB] = "长沙银行"
	BankTypeName[BANK_CYB] = "集友银行"
	BankTypeName[BANK_CZB] = "浙商银行"
	BankTypeName[BANK_CZCB] = "浙江稠州商业银行"
	BankTypeName[BANK_GNXS] = "广州市农信社"
	BankTypeName[BANK_HKBCHINA] = "汉口银行"
	BankTypeName[BANK_ALIPAY] = "支付宝"

	PaymentChannelName = make(map[PaymentChannel]string)
	PaymentChannelName[PAY_CHANNEL_HUIYI] = "汇易"
	PaymentChannelName[PAY_CHANNEL_WOHUIBAO] = "沃汇宝"
	PaymentChannelName[PAY_CHANNEL_HONGJIA] = "宏佳"
}
