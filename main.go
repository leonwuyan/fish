package main

import (
	"fish/configs"
	"fish/controllers"
	"fish/fishServer"
	"fish/managers"
	"fish/models"
	_ "fish/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

var logPath = "./logs"
var logCfg = `{"filename":"%s/log.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(logCfg, logPath))
	beego.SetLogFuncCall(true)
}
func main() {
	models.RegisterDB()
	managers.TaskInstance.Init()
	gameServerIp := configs.GameServer["ip"]
	gameServerPort, _ := strconv.Atoi(configs.GameServer["port"])
	gameServerSign, _ := strconv.Atoi(configs.GameServer["sign"])
	go fishServer.Start(gameServerIp, gameServerPort, gameServerSign)
	beego.Run()
}
