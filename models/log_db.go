package models

import (
	"time"
)

//代理申请表
type AgentApply struct {
	Id        int       `orm:"column(Id);size(11);pk;auto;" json:"id"`
	UserId    int       `orm:"column(UserId);size(11);" json:"user_id"`    //用户Id
	ApplyTime time.Time `orm:"column(ApplyTime);" json:"apply_time"`       //申请时间
	Name      string    `orm:"column(Name);size(45);" json:"name"`         //姓名
	Phone     string    `orm:"column(Phone);size(45);" json:"phone"`       //电话
	Email     string    `orm:"column(Email);size(145);" json:"email"`      //邮箱
	QQ        string    `orm:"column(qq);size(145);" json:"qq"`            //QQ
	WeiXin    string    `orm:"column(WeiXin);size(145);" json:"wei_xin"`   //微信
	Message   string    `orm:"column(Message);size(1445);" json:"message"` //消息内容
	IsDeal    bool      `orm:"column(IsDeal);" json:"is_deal"`             //是否处理
	DealTime  time.Time `orm:"column(DealTime);" json:"deal_time"`         //处理时间
	Reply     string    `orm:"column(Reply);size(1445);" json:"reply"`     //回复内容
}
type AdminAccount struct {
	Id          int       `orm:"column(id)" json:"id"`
	Name        string    `orm:"column(name)" json:"name"`
	Password    string    `orm:"column(password)" json:"-"`
	Permissions string    `orm:"column(permissions)" json:"permissions"`
	CreateTime  time.Time `orm:"column(create_time)" json:"create_time"`
	FrozenTime  time.Time `orm:"column(frozen_time)" json:"frozen_time"`
}
type AgentFeeLog struct {
	Id         int       `orm:"column(id)" json:"id"`
	LogId      int       `orm:"column(log_id)" json:"log_id"`
	UserId     int       `orm:"column(user_id)" json:"user_id"`
	UserName   string    `orm:"column(user_name)" json:"user_name"`
	AgentId    int       `orm:"column(agent_id)" json:"agent_id"`
	AgentChild int       `orm:"column(agent_child)" json:"agent_child"`
	GameId     int       `orm:"column(game_id)" json:"game_id"`
	RoomId     int       `orm:"column(room_id)" json:"room_id"`
	Tax        float64   `orm:"column(tax)" json:"tax"`
	Rate       int       `orm:"column(rate)" json:"rate"`
	ChildRate  int       `orm:"column(child_rate)" json:"child_rate"`
	Fee        float64   `orm:"column(fee);digits(18);decimals(6);" json:"fee"`
	WinLose    float64   `orm:"column(win_lose);digits(18);decimals(6);" json:"win_lose"`
	LogTime    time.Time `orm:"column(log_time)" json:"log_time"`
}
type RechargeLog struct {
	Id              int       `orm:"column(Id)" json:"id"`
	UserId          int       `orm:"column(UserId)" json:"user_id"`
	RechargeType    int       `orm:"column(RechargeType)" json:"recharge_type"`
	RechargeMoney   float32   `orm:"column(RechargeMoney)" json:"recharge_money"`
	PlayerGold      int64     `orm:"column(PlayerGold)" json:"player_gold"`
	BankGold        int64     `orm:"column(BankGold)" json:"bank_gold"`
	GoldChange      int64     `orm:"column(GoldChange)" json:"gold_change"`
	RechargeTime    time.Time `orm:"column(RechargeTime)" json:"recharge_time"`
	ChannelType     int       `orm:"column(ChannelType)" json:"channel_type"`
	RechargeChannel int       `orm:"column(RechargeChannel)" json:"recharge_channel"`
	TransactionId   string    `orm:"column(TransactionId)" json:"transaction_id"`
	Finished        byte      `orm:"column(Finished)" json:"finished"`
	IsRead          bool      `orm:"column(IsRead)" json:"is_read"`
	FinishTime      time.Time `orm:"column(FinishTime)" json:"finish_time"`
	IsUse           bool      `orm:"column(IsUse)" json:"is_use"`
	IsSend          bool      `orm:"column(IsSend)" json:"is_send"`
	SendTime        time.Time `orm:"column(SendTime)" json:"send_time"`
	AgentId         int       `orm:"column(AgentId)" json:"agent_id"`
	FlashId         string    `orm:"column(FlashId)" json:"flash_id"`
	BuyerLogonId    string    `orm:"column(buyer_logon_id)" json:"buyer_logon_id"`
	PayerAlipayId   string    `orm:"column(payer_alipay_id)" json:"payer_alipay_id"`
}
type AgentAccount struct {
	Id                     int       `orm:"column(id)" json:"id"`
	Name                   string    `orm:"column(name)" json:"name"`
	NickName               string    `orm:"column(nick_name)" json:"nick_name"`
	Password               string    `orm:"column(password)" json:"-"`
	Level                  int       `orm:"column(level)" json:"level"`
	Gold                   float64   `orm:"column(gold);digits(18);decimals(6);" json:"gold"`
	ParentId               int       `orm:"column(parent_id)" json:"parent_id"`
	Mobile                 string    `orm:"column(mobile)" json:"mobile"`
	TotalChildrenImmediate int       `orm:"column(total_children_immediate)" json:"total_children_immediate"`
	TotalChildrenOther     int       `orm:"column(total_children_other)" json:"total_children_other"`
	TotalPlayersImmediate  int       `orm:"column(total_players_immediate)" json:"total_players_immediate"`
	TotalPlayersOther      int       `orm:"column(total_players_other)" json:"total_players_other"`
	TotalTax               float64   `orm:"column(total_tax);digits(18);decimals(6);" json:"total_tax"`
	TotalFee               float64   `orm:"column(total_fee);digits(18);decimals(6);" json:"total_fee"`
	TotalCash              float64   `orm:"column(total_cash);digits(18);decimals(6);" json:"total_cash"`
	PlayerWin              float64   `orm:"column(player_win);digits(18);decimals(6);" json:"player_win"`
	PlayerLose             float64   `orm:"column(player_lose);digits(18);decimals(6);" json:"player_lose"`
	Rate                   int       `orm:"column(rate)" json:"rate"`
	RegisterTime           time.Time `orm:"column(register_time)" json:"register_time"`
}
type SmsLog struct {
	Id           int       `orm:"column(Id)"`
	PhoneNumber  string    `orm:"column(PhoneNumber)"`
	Text         string    `orm:"column(Text)"`
	CreationTime time.Time `orm:"column(CreationTime)"`
	IsSend       bool      `orm:"column(IsSend)"`
	ChannelId    int       `orm:"column(ChannelId)"`
}
type AgentCashLog struct {
	Id                 int       `orm:"column(id)" json:"id"`
	AgentId            int       `orm:"column(agent_id)" json:"agent_id"`
	Gold               int       `orm:"column(gold)" json:"gold"`
	OrderId            string    `orm:"column(order_id)" json:"order_id"`
	Fee                int       `orm:"column(fee)" json:"fee"`
	GetMoney           int       `orm:"column(get_money)" json:"get_money"`
	TxType             int       `orm:"column(tx_type)" json:"tx_type"`
	BankCardTypeId     int       `orm:"column(bank_type)" json:"bank_card_type_id"`
	BankCardNo         string    `orm:"column(bank_card_no)" json:"bank_card_no"`
	RealName           string    `orm:"column(real_name)" json:"real_name"`
	Alipay             string    `orm:"column(alipay)" json:"alipay"`
	AlipayName         string    `orm:"column(alipay_name)" json:"alipay_name"`
	Mobile             string    `orm:"column(mobile)" json:"mobile"`
	WithdrawalsLogTime time.Time `orm:"column(WithdrawalsLogTime)" json:"withdrawals_log_time"`
	State              byte      `orm:"column(state)" json:"state"`
}
type PlayerCashLog struct {
	Id                 int       `orm:"column(Id)" json:"id"`
	UserId             int       `orm:"column(UserID)" json:"user_id"`
	Gold               int64     `orm:"column(Gold)" json:"gold"`
	GetMoney           int64     `orm:"column(GetMoney)" json:"get_money"`
	OrderId            string    `orm:"column(OrderId)" json:"order_id"`
	Fee                int64     `orm:"column(Fee)" json:"fee"`
	WithdrawalsLogTime time.Time `orm:"column(WithdrawalsLogTime)" json:"withdrawals_log_time"`
	State              byte      `orm:"column(State)" json:"state"`
	HandleTime         time.Time `orm:"column(HandleTime)" json:"handle_time"`
	BillNumber         string    `orm:"column(BillNumber)" json:"bill_number"`
	WithdrawalsChannel int       `orm:"column(WithdrawalsChannel)" json:"withdrawals_channel"`
	HandleComment      string    `orm:"column(HandleComment)" json:"handle_comment"`
	Alipay             string    `orm:"column(Alipay)" json:"alipay"`
	AlipayName         string    `orm:"column(Alipayname)" json:"alipay_name"`
	Username           string    `orm:"column(Username)" json:"username"`
	TxType             int       `orm:"column(TxType)" json:"tx_type"`
	BankCardTypeId     int       `orm:"column(BankCardTypeId)" json:"bank_card_type_id"`
	BankCardNo         string    `orm:"column(BankCardNo)" json:"bank_card_no"`
	RealName           string    `orm:"column(RealName)" json:"real_name"`
	DirtyGold          int64     `orm:"column(DirtyGold)" json:"dirty_gold"`
	IsProcessed        bool      `orm:"column(IsProcessed)" json:"is_processed"`
}
type User2Agent struct {
	Id        int `orm:"column(Id)" json:"id"`
	UserId    int `orm:"column(UserId)" json:"user_id"`
	QmAgentId int `orm:"column(QmAgentId)" json:"qm_agent_id"`
}
type BankCardConfig struct {
	Id             int    `orm:"column(Id)" json:"id"`
	BankCardTypeId int    `orm:"column(BankCardTypeId)" json:"bank_card_type_id"`
	BankName       string `orm:"column(BankName)" json:"bank_name"`
	BankCode       string `orm:"column(BankCode)" json:"bank_code"`
	IsEnabled      bool   `orm:"column(IsEnabled)" json:"is_enabled"`
}
type BankCardInfo struct {
	Id         int    `orm:"column(id)" json:"id"`
	AgentId    int    `orm:"column(agent_id)" json:"agent_id"`
	CashType   int    `orm:"column(cash_type)" json:"cash_type"`
	BankType   int    `orm:"column(bank_type)" json:"bank_card_type_id"`
	BankName   string `orm:"column(bank_name)" json:"bank_name"`
	BankCardNo string `orm:"column(bank_card_no)" json:"bank_card_no"`
	RealName   string `orm:"column(real_name)" json:"real_name"`
	IsDefault  bool   `orm:"column(is_default)" json:"is_default"`
}

type LoginLog struct {
	Id        int       `orm:"column(Id)" json:"id"`
	UserId    int       `orm:"column(UserId)" json:"user_id"`
	LoginTime time.Time `orm:"column(LoginTime)"json:"login_time"`
	Ip        int       `orm:"column(Ip)" json:"ip"`
}
type PlayLog struct {
	Id         int       `orm:"column(Id)" json:"id"`
	GameLogId  string    `orm:"column(GameLogId)" json:"game_log_id"`
	StartGold  int64     `orm:"column(StartGold)" json:"start_gold"`
	GoldChange int64     `orm:"column(GoldChange)" json:"gold_change"`
	CreateTime time.Time `orm:"column(CreateTime)" json:"create_time"`
	UserId     int       `orm:"column(UserId)" json:"user_id"`
	GameId     int       `orm:"column(GameId)" json:"game_id"`
	RoomId     int       `orm:"column(RoomId)" json:"room_id"`
	AgentId    int       `orm:"column(AgentId)" json:"agent_id"`
}
type GoldChangeLog struct {
	Id           int       `orm:"column(Id)" json:"id"`
	UserId       int       `orm:"column(UserId)" json:"user_id"`
	PlayerGold   int64     `orm:"column(PlayerGold)" json:"player_gold"`
	BankGold     int64     `orm:"column(BankGold)" json:"bank_gold"`
	GoldChange   int64     `orm:"column(GoldChange)" json:"gold_change"`
	ChangeType   byte      `orm:"column(ChangeType)" json:"change_type"`
	ChangeText   string    `orm:"column(ChangeText)" json:"change_text"`
	ChangeTime   time.Time `orm:"column(ChangeTime)" json:"change_time"`
	ExPlayerGold int64     `orm:"column(ExPlayerGold)" json:"ex_player_gold"`
	ExBankGold   int64     `orm:"column(ExBankGold)" json:"ex_bank_gold"`
	AgentId      int       `orm:"column(AgentId)" json:"agent_id"`
}
type PumpLogDetail struct {
	Id         int       `orm:"column(Id)" json:"id"`
	UserId     int       `orm:"column(UserId)" json:"user_id"`
	CreateTime time.Time `orm:"column(CreateTime)" json:"create_time"`
	Pumping    int       `orm:"column(Pumping)" json:"pumping"`
	GameId     int       `orm:"column(GameId)" json:"game_id"`
	RoomId     int       `orm:"column(RoomId)" json:"room_id"`
	AgentId    int       `orm:"column(AgentId)" json:"agent_id"`
}
type LogPageVisit struct {
	Id        int       `orm:"column(id)" json:"id"`
	Page      string    `orm:"column(page)" json:"page"`
	Method    string    `orm:"column(method)" json:"method"`
	Params    string    `orm:"column(params)" json:"params"`
	User      string    `orm:"column(user)" json:"user"`
	VisitTime time.Time `orm:"column(visit_time)" json:"visit_time"`
}
type ChatMessages struct {
	Id            int       `orm:"column(Id)" json:"id"`
	CreationTime  time.Time `orm:"column(CreationTime)" json:"creation_time"`
	UserId        int       `orm:"column(UserId)" json:"user_id"`
	MessageType   int       `orm:"column(MessageType)" json:"message_type"`
	Message       string    `orm:"column(Message)" json:"message"`
	IsProcessed   bool      `orm:"column(IsProcessed)" json:"is_processed"`
	KefuId        int       `orm:"column(KefuId)" json:"kefu_id"`
	IsUserMessage bool      `orm:"column(IsUserMessage)" json:"is_user_message"`
	IsDeleted     bool      `orm:"column(IsDeleted)" json:"is_deleted"`
}

//配置部分
type Channel struct {
	Id                      int    `orm:"column(Id)"`
	ChannelId               int    `orm:"column(ChannelId)"`               //渠道编号
	ChannelName             string `orm:"column(ChannelName)"`             //渠道名称
	MainChannel             string `orm:"column(MainChannel)"`             //主渠道
	Platform                string `orm:"column(Platform)"`                //平台
	ServerVersion           int    `orm:"column(ServerVersion)"`           //服务器版本
	GameShow                bool   `orm:"column(gameshow)"`                //是否在游戏中显示
	IsTiXianEnabled         bool   `orm:"column(IsTiXianEnabled)"`         //是否启用提现
	MaxTiXianGold           int    `orm:"column(MaxTiXianGold)"`           //最大提现金额
	PayUrl                  string `orm:"column(PayUrl)"`                  //充值的Url
	GameIdentifier          string `orm:"column(GameIdentifier)"`          //不知道
	IapName                 string `orm:"column(IapName)"`                 //不知道
	InitMoney               int    `orm:"column(InitMoney)"`               //游客注册赠送金币
	UpAccountMoney          int    `orm:"column(UpAccountMoney)"`          //升级正式账号赠送金币
	Remarks                 string `orm:"column(Remarks)"`                 //介绍？
	IsAgentTiXianEnabled    bool   `orm:"column(IsAgentTiXianEnabled)"`    //代理提现是否启用
	IsBankCardTiXianEnabled bool   `orm:"column(IsBankCardTiXianEnabled)"` //银行卡提现是否启用
	IsBankBindingEnabled    bool   `orm:"column(IsBankBindingEnabled)"`    //银行卡绑定是否启用
	IsUnionPayEnabled       bool   `orm:"column(IsUnionPayEnabled)"`       //银联充值是否启用
	IsQQPayEnabled          bool   `orm:"column(IsQQPayEnabled)"`          //QQ充值是否开启
	IsJDPayEnabled          bool   `orm:"column(IsJDPayEnabled)"`          //京东充值是否启用
	IsWxPayEnabled          bool   `orm:"column(IsWxPayEnabled)"`          //微信充值是否开启
	IsAlipayEnabled         bool   `orm:"column(IsAlipayEnabled)"`         //支付宝充值是否开启
	UpdateCount             string `orm:"column(UpdateCount)"`             //更新次数
	IsDkPayEnabled          bool   `orm:"column(IsDkPayEnabled)"`          //点卡支付是否开启
	IsAliTixianEnabeld      bool   `orm:"column(IsAliTixianEnabeld)"`      //是否允许支付宝提现
	IsAliBindingEnabled     bool   `orm:"column(IsAliBindingEnabled)"`     //是否打开支付宝绑定
	SubPlatformName         string `orm:"column(SubPlatformName)"`         //分平台名称
}
type AgentShow struct {
	AgentId          int    `orm:"column(AgentId);pk"`
	AgentName        string `orm:"column(AgentName)"`        //代理商名称，登陆页面使用
	Password         string `orm:"column(Password)"`         //登陆密码
	Payment          string `orm:"column(Payment)"`          //充值密码
	SignKey          string `orm:"column(SignKey)"`          //加密串密钥
	Remarks          string `orm:"column(Remarks)"`          //代理商描述，游戏内显示名称
	QQ               string `orm:"column(QQ)"`               //代理商QQ号码
	WenXin           string `orm:"column(WenXin)"`           //代理商微信号码
	Gold             int64  `orm:"column(Gold)"`             //代理商余额,单位分
	IsUse            bool   `orm:"column(IsUse)"`            //是否可用
	Level            int    `orm:"column(Level)"`            //等级:1.内部商人、优先级最高;2.集团外商、优先级次等;3.嫡系商人、非集团控制的嫡系合作商人;4.普通商人、市面上合作的商人\\\\n
	ShowInGame       bool   `orm:"column(ShowInGame)"`       //是否在游戏里显示该代理商
	MinGold          int64  `orm:"column(MinGold)"`          //保底金额
	IsMinGoldEnabled bool   `orm:"column(IsMinGoldEnabled)"` //是否启用保底金额
	UpdateCount      int    `orm:"column(UpdateCount)"`      //更新次数
	ChannelId        int    `orm:"column(ChannelId)"`        //渠道Id
}
type Notice struct {
	Id         int       `orm:"column(Id)"`         //自增长编号//,
	ChannelId  int       `orm:"column(ChannelId)"`  //渠道编号
	StartTime  time.Time `orm:"column(StartTime)"`  //开始时间
	EndTime    time.Time `orm:"column(EndTime)"`    //结束时间
	Content    string    `orm:"column(Content)"`    //内容
	Title      string    `orm:"column(Title)"`      //标题
	CreateTime time.Time `orm:"column(CreateTime)"` //创建时间
	IsActive   bool      `orm:"column(IsActive)"`   //是否启用
	Url        string    `orm:"column(Url)"`        //官网Url
}

func (this *AgentApply) TableName() string {
	return "agentapply"
}
func (this *SmsLog) TableName() string {
	return "smslog"
}
func (this *PlayerCashLog) TableName() string {
	return "withdrawalslog"
}
func (this *User2Agent) TableName() string {
	return "qmagent_user2agent"
}
func (this *RechargeLog) TableName() string {
	return "rechargelog"
}
func (this *BankCardConfig) TableName() string {
	return "bankcardconfig"
}
func (this *PlayLog) TableName() string {
	return "v_playlog"
}
func (this *LoginLog) TableName() string {
	return "loginlog"
}
func (this *GoldChangeLog) TableName() string {
	return "v_goldchangelog"
}
func (this *PumpLogDetail) TableName() string {
	return "v_pumplogdetail"
}
func (this *LogPageVisit) TableName() string {
	return "log_page_visit"
}
func (this *ChatMessages) TableName() string {
	return "chatmessages"
}
func (this *Channel) TableName() string {
	return "channel"
}
func (this *AgentShow) TableName() string {
	return "agent"
}
func (this *Notice) TableName() string {
	return "notice"
}

func (this *AgentAccount) TableUnique() [][]string {
	return [][]string{
		[]string{"name"},
	}
}
