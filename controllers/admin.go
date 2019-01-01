package controllers

import (
	"fish/configs"
	"fish/enums"
	"fish/managers"
	"fish/models"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"time"
)

type AdminController struct {
	baseController
	baseUrl string
	admin   models.AdminAccount
}

func (c *AdminController) verify_page(code int) {
	if !verifyPower(c.admin, code) {
		c.Abort("10003")
	}
}
func (c *AdminController) verify_ajax(code int) {
	if !verifyPower(c.admin, code) {
		c.Data["json"] = c.jsonData(enums.NOT_ENOUGH_PERMISSION)
		c.ServeJSON()
		c.StopRun()
	}
}
func (c *AdminController) Prepare() {
	c.baseController.Prepare()
	c.baseUrl = configs.Domain["domain"] + configs.Domain["admin"]
	c.Data["domain"] = configs.Domain["domain"]
	c.Data["host"] = c.baseUrl
	c.Data["powers"] = configs.AdminPower
	if !strings.Contains(c.Ctx.Input.URL(), "login") {
		c.Layout = "LayoutAdmin.tpl"
		if !c.checkSession() {
			if c.Ctx.Input.IsPost() || c.Ctx.Input.IsPut() {
				c.Data["json"] = c.jsonData(enums.AGENT_NOT_LOGIN)
				c.ServeJSON()
			} else {
				c.Redirect(c.baseUrl+"login", 302)
			}
		} else {
			c.admin = c.GetSession("admin").(models.AdminAccount)
			c.Data["user"] = c.admin
		}
	}
}
func (c *AdminController) Index() {

}
func (c *AdminController) Login() {
	if c.Ctx.Input.IsPost() {
		if !cpt.VerifyReq(c.Ctx.Request) {
			c.Data["json"] = c.jsonData(enums.CAPTCHA_ERROR)
		} else {
			name := c.GetString("name")
			pwd := c.GetString("pwd")
			admin, err := managers.AdminInstance.GetAdminByName(name)
			if err != nil {
				c.Data["json"] = c.jsonData(enums.AGENT_NOT_FOUND)
			} else {
				result := c.checkAdmin(admin, pwd)
				if result == enums.SUCCESS {
					c.updateSession(admin)
				}
				c.Data["json"] = c.jsonData(result)
			}
		}
		c.ServeJSON()
	}
}
func (c *AdminController) Logout() {
	c.SetSession("admin", nil)
	c.Redirect(c.baseUrl, 302)
}
func (c *AdminController) ChangePwd() {
	if c.Ctx.Input.IsPost() {
		oldPass := c.GetString("old")
		newPass := c.GetString("new")
		newPass1 := c.GetString("new1")
		if c.admin.Password == oldPass {
			if newPass == newPass1 {
				result := managers.AdminInstance.ChangePwd(c.admin, oldPass, newPass)
				c.Data["json"] = c.jsonData(result)
				if result == enums.SUCCESS {
					c.admin.Password = newPass
					c.updateSession(c.admin)
				}
			} else {
				c.Data["json"] = c.jsonData(enums.AGENT_PASSWORD_ERROR)
			}
		} else {
			c.Data["json"] = c.jsonData(enums.AGENT_OLDPASSWORD_ERROR)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) AdminList() {
	c.verify_page(configs.AdminPower["管理员"]["管理员列表"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetAdmins(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		action := c.GetString("action")
		switch action {
		case "add_admin":
			c.verify_ajax(configs.AdminPower["管理员"]["添加管理员"])
			name := c.GetString("admin_name")
			pass := c.GetString("admin_pass")
			powers := c.GetString("admin_powers")
			admin := models.AdminAccount{
				Name:        name,
				Password:    pass,
				Permissions: powers,
				CreateTime:  time.Now(),
				FrozenTime:  time.Now(),
			}
			c.Data["json"] = c.jsonData(managers.AdminInstance.Register(admin))
			break
		case "edit_admin":
			c.verify_ajax(configs.AdminPower["管理员"]["编辑管理员"])
			id, _ := c.GetInt("admin_id")
			powers := c.GetStrings("admin_powers")
			var s_powers string
			for _, v := range powers {
				s_powers += v + ","
			}
			admin := models.AdminAccount{
				Id:          id,
				Permissions: s_powers,
			}
			c.Data["json"] = c.jsonData(managers.AdminInstance.UpdateAdmin(admin))
			break
		default:
			c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
			break
		}
		c.ServeJSON()
	}
}
func (c *AdminController) PlayerList() {
	c.verify_page(configs.AdminPower["玩家"]["玩家列表"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetPlayers(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) PlayerCash() {
	c.verify_page(configs.AdminPower["玩家"]["玩家提现"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetPlayerCashLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		c.verify_ajax(configs.AdminPower["玩家"]["玩家提现"])
		if c.GetString("action") == "action" {
			id, _ := c.GetInt("id")
			state, _ := c.GetInt("state")
			result := managers.AdminInstance.UpdatePlayerCashLogState(c.admin, id, state)
			c.Data["json"] = c.jsonData(result)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) AgentList() {
	c.verify_page(configs.AdminPower["代理"]["代理列表"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetAllAgents(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		action := c.GetString("action")
		switch action {
		case "add_agent":
			c.verify_ajax(configs.AdminPower["管理员"]["添加代理"])
			agent_name := c.GetString("agent_name")
			default_pass := c.GetString("default_pass")
			mobile := c.GetString("mobile")
			nickname := c.GetString("nickname")
			rate, _ := c.GetInt("rate")
			agent := models.AgentAccount{
				Name:         agent_name,
				NickName:     nickname,
				Password:     default_pass,
				Mobile:       mobile,
				Rate:         rate,
				RegisterTime: time.Now()}
			c.Data["json"] = c.jsonData(managers.AgentInstance.Register(agent))
			break
		default:
			c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) AgentApply() {
	c.verify_page(configs.AdminPower["代理"]["代理申请"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetAgentCashLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) AgentCash() {
	c.verify_page(configs.AdminPower["代理"]["代理提现"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.GetAgentCashLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		c.verify_ajax(configs.AdminPower["代理"]["代理提现"])
		if c.GetString("action") == "action" {
			id, _ := c.GetInt("id")
			state, _ := c.GetInt("state")
			result := managers.AdminInstance.UpdateAgentCashLogState(c.admin, id, state)
			c.Data["json"] = c.jsonData(result)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) Recharge() {
	c.verify_page(configs.AdminPower["充值"]["后台充值"])
	if c.Ctx.Input.IsPost() {
		playerId, _ := c.GetInt("id")
		amount, _ := c.GetFloat("amount")
		c.Data["json"] = c.jsonData(managers.AdminInstance.RechargePlayer(playerId, amount))
		c.ServeJSON()
	}
}

func (c *AdminController) Online() {
	c.verify_page(configs.AdminPower["数据统计"]["在线数据"])
	if c.Ctx.Input.IsPost() {
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AdminInstance.Online(searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}

func (c *AdminController) LogRecharge() {
	c.verify_page(configs.AdminPower["日志"]["充值日志"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, rechargeLogs, err := managers.AdminInstance.GetRechargeLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, rechargeLogs, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) LogPlay() {
	c.verify_page(configs.AdminPower["日志"]["战绩日志"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, playLog, err := managers.AdminInstance.GetPlayLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, playLog, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}

func (c *AdminController) SystemConfig() {
	c.verify_page(configs.AdminPower["系统"]["系统配置"])
	c.Data["site"] = configs.Site
	c.Data["db"] = configs.Db
	c.Data["domain"] = configs.Domain
	c.Data["sms"] = configs.Sms
	c.Data["data"] = configs.Data
	c.Data["payment"] = configs.PaymentGate
	if c.Ctx.Input.IsPost() {
		if err := managers.SystemInstanse.ChangeConfig(c.GetString("k"), c.GetString("v")); err != nil {
			logs.Error(err)
			c.Data["json"] = c.jsonData(enums.CHANGE_CONFIG_FAILED)
		} else {
			c.Data["json"] = c.jsonData(enums.SUCCESS)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) checkAdmin(admin models.AdminAccount, pwd string) enums.ReturnCode {
	//if !agent. {
	//	return enums.AGENT_UNUSED
	//}
	if admin.Password != pwd {
		return enums.AGENT_PASSWORD_ERROR
	}
	return enums.SUCCESS
}
func (c *AdminController) updateSession(admin models.AdminAccount) {
	c.admin = admin
	c.SetSession("admin", admin)
}
func (c *AdminController) checkSession() bool {
	admin := c.GetSession("admin")
	if reflect.TypeOf(admin) == reflect.TypeOf(models.AdminAccount{}) {
		if admin.(models.AdminAccount).Id > 0 {
			return true
		}
	}
	return false
}
