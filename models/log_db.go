package models

import (
	"time"
)

type Player struct {
	Id                int    `json:"id"`
	NickName          string `json:"nick_name"`
	GlobalNum         int    `json:"global_num"`
	BankNum           int    `json:"bank_num"`
	Tax               int    `json:"tax"`
	TotalRechargeSum  int    `json:"total_recharge_sum"`
	AllWithdrawAmount int    `json:"all_withdraw_amount"`
	QmAgentId         int    `json:"qm_agent_id"`
}

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

//管理员给代理充值记录
type AdminRechargeLog struct {
	Id             int       `orm:"column(Id);size(11);" json:"id"`
	RewardMoney    int       `orm:"column(RewardMoney);size(11);" json:"reward_money"`
	AgentId        int       `orm:"column(AgentId);size(11);" json:"agent_id"`
	RechargeTime   time.Time `orm:"column(RechargeTime);" json:"recharge_time"`
	RechargeMoney  int64     `orm:"column(RechargeMoney);size(20);" json:"recharge_money"`
	IsEffective    bool      `orm:"column(IsEffective);" json:"is_effective"`
	IsProcessed    bool      `orm:"column(IsProcessed);" json:"is_processed"`
	RechargeUserId int64     `orm:"column(RechargeUserId);size(20);" json:"recharge_user_id"`
}
type PlayLog struct {
	Id         int       `orm:"column(Id)"`
	GameLogId  string    `orm:"column(GameLogId)"`
	StartGold  int64     `orm:"column(StartGold)"`
	GoldChange int64     `orm:"column(GoldChange)"`
	CreateTime time.Time `orm:"column(CreateTime)"`
	UserId     int       `orm:"column(UserId)"`
	GameId     int       `orm:"column(GameId)"`
	RoomId     int       `orm:"column(RoomId)"`
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
	Alipay             string    `orm:"column(alipay)" json:"alipay"`
	RealName           string    `orm:"column(real_name)" json:"real_name"`
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
type AgentCashInfo struct {
	AgentId        int    `orm:"column();pk;" json:"agent_id"`
	Alipay         string `orm:"column()" json:"alipay"`
	AlipayRealName string `orm:"column()" json:"alipay_real_name"`
	BankType       int    `orm:"column()" json:"bank_type"`
	BankInfo       string `orm:"column()" json:"bank_info"`
	BankCardNo     string `orm:"column()" json:"bank_card_no"`
	BankRealName   string `orm:"column()" json:"bank_real_name"`
}

func (this *PlayLog) TableName() string {
	return "playlog"
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

func (this *AgentAccount) TableUnique() [][]string {
	return [][]string{
		[]string{"name"},
	}
}
