package enums

var SystemMsg map[ReturnCode]string
var CashStateMsg map[CashState]string
var CashTypeMsg map[CashType]string
var BankTypeName map[BankType]string
var PaymetChannelName map[PaymentChannel]string

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
	BankTypeName[ICBC] = "工商银行"
	BankTypeName[ABC] = "农业银行"
	BankTypeName[BOC] = "中国银行"
	BankTypeName[CCB] = "建设银行"
	BankTypeName[CMB] = "招商银行"
	BankTypeName[COMM] = "交通银行"
	BankTypeName[CEB] = "光大银行"
	BankTypeName[HXB] = "华夏银行"
	BankTypeName[GDB] = "广发银行"
	BankTypeName[BCCB] = "北京银行"
	BankTypeName[BJRCB] = "北京农商行"
	BankTypeName[BOS] = "上海银行"
	BankTypeName[SHRCB] = "上海农商银行"
	BankTypeName[CBHB] = "渤海银行"
	BankTypeName[HCCB] = "杭州银行"
	BankTypeName[GZCB] = "广州市商业银行"
	BankTypeName[CITIC] = "中信银行"
	BankTypeName[PSBC] = "中国邮储银行"
	BankTypeName[CIB] = "兴业银行"
	BankTypeName[CMBC] = "民生银行"
	BankTypeName[SZPAB] = "平安银行"
	BankTypeName[SPDB] = "浦发银行"
	BankTypeName[URCB] = "杭州联合银行"
	BankTypeName[NBCB] = "宁波银行"
	BankTypeName[NJCB] = "南京银行"
	BankTypeName[WZCB] = "温州市商业银行"
	BankTypeName[CSCB] = "长沙银行"
	BankTypeName[CYB] = "集友银行"
	BankTypeName[CZB] = "浙商银行"
	BankTypeName[CZCB] = "浙江稠州商业银行"
	BankTypeName[GNXS] = "广州市农信社"
	BankTypeName[HKBCHINA] = "汉口银行"
	BankTypeName[ALIPAY] = "支付宝"

	PaymetChannelName = make(map[PaymentChannel]string)
	PaymetChannelName[PAY_CHANNEL_HUIYI] = "汇易"
	PaymetChannelName[PAY_CHANNEL_WOHUIBAO] = "沃汇宝"
	PaymetChannelName[PAY_CHANNEL_HONGJIA] = "宏佳"
}
