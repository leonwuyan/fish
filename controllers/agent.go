package controllers

import (
	"bytes"
	"fish/configs"
	"fish/enums"
	"fish/managers"
	"fish/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type AgentController struct {
	baseController
	baseUrl string
	agent   models.AgentAccount
}

func (c *AgentController) Prepare() {
	c.baseController.Prepare()
	c.baseUrl = configs.Domain["domain"] + configs.Domain["agent"]
	c.Data["domain"] = configs.Domain["domain"]
	c.Data["host"] = c.baseUrl
	if !strings.Contains(c.Ctx.Input.URL(), "login") {
		c.Layout = "LayoutAgent.tpl"
		if !c.checkSession() {
			if c.Ctx.Input.IsPost() || c.Ctx.Input.IsPut() {
				c.Data["json"] = c.jsonData(enums.AGENT_NOT_LOGIN)
				c.ServeJSON()
			} else {
				c.Redirect(c.baseUrl+"login", 302)
			}
		} else {
			c.agent = c.GetSession("agent").(models.AgentAccount)
			c.Data["user"] = c.agent
		}
	}
	managers.SystemInstance.PageVisitor(c.Ctx.Input, c.GetSession("agent"))
}
func (c *AgentController) Index() {
	if c.Ctx.Input.IsPost() {
		c.Data["json"] = c.jsonData(enums.SUCCESS, managers.AgentInstance.GetAgentInfo(c.agent), 0)
		c.ServeJSON()
	}
}
func (c *AgentController) Login() {
	if c.Ctx.Input.IsPost() {
		if !cpt.VerifyReq(c.Ctx.Request) {
			c.Data["json"] = c.jsonData(enums.CAPTCHA_ERROR)
		} else {
			name := c.GetString("name")
			pwd := c.GetString("pwd")
			agent, err := managers.AgentInstance.GetAgentByName(name)
			if err != nil {
				c.Data["json"] = c.jsonData(enums.AGENT_NOT_FOUND)
			} else {
				result := c.checkAgent(agent, pwd)
				if result == enums.SUCCESS {
					c.updateSession(agent)
				}
				c.Data["json"] = c.jsonData(result)
			}
		}
		c.ServeJSON()
	}
}
func (c *AgentController) Logout() {
	c.SetSession("agent", nil)
	c.Redirect(c.baseUrl, 302)
}
func (c *AgentController) ChangePwd() {
	if c.Ctx.Input.IsPost() {
		oldPass := c.GetString("old")
		newPass := c.GetString("new")
		newPass1 := c.GetString("new1")
		if c.agent.Password == oldPass {
			if newPass == newPass1 {
				result := managers.AgentInstance.ChangePwd(c.agent, oldPass, newPass)
				c.Data["json"] = c.jsonData(result)
				if result == enums.SUCCESS {
					c.agent.Password = newPass
					c.updateSession(c.agent)
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
func (c *AgentController) Generalize() {
	adUrl := configs.Domain["domain"] + "advertise?agentId=" + strconv.Itoa(c.agent.Id)
	c.Data["ad_url"] = adUrl
}
func (c *AgentController) GeneralizeQr() {
	templateId := c.Ctx.Input.Param(":id")
	adUrl := configs.Domain["domain"] + "advertise?agentId=" + strconv.Itoa(c.agent.Id)
	img, _ := c.createQr(fmt.Sprintf("static/img/bg%s.png", templateId), adUrl)
	c.Ctx.Output.ContentType("png")
	var b bytes.Buffer
	jpeg.Encode(&b, img, &jpeg.Options{Quality: 75})
	c.Ctx.Output.Body(b.Bytes())
}
func (c *AgentController) Agents() {
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AgentInstance.GetChildren(c.agent, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		switch c.GetString("action") {
		case "add_agent_child":
			parent_id, _ := c.GetInt("agent_id")
			agent_name := c.GetString("agent_name")
			default_pass := c.GetString("default_pass")
			mobile := c.GetString("mobile")
			nickname := c.GetString("nickname")
			rate, _ := c.GetInt("rate")
			if rate <= c.agent.Rate {
				newAgent := models.AgentAccount{
					Name:         agent_name,
					Password:     default_pass,
					Mobile:       mobile,
					NickName:     nickname,
					Rate:         rate,
					ParentId:     parent_id,
					RegisterTime: time.Now(),
				}
				c.Data["json"] = c.jsonData(managers.AgentInstance.Register(newAgent))
			} else {
				c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
			}
			break
			//case "change_rate":
			//	agent_id, _ := c.GetInt("agent_id")
			//	rate, _ := c.GetInt("rate")
			//	if rate <= c.agent.Rate {
			//		c.Data["json"] = c.jsonData(managers.AgentInstance.ChangeChildRate(c.agent, agent_id, rate))
			//	} else {
			//		c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
			//	}
			//	break
		default:
			c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
			break
		}
		c.ServeJSON()
	}
}
func (c *AgentController) ChildTax() {
	if c.Ctx.Input.IsPost() {
		agent_id, _ := c.GetInt("id")
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		child := models.AgentAccount{Id: agent_id}
		total, agents, err := managers.AgentInstance.GetTax(child, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AgentController) Players() {
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, players, err := managers.AgentInstance.GetPlayers(c.agent, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, players, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AgentController) Tax() {
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AgentInstance.GetTax(c.agent, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AgentController) Cash() {
	if c.Ctx.Input.IsPost() {
		pageSize, _ := c.GetInt("pageSize")
		pageIndex, _ := c.GetInt("pageIndex")
		searchParams := c.GetString("searchParams")
		total, agents, err := managers.AgentInstance.GetCashLog(c.agent, pageSize, pageIndex, searchParams)
		if err == nil {
			c.Data["json"] = c.jsonData(enums.SUCCESS, agents, total)
		} else {
			c.Data["json"] = c.jsonData(enums.QUERY_DATA_ERROR)
			logs.Error(err)
		}
		c.ServeJSON()
	}
}
func (c *AgentController) CashApply() {
	if c.Ctx.Input.IsPost() {
		infos, _ := managers.AgentInstance.GetBankInfos(c.agent)
		c.Data["json"] = c.jsonData(enums.SUCCESS, infos, len(infos))
		c.ServeJSON()
	}
	if c.Ctx.Input.IsPut() {
		action := c.GetString("action")
		switch action {
		case "save":
			cashType, _ := c.GetInt("cash_type")
			bankCode, _ := c.GetInt("bank_code")
			bankName := c.GetString("bank_name")
			bankCard := c.GetString("bank_card")
			realName := c.GetString("real_name")
			data := models.BankCardInfo{
				AgentId:    c.agent.Id,
				CashType:   cashType,
				BankType:   bankCode,
				BankName:   bankName,
				BankCardNo: bankCard,
				RealName:   realName,
			}
			c.Data["json"] = c.jsonData(managers.AgentInstance.AddBankInfo(data))
			break
		case "apply":
			amount, _ := c.GetInt("amount")
			bankInfoId, _ := c.GetInt("bank_info_id")
			if amount/100 < 1 || amount%100 > 0 {
				c.Data["json"] = c.jsonData(enums.AMOUNT_MUST_100)
			} else {
				c.Data["json"] = c.jsonData(managers.AgentInstance.CashApply(c.agent, amount, bankInfoId))
			}
			break
		default:
			c.Data["json"] = c.jsonData(enums.INVALID_ACTION)
			break
		}
		c.ServeJSON()
	}
}
func (c *AgentController) checkAgent(agent models.AgentAccount, pwd string) enums.ReturnCode {
	//if !agent. {
	//	return enums.AGENT_UNUSED
	//}
	if agent.Password != pwd {
		return enums.AGENT_PASSWORD_ERROR
	}
	return enums.SUCCESS
}
func (c *AgentController) updateSession(agent models.AgentAccount) {
	c.agent = agent
	c.SetSession("agent", agent)
}
func (c *AgentController) checkSession() bool {
	agent := c.GetSession("agent")
	if reflect.TypeOf(agent) == reflect.TypeOf(models.AgentAccount{}) {
		if agent.(models.AgentAccount).Id > 0 {
			return true
		}
	}
	return false
}

func (c *AgentController) createQr(bgPath, info string) (newImg draw.Image, err error) {
	bgFile, err := os.Open(bgPath)
	defer bgFile.Close()
	if err != nil {
		return
	}
	logoFile, err := os.Open("static/img/logo.png")
	defer logoFile.Close()
	if err != nil {
		return
	}
	bgImg, _ := jpeg.Decode(bgFile)
	logoImg, _ := jpeg.Decode(logoFile)

	qrCode, _ := qrcode.New(info, qrcode.Highest)
	qrImg := qrCode.Image(232)
	logoImgSize := qrImg.Bounds().Max.X / 4
	logoImg = resize.Thumbnail(uint(logoImgSize), uint(logoImgSize), logoImg, resize.Lanczos3)
	newImg = image.NewRGBA64(bgImg.Bounds())
	global_offset_Y := 657
	qrImg_offset := image.Pt(bgImg.Bounds().Max.X/2-qrImg.Bounds().Max.X/2, global_offset_Y-qrImg.Bounds().Max.Y/2)
	logoImg_offset := qrImg_offset.Add(image.Pt(qrImg.Bounds().Max.X/2-logoImg.Bounds().Max.X/2, qrImg.Bounds().Max.Y/2-logoImg.Bounds().Max.Y/2))
	draw.Draw(newImg, bgImg.Bounds(), bgImg, bgImg.Bounds().Min, draw.Over)
	draw.Draw(newImg, qrImg.Bounds().Add(qrImg_offset), qrImg, qrImg.Bounds().Min, draw.Src)
	draw.Draw(newImg, qrImg.Bounds().Add(logoImg_offset), logoImg, logoImg.Bounds().Min, draw.Src)
	return
}
