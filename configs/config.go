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
var AdminPower map[string]map[string]int

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
func init() {
	loadPermissions()
}
