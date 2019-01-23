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
	NOT_ENOUGH_PERMISSION   ReturnCode = 10007 //没有权限
	AGENT_NOT_LOGIN         ReturnCode = 20000 //代理没有登录
	AGENT_NOT_FOUND         ReturnCode = 20001 //代理账号不存在
	AGENT_UNUSED            ReturnCode = 20002 //不可用
	AGENT_PASSWORD_ERROR    ReturnCode = 20003 //密码错误
	AGENT_OLDPASSWORD_ERROR ReturnCode = 20004 //旧密码错误
	AGENT_NEWPASSWORD_ERROR ReturnCode = 20005 //新密码不一致
	AGENT_NOT_ENOUGH_GOLD   ReturnCode = 20006 //新密码不一致
	PLAYER_NOT_FOUND        ReturnCode = 30001 //玩家账号不存在
	AMOUNT_MUST_100         ReturnCode = 30002
	ADMIN_NOT_LOGIN         ReturnCode = 40000 //管理没有登录
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

type BankType int

const (
	_        BankType = iota
	ICBC     BankType = 1
	ABC      BankType = 2
	BOC      BankType = 3
	CCB      BankType = 4
	CMB      BankType = 5
	COMM     BankType = 6
	CEB      BankType = 7
	HXB      BankType = 8
	GDB      BankType = 9
	BCCB     BankType = 10
	BJRCB    BankType = 11
	BOS      BankType = 12
	SHRCB    BankType = 13
	CBHB     BankType = 14
	HCCB     BankType = 15
	GZCB     BankType = 16
	CITIC    BankType = 17
	PSBC     BankType = 18
	CIB      BankType = 19
	CMBC     BankType = 20
	SZPAB    BankType = 21
	SPDB     BankType = 22
	URCB     BankType = 23
	NBCB     BankType = 24
	NJCB     BankType = 25
	WZCB     BankType = 26
	CSCB     BankType = 27
	CYB      BankType = 28
	CZB      BankType = 29
	CZCB     BankType = 30
	GNXS     BankType = 31
	HKBCHINA BankType = 32
	ALIPAY   BankType = 33
)

type AdminActionType int

const (
	ADMIN_RECHARGE_PLAYER         AdminActionType = 1
	ADMIN_CHANGE_AGENT_RATE       AdminActionType = 2
	ADMIN_CHANGE_ADMIN_PERMISSION AdminActionType = 3
)

type PaymentChannel int

const (
	PAY_CHANNEL_HUIYI    PaymentChannel = 1
	PAY_CHANNEL_WOHUIBAO PaymentChannel = 2
	PAY_CHANNEL_HONGJIA  PaymentChannel = 3
	PAY_CHANNEL_ZONGHENG PaymentChannel = 4
	PAY_CHANNEL_YIJIA    PaymentChannel = 5
	PAY_CHANNEL_ALIPAY   PaymentChannel = 6
	PAY_CHANNEL_SUNAPI   PaymentChannel = 7
)

type GoldChangeType int

const (
	GOLD_CHANGE_GAME_CATCH_FISH    GoldChangeType = 1  //2.捕鱼游戏输赢
	GOLD_CHANGE_GAME_GOLDEN_FLOWER GoldChangeType = 2  // 2.金花游戏输赢
	GOLD_CHANGE_GAME_DOUDIZU       GoldChangeType = 3  // 3.斗地主输赢
	GOLD_CHANGE_GAME_NIUNIU_M      GoldChangeType = 4  // 4.百人牛牛输赢
	GOLD_CHANGE_GAME_KING_QUEE     GoldChangeType = 5  // 5.红黑大战输赢
	GOLD_CHANGE_RECHARGE           GoldChangeType = 6  // 6.充值
	GOLD_CHANGE_CASH               GoldChangeType = 7  // 7.提现
	GOLD_CHANGE_BANK               GoldChangeType = 8  // 8.银行操作
	GOLD_CHANGE_AWARD              GoldChangeType = 9  // 9.补偿
	GOLD_CHANGE_RETURN             GoldChangeType = 10 // 10.还款
	GOLD_CHANGE_GAME_NIUNIU_Z      GoldChangeType = 11 // 11.抢庄牛牛
	GOLD_CHANGE_GAME_DRAGON_TIGER  GoldChangeType = 12 // 12.龙虎斗
	GOLD_CHANGE_CASH_FAIL          GoldChangeType = 13 // 13.提现失败返还的金币
	GOLD_CHANGE_RED_PACKET         GoldChangeType = 14 // 14.领取红包',
)
