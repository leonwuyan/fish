package routers

import (
	"fish/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/forms/*", &controllers.MainController{}, "*:GetPostForms")
	//推广
	beego.Router("/advertise", &controllers.MainController{}, "*:Advertise")
	beego.Router("/advertise/:agent([0-9]+)_:id([0-9]+).png", &controllers.MainController{}, "*:GeneralizeQr")
	//充值
	beego.Router("/pay", &controllers.MainController{}, "*:Recharge")
	//beego.Router("/pay", &controllers.MainController{}, "*:RechargeHuiYi")
	//beego.Router("/pay", &controllers.MainController{}, "*:RechargeWoHuiBao")
	//充值返回
	beego.Router("/notify/hui_yi", &controllers.MainController{}, "*:Notify_Hui_Yi")
	beego.Router("/notify/wo_hui_bao", &controllers.MainController{}, "*:Notify_Wo_Hui_Bao")
	beego.Router("/notify/hong_jia", &controllers.MainController{}, "*:Notify_Hong_Jia")
	beego.Router("/notify/zong_heng", &controllers.MainController{}, "*:Notify_Zong_Heng")
	beego.Router("/notify/yi_jia", &controllers.MainController{}, "*:Notify_Yi_Jia")
	ns_admin := beego.NewNamespace("/admin",
		beego.NSRouter("/", &controllers.AdminController{}, "*:Index"),
		beego.NSRouter("/login", &controllers.AdminController{}, "*:Login"),
		beego.NSRouter("/logout", &controllers.AdminController{}, "*:Logout"),
		beego.NSRouter("/changepwd", &controllers.AdminController{}, "*:ChangePwd"),
		beego.NSRouter("/sys/config", &controllers.AdminController{}, "*:SystemConfig"),
		beego.NSRouter("/sys/channel", &controllers.AdminController{}, "*:ChannelInfo"),

		beego.NSRouter("/admin/list", &controllers.AdminController{}, "*:AdminList"),

		beego.NSRouter("/player/cash", &controllers.AdminController{}, "*:PlayerCash"),
		beego.NSRouter("/player/list", &controllers.AdminController{}, "*:PlayerList"),

		beego.NSRouter("/agent/list", &controllers.AdminController{}, "*:AgentList"),
		beego.NSRouter("/agent/apply", &controllers.AdminController{}, "*:AgentApply"),
		beego.NSRouter("/agent/cash", &controllers.AdminController{}, "*:AgentCash"),

		beego.NSRouter("/recharge", &controllers.AdminController{}, "*:Recharge"),

		beego.NSRouter("/services/message", &controllers.AdminController{}, "*:ChatMessages"),

		beego.NSRouter("/statistic/online", &controllers.AdminController{}, "*:Online"),
		beego.NSRouter("/statistic/game", &controllers.AdminController{}, "*:StatisticsPlay"),

		beego.NSRouter("/log/recharge", &controllers.AdminController{}, "*:LogRecharge"),
		beego.NSRouter("/log/play", &controllers.AdminController{}, "*:LogPlay"),

		beego.NSRouter("/balance", &controllers.AdminController{}, "*:Balance"),
		beego.NSRouter("/ws", &controllers.AdminController{}, "*:SysMessage"),
	)
	ns_agent := beego.NewNamespace("/agent",
		beego.NSRouter("/", &controllers.AgentController{}, "*:Index"),
		beego.NSRouter("/login", &controllers.AgentController{}, "*:Login"),
		beego.NSRouter("/logout", &controllers.AgentController{}, "*:Logout"),
		beego.NSRouter("/agents", &controllers.AgentController{}, "*:Agents"),
		beego.NSRouter("/childtax", &controllers.AgentController{}, "*:ChildTax"),
		beego.NSRouter("/generalize", &controllers.AgentController{}, "*:Generalize"),
		//beego.NSRouter("/generalize_qr/:id([0-9]+).png", &controllers.AgentController{}, "*:GeneralizeQr"),
		beego.NSRouter("/players", &controllers.AgentController{}, "*:Players"),
		beego.NSRouter("/tax", &controllers.AgentController{}, "*:Tax"),
		beego.NSRouter("/cash", &controllers.AgentController{}, "*:Cash"),
		beego.NSRouter("/cash/apply", &controllers.AgentController{}, "*:CashApply"),
		beego.NSRouter("/changepwd", &controllers.AgentController{}, "*:ChangePwd"),
	)
	ns_api := beego.NewNamespace("/api",
		beego.NSRouter("/", &controllers.ApiController{}),
	)
	beego.AddNamespace(ns_admin, ns_agent, ns_api)
}
