package enums

type ReturnCode uint

const (
	SUCCESS                 ReturnCode = iota
	CAPTCHA_ERROR           ReturnCode = 10001 //验证码错误
	QUERY_DATA_ERROR        ReturnCode = 10002 //查询数据失败
	DB_ACTION_ERROR         ReturnCode = 10003 //数据库存储失败
	PARAMS_PARSE_ERROR      ReturnCode = 10004 //参数解析错误
	CHANGE_CONFIG_FAILED    ReturnCode = 10005 //修改配置失败
	INVALID_ACTION          ReturnCode = 10006 //无效的操作
	AGENT_NOT_LOGIN         ReturnCode = 20000 //没有登录
	AGENT_NOT_FOUND         ReturnCode = 20001 //代理账号不存在
	AGENT_UNUSED            ReturnCode = 20002 //不可用
	AGENT_PASSWORD_ERROR    ReturnCode = 20003 //密码错误
	AGENT_OLDPASSWORD_ERROR ReturnCode = 20004 //旧密码错误
	AGENT_NEWPASSWORD_ERROR ReturnCode = 20005 //新密码不一致
	PLAYER_NOT_FOUND        ReturnCode = 30001 //玩家账号不存在
)

func (this ReturnCode) String() string {
	return SystemMsg[this]
}

type AgentLevel uint

const (
	_             AgentLevel = iota
	AGENT_LEVEL_1 AgentLevel = 1 //内部商人、优先级最高
	AGENT_LEVEL_2 AgentLevel = 2 //内部商人、优先级最高
	AGENT_LEVEL_3 AgentLevel = 3 //内部商人、优先级最高
	AGENT_LEVEL_4 AgentLevel = 4 //内部商人、优先级最高
	AGENT_LEVEL_5 AgentLevel = 5 //内部商人、优先级最高
)

type GameType int

const (
	GAME_CATCH_FISH    GameType = 1 //捕鱼
	GAME_GOLDEN_FLOWER GameType = 2 //炸金花
	GAME_DOUDIZHU      GameType = 3 //斗地主
	GAME_NIUNIU_M      GameType = 4 //百人牛牛
	GAME_KING_QUEE     GameType = 5 //红黑大战
	GAME_NIUNIU_Z      GameType = 6 //抢庄牛牛
	GAME_DRAGON_TIGER  GameType = 7 //龙虎斗
)

type CashState int

const (
	CASH_STATE_APPLY   CashState = 0 //申请
	CASH_STATE_WAIT    CashState = 1 //等待
	CASH_STATE_REFUSED CashState = 2 //拒绝
	CASH_STATE_SUCCESS CashState = 3 //成功
	CASH_STATE_REFUND  CashState = 4 //退款
)

type CashType int

const (
	CASH_TYPE_ALIPAY   CashType = 0 //支付宝
	CASH_TYPE_BANKCARD CashType = 1 //银行卡
	CASH_TYPE_AGENT    CashType = 2 //代理商
)
