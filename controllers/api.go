package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	action := c.GetString("action")
	switch action {
	case "CustomServiceMsg":
		msgId, _ := c.GetInt("user_id")
		content := c.GetString("content")
		addCustomServiceMsg(msgId, content)
		break
	}
}
func addCustomServiceMsg(userId int, content string) {
	//managers.AdminInstance.AddCustomServiceMsg(userId, content)
}
