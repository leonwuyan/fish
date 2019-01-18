package controllers

import (
	"fish/configs"
	"fish/enums"
	"fish/managers"
	"fish/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var IsRunning = false
var clients map[*websocket.Conn]string

type AdminController struct {
	baseController
	baseUrl   string
	admin     models.AdminAccount
	wsUpgrade websocket.Upgrader
}

func init() {
	clients = make(map[*websocket.Conn]string)
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
			if c.Ctx.Request.Header.Get("Upgrade") == "websocket" {
				//println(c.Ctx.Request.Header.Get("Upgrade"))
				ws, err := c.wsUpgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
				if err != nil {
					logs.Error(err)
				}
				ws.WriteJSON(c.jsonData(enums.ADMIN_NOT_LOGIN))
				ws.Close()
			}
			if c.Ctx.Input.IsPost() || c.Ctx.Input.IsPut() {
				c.Data["json"] = c.jsonData(enums.ADMIN_NOT_LOGIN)
				c.ServeJSON()
			} else {
				c.Redirect(c.baseUrl+"login", 302)
			}
		} else {
			c.admin = c.GetSession("admin").(models.AdminAccount)
			c.Data["user"] = c.admin
		}
	}
	managers.SystemInstance.PageVisitor(c.Ctx.Input, c.GetSession("admin"))
}
func (c *AdminController) Index() {
	o := orm.NewOrm()
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日注册"]) {
		player_reg_count, _ := o.QueryTable(new(models.PlayerAccount)).Filter("RsgLogTime__gte", time.Today()).Count()
		c.Data["player_reg_count"] = player_reg_count
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日活跃"]) {
		var player_login_count int
		o.Raw("SELECT COUNT(DISTINCT UserId) FROM loginlog WHERE LoginTime>=?", time.Today()).QueryRow(&player_login_count)
		c.Data["player_login_count"] = player_login_count
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日有效"]) {
		var player_play_count int
		o.Raw("SELECT COUNT(DISTINCT UserId) FROM playlog WHERE CreateTime>=?", time.Today()).QueryRow(&player_play_count)
		c.Data["player_play_count"] = player_play_count
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日充值"]) {
		var player_pay_count, player_pay_gold int
		o.Raw("SELECT COUNT(DISTINCT UserId),SUM(GoldChange) FROM rechargelog WHERE Finished=1 AND RechargeType<100 AND RechargeTime>=?", time.Today()).QueryRow(&player_pay_count, &player_pay_gold)
		c.Data["player_pay_count"] = player_pay_count
		c.Data["player_pay_gold"] = strconv.FormatFloat(float64(player_pay_gold)/100, 'f', 2, 64)
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日税收"]) {
		var player_tax_count, player_tax_gold int
		o.Raw("SELECT COUNT(DISTINCT UserId),SUM(Pumping) FROM pumplogdetail WHERE CreateTime>=?", time.Today()).QueryRow(&player_tax_count, &player_tax_gold)
		c.Data["player_tax_count"] = player_tax_count
		c.Data["player_tax_gold"] = strconv.FormatFloat(float64(player_tax_gold)/100, 'f', 2, 64)
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["玩家提现"]) {
		var player_cash_count, player_cash_gold int
		o.Raw("SELECT COUNT(DISTINCT UserId),SUM(Gold) FROM withdrawalslog WHERE State=3 AND WithdrawalsLogTime>=?", time.Today()).QueryRow(&player_cash_count, &player_cash_gold)
		c.Data["player_cash_count"] = player_cash_count
		c.Data["player_cash_gold"] = strconv.FormatFloat(float64(player_cash_gold)/100, 'f', 2, 64)
	}
	if verifyPower(c.admin, configs.AdminPower["概况"]["今日盈利"]) {
	}
}
func (c *AdminController) SysMessage() {
	c.TplName = "index.tpl"
	ws, err := c.wsUpgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		logs.Error(err)
	}
	clients[ws] = c.Ctx.Input.IP()
	if !IsRunning {
		go c.handlerClientMessage()
		IsRunning = true
	}
}
func (c *AdminController) handlerClientMessage() {
	for {
		cashApply, chat := managers.AdminInstance.IsNewMessage()
		var messages []models.Result
		if cashApply > 0 {
			messages = append(messages, c.jsonData(enums.SUCCESS, models.ClientMessage{0, "有新的提现申请，请尽快处理"}))
		}
		if chat > 0 {
			messages = append(messages, c.jsonData(enums.SUCCESS, models.ClientMessage{1, "有新的玩家消息，请尽快处理"}))
		}
		if len(messages) == 0 {
			messages = append(messages, c.jsonData(enums.SUCCESS, models.ClientMessage{2, "暂无新的消息"}))
		}
		//清理客户端连接
		//println(len(clients))
		for client, _ := range clients {
			for _, msg := range messages {
				if client.WriteJSON(msg) != nil {
					delete(clients, client)
				}
			}
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

//游戏配置
func (c *AdminController) ChannelList() {

}
func (c *AdminController) ChannelInfo() {
	id, _ := c.GetInt("id")
	println(id)
}
func (c *AdminController) ShowAgentList() {

}
func (c *AdminController) NoticeList() {

}

//功能部分
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
			if agents == nil {
				agents = []models.AdminAccount{}
			}
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
			if agents == nil {
				agents = []models.PlayerAccount{}
			}
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
			if agents == nil {
				agents = []models.PlayerCashLog{}
			}
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
			if agents == nil {
				agents = []models.AgentAccount{}
			}
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
		total, agents, err := managers.AdminInstance.GetAgentApply(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			if agents == nil {
				agents = []models.AgentApply{}
			}
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
			if agents == nil {
				agents = []models.AgentCashLog{}
			}
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
		rechargeType, _ := c.GetInt("type")
		c.Data["json"] = c.jsonData(managers.AdminInstance.RechargePlayer(playerId, amount, rechargeType))
		c.ServeJSON()
	}
}

//客服
func (c *AdminController) ChatMessages() {
	c.verify_page(configs.AdminPower["客服"]["消息"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, msg, err := managers.AdminInstance.GetChatMessages(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			if msg == nil {
				msg = []models.ChatMessages{}
			}
			c.Data["json"] = c.jsonData(enums.SUCCESS, msg, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		action := c.GetString("action")
		switch action {
		case "add_service_msg":
			c.verify_ajax(configs.AdminPower["客服"]["回复消息"])
			userId, _ := c.GetInt("user_id")
			msgId, _ := c.GetInt("msg_id")
			//msgType, _ := c.GetInt("msg_type")
			content := c.GetString("msg_content")
			managers.AdminInstance.UpdateChatMessage(msgId)
			c.Data["json"] = c.jsonData(managers.AdminInstance.AddCustomServiceMsg(c.admin, userId, content))
			break
		case "set_state":
			msgId, _ := c.GetInt("id")
			c.Data["json"] = c.jsonData(managers.AdminInstance.UpdateChatMessage(msgId))
			break
		default:
			c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) ChatWindow() {
	c.verify_page(configs.AdminPower["客服"]["消息"])
	if c.Ctx.Input.IsPost() {
		//userId, _ := c.GetInt("uid")
		//msgs := []models.ChatMessages{}
	}
}

//数据统计部分
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
func (c *AdminController) StatisticsRecharge() {
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
func (c *AdminController) StatisticsPlay() {
	c.verify_page(configs.AdminPower["数据统计"]["游戏数据"])
	if c.Ctx.Input.IsPost() {
		beginTime, _ := time.ParseInLocation("2006-01-02", c.GetString("begin"), time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02", c.GetString("end"), time.Local)
		total, data, err := managers.AdminInstance.StatisticsPlay(beginTime, endTime)
		if err == nil {
			if data == nil {
				data = []models.StatisticsPlay{}
			}
			c.Data["json"] = c.jsonData(enums.SUCCESS, data, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}

//日志部分
func (c *AdminController) LogRecharge() {
	c.verify_page(configs.AdminPower["日志"]["充值日志"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, rechargeLogs, err := managers.AdminInstance.GetRechargeLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			if rechargeLogs == nil {
				rechargeLogs = []models.RechargeLog{}
			}
			c.Data["json"] = c.jsonData(enums.SUCCESS, rechargeLogs, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AdminController) LogPlay() {
	c.Data["beginT"] = time.Today().Format("2006-01-02T15:04")
	c.Data["endT"] = time.Now().Format("2006-01-02T15:04")
	c.verify_page(configs.AdminPower["日志"]["战绩日志"])
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, playLog, err := managers.AdminInstance.GetPlayLogs(c.admin, pageSize, pageIndex, searchParams)
		if err == nil {
			if playLog == nil {
				playLog = []models.PlayLog{}
			}
			c.Data["json"] = c.jsonData(enums.SUCCESS, playLog, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}

//结算
func (c *AdminController) Balance() {
	c.verify_page(configs.AdminPower["结算"]["收益结算"])
	if c.Ctx.Input.IsPost() {
		beginTime, _ := time.ParseInLocation("2006-01-02", c.GetString("begin"), time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02", c.GetString("end"), time.Local)
		endTime = endTime.AddDate(0, 0, 1)
		var sysWinA, sysFeeA, sysWin, sysFee int
		var agentFee, agentTax float64
		data := make(map[string]string)
		if verifyPower(c.admin, configs.AdminPower["结算"]["代理结算"]) {
			sysWinA, sysFeeA = managers.AdminInstance.GetHasAgentSystemWin(beginTime, endTime)
			agentFee = managers.AdminInstance.GetAgentFee(beginTime, endTime)
			agentTax = managers.AdminInstance.GetAgentTax(beginTime, endTime)
			//有代理系统输赢
			data["sysWinA"] = strconv.FormatFloat(float64(sysWinA)/100, 'f', 2, 64)
			//有代理系统税收
			data["sysFeeA"] = strconv.FormatFloat(float64(sysFeeA)/100, 'f', 2, 64)
			//代理税收
			data["agentFee"] = strconv.FormatFloat(float64(agentFee)/100, 'f', 2, 64)
			//代理收益
			data["agentTax"] = strconv.FormatFloat(float64(agentTax)/100, 'f', 2, 64)
		}
		if verifyPower(c.admin, configs.AdminPower["结算"]["系统结算"]) {
			sysWin, sysFee = managers.AdminInstance.GetSystemWin(beginTime, endTime)
			//系统输赢
			data["sysWin"] = strconv.FormatFloat(float64(sysWin)/100, 'f', 2, 64)
			//系统税收
			data["sysFee"] = strconv.FormatFloat(float64(sysFee)/100, 'f', 2, 64)
			//非代理系统输赢
			data["sysWinO"] = strconv.FormatFloat(float64(sysWin-sysWinA)/100, 'f', 2, 64)
			//非代理系统税收
			data["sysFeeO"] = strconv.FormatFloat(float64(sysFee-sysFeeA)/100, 'f', 2, 64)
		}
		if verifyPower(c.admin, configs.AdminPower["结算"]["其他结算"]) {
			data["playerRecharge"] = strconv.FormatFloat(float64(managers.AdminInstance.GetPlayerRecharge(beginTime, endTime))/100, 'f', 2, 64)
			data["playerRechargeAward"] = strconv.FormatFloat(float64(managers.AdminInstance.GetPlayerRechargeAward(beginTime, endTime))/100, 'f', 2, 64)
			data["playerRechargeForTest"] = strconv.FormatFloat(float64(managers.AdminInstance.GetPlayerRechargeForTest(beginTime, endTime))/100, 'f', 2, 64)
			data["sysAward"] = strconv.FormatFloat(float64(managers.AdminInstance.GetSystemAward(beginTime, endTime))/100, 'f', 2, 64)
			data["playerCash"] = strconv.FormatFloat(float64(managers.AdminInstance.GetPlayerCash(beginTime, endTime))/100, 'f', 2, 64)
			data["sysPunish"] = strconv.FormatFloat(float64(managers.AdminInstance.GetSystemPunish(beginTime, endTime))/100, 'f', 2, 64)
			data["playerRemain"] = strconv.FormatFloat(float64(managers.AdminInstance.GetPlayerRemain())/100, 'f', 2, 64)
			data["dev"] = strconv.FormatFloat((float64(sysWin-sysWinA+sysFee-sysFeeA)*0.5+(float64(sysWinA+sysFeeA)-agentFee)*0.35+float64(agentFee-agentTax)*0.5)/100, 'f', 2, 64)
		}
		c.Data["json"] = c.jsonData(enums.SUCCESS, data, 0)
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
		if err := managers.SystemInstance.ChangeConfig(c.GetString("k"), c.GetString("v")); err != nil {
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
