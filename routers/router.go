package routers

import (
	"fish/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/pay", &controllers.MainController{}, "*:Recharge")
	beego.Router("/advertise", &controllers.MainController{}, "*:Advertise")
	beego.Router("/forms/*", &controllers.MainController{}, "*:GetPostForms")
	ns_admin := beego.NewNamespace("/admin",
		beego.NSRouter("/", &controllers.AdminController{}, "*:Index"),
		beego.NSRouter("/login", &controllers.AdminController{}, "*:Login"),
		beego.NSRouter("/logout", &controllers.AdminController{}, "*:Logout"),
		beego.NSRouter("/changepwd", &controllers.AdminController{}, "*:ChangePwd"),
		beego.NSRouter("/sysconfig", &controllers.AdminController{}, "*:SystemConfig"),

		beego.NSRouter("/player/cash", &controllers.AdminController{}, "*:PlayerCash"),
		beego.NSRouter("/player/list", &controllers.AdminController{}, "*:PlayerList"),

		beego.NSRouter("/agent/list", &controllers.AdminController{}, "*:AgentList"),
		beego.NSRouter("/agent/apply", &controllers.AdminController{}, "*:AgentApply"),
		beego.NSRouter("/agent/cash", &controllers.AdminController{}, "*:AgentCash"),

		beego.NSRouter("/recharge", &controllers.AdminController{}, "*:Recharge"),
	)
	ns_agent := beego.NewNamespace("/agent",
		beego.NSRouter("/", &controllers.AgentController{}, "*:Index"),
		beego.NSRouter("/login", &controllers.AgentController{}, "*:Login"),
		beego.NSRouter("/logout", &controllers.AgentController{}, "*:Logout"),
		beego.NSRouter("/agents", &controllers.AgentController{}, "*:Agents"),
		beego.NSRouter("/childtax", &controllers.AgentController{}, "*:ChildTax"),
		beego.NSRouter("/generalize", &controllers.AgentController{}, "*:Generalize"),
		beego.NSRouter("/players", &controllers.AgentController{}, "*:Players"),
		beego.NSRouter("/tax", &controllers.AgentController{}, "*:Tax"),
		beego.NSRouter("/cash", &controllers.AgentController{}, "*:Cash"),
		beego.NSRouter("/changepwd", &controllers.AgentController{}, "*:ChangePwd"),
	)
	beego.AddNamespace(ns_admin, ns_agent)
}
