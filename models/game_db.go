package models

import "time"

type PlayerAccount struct {
	UserId            int       `orm:"column(UserId);pk" json:"user_id"`
	AccountName       string    `orm:"column(AccountName)" json:"account_name"`
	NickName          string    `orm:"column(NickName)" json:"nick_name"`
	GlobalNum         int       `orm:"column(GlobalNum)" json:"global_num"`
	BankNum           int       `orm:"column(BankNum)" json:"bank_num"`
	Tax               int       `orm:"column(Tax)" json:"tax"`
	TotalRechargeSum  int       `orm:"column(TotalRechargeSum)" json:"total_recharge_sum"`
	TotalCashSum      int       `orm:"column(TotalCashSum)" json:"total_cash_sum"`
	AllWithdrawAmount int       `orm:"column(AllWithdrawAmount)" json:"all_withdraw_amount"`
	RsgLogTime        time.Time `orm:"column(RsgLogTime)" json:"rsg_log_time"`
	LastLogonTime     time.Time `orm:"column(LastLogonTime)" json:"last_logon_time"`
	IsFreeze          time.Time `orm:"column(IsFreeze)" json:"is_freeze"`
	FreezeEndTime     time.Time `orm:"column(FreezeEndTime)" json:"freeze_end_time"`
	OnlineTime        int       `orm:"column(OnlineTime)" json:"online_time"`
	AgentId           int       `orm:"column(AgentId)" json:"agent_id"`
}

func (this *PlayerAccount) TableName() string {
	return "v_playeraccount"
}
