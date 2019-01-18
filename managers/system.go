package managers

import (
	"fish/alisms"
	"fish/configs"
	"fish/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"time"
)

type SystemMgr struct {
}

var SystemInstance = newSystem()

func newSystem() *SystemMgr {
	return new(SystemMgr)
}
func (this *SystemMgr) Recharge() {

}
func (this *SystemMgr) PreRecharge(userId, channel, payType, amount int, payOrder string) (err error) {
	o := orm.NewOrm()
	player, err := PlayerInstance.GetPlayerById(userId)
	if err != nil {
		return
	}
	rechargeData := models.RechargeLog{
		UserId:          userId,
		RechargeChannel: channel,
		RechargeType:    payType,
		TransactionId:   payOrder,
		GoldChange:      int64(amount),
		RechargeTime:    time.Now(),
		FinishTime:      time.Unix(0, 0),
		SendTime:        time.Unix(0, 0),
		AgentId:         player.AgentId,
	}
	_, err = o.Insert(&rechargeData)
	if err != nil {
		this.PreRecharge(userId, channel, payType, amount, payOrder)
	}
	return
}
func (this *SystemMgr) FinishRecharge(payOrder string) (err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(models.RechargeLog)).Filter("TransactionId", payOrder).Filter("Finished", 0).Update(orm.Params{"Finished": 1, "FinishTime": time.Now()}); err != nil {
		return
	}
	return
}
func (this *SystemMgr) SendSms(phone, text string) (err error) {
	err = alisms.SendSms(
		configs.Sms["access_key_id"],
		configs.Sms["access_secret"],
		phone,
		configs.Sms["sign_name"],
		"{\"code\":\""+text+"\"}",
		configs.Sms["template_code"])
	return
}
func (this *SystemMgr) ChangeConfig(key, value string) (err error) {
	if err = beego.AppConfig.Set(key, value); err == nil {
		err = beego.AppConfig.SaveConfigFile("conf/app.conf")
	}
	return
}
func (this *SystemMgr) PageVisitor(input *context.BeegoInput, session interface{}) {
	visitData := models.LogPageVisit{
		Page:      input.URI(),
		Method:    input.Method(),
		Params:    fmt.Sprintf("%+v", input.Context.Request.PostForm),
		User:      fmt.Sprintf("%+v", session),
		VisitTime: time.Now(),
	}
	o := orm.NewOrm()
	o.Insert(&visitData)
}
