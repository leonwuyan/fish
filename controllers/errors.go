package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["code"] = 404
	c.Data["content"] = "你访问的页面被外星人偷走啦"
	c.TplName = "error.tpl"
}
func (c *ErrorController) Error10001() {
	c.Data["code"] = 10001
	c.Data["content"] = "创建充值订单失败"
	c.TplName = "error.tpl"
}
func (c *ErrorController) Error10002() {
	c.Data["code"] = 10002
	c.Data["content"] = "无效的充值渠道"
	c.TplName = "error.tpl"
}

