package main

import (
	"bytes"
	bs "encoding/base64"
	"fish/configs"
	"fish/controllers"
	"fish/fishServer"
	"fish/managers"
	"fish/models"
	_ "fish/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skip2/go-qrcode"
	"html/template"
	"image/png"
	"strconv"
)

var logPath = "./logs"
var logCfg = `{"filename":"%s/log.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(logCfg, logPath))
	beego.SetLogFuncCall(true)
	beego.AddFuncMap("create_qr", createQr)
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
func createQr(msg string) template.HTML {
	code, _ := qrcode.New(msg, qrcode.Medium)
	img := code.Image(220)
	emptyBuff := bytes.NewBuffer([]byte{})         //开辟一个新的空buff
	png.Encode(emptyBuff, img)                     //img写入到buff
	dist := make([]byte, 50000)                    //开辟存储空间
	bs.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	n := 0
	for i := 0; i < len(dist); i++ {
		if dist[i] == 0 {
			n = i
			break
		}
	}
	str := string(dist[:n])
	ls1 := "<img src='data:image/png;base64,"
	ls2 := "'/>"
	ls := ls1 + str + ls2
	return template.HTML(ls)
}
