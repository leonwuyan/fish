package models

import (
	"fish/configs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", configs.Db["log"])
	orm.RegisterModel(
		new(AdminAccount),
		new(AgentAccount),
		new(AgentFeeLog),
		new(PlayLog),
		new(SmsLog),
		new(AgentCashLog),
		new(PlayerCashLog),
		new(User2Agent),
		new(RechargeLog),
		new(PlayerAccount),
		new(BankCardInfo),
	)
	orm.RunSyncdb("default", false, true)
	orm.Debug = beego.BConfig.RunMode == "dev"
}
