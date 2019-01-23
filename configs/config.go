package configs

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

var Site, _ = beego.AppConfig.GetSection("site")
var Db, _ = beego.AppConfig.GetSection("db")
var Data, _ = beego.AppConfig.GetSection("data")
var Domain, _ = beego.AppConfig.GetSection("domain")
var Sms, _ = beego.AppConfig.GetSection("sms")
var PaymentGate, _ = beego.AppConfig.GetSection("payment")
var AdminPower map[string]map[string]int
var TaskEnabled, _ = beego.AppConfig.Bool("task_enabled")
var GameServer, _ = beego.AppConfig.GetSection("game_server")
var PayAmountFloor, _ = beego.AppConfig.Bool("pay_amount_floor")
var DownUrls []string

func loadPermissions() {
	data, err := ioutil.ReadFile("conf/admin.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &AdminPower)
	if err != nil {
		return
	}
}
func loadDownloadUrls() {
	data, err := ioutil.ReadFile("conf/download.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &DownUrls)
	if err != nil {
		return
	}
}
func init() {
	loadPermissions()
	loadDownloadUrls()
}
