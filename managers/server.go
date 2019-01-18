package managers

import (
	"fish/fishServer"
	"fish/models"
	"github.com/astaxie/beego/orm"
)

type ServerMgr struct {
}

var ServerInstance = newServer()

func newServer() *ServerMgr {
	return new(ServerMgr)
}
func (this *ServerMgr) SendMsg(userId int, serviceId int) {
	o := orm.NewOrm()
	var messages []models.ChatMessages
	o.QueryTable(new(models.ChatMessages)).Filter("UserId", userId).Filter("MessageType", 1).Filter("IsUserMessage", false).Filter("IsProcessed", false).All(&messages)
	for _, msg := range messages {
		fishServer.FishInstance.SendServiceMsg(userId, msg.Id)
		msg.IsProcessed = true
		//msg.KefuId = serviceId
		o.Update(&msg, "IsProcessed")
	}
}
