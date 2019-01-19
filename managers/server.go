package managers

import (
	"fish/fishServer"
	"fish/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ServerMgr struct {
}

var ServerInstance = newServer()

func newServer() *ServerMgr {
	return new(ServerMgr)
}
func (this *ServerMgr) SendMsg(userId int) {
	o := orm.NewOrm()
	var messages []models.ChatMessages
	o.QueryTable(new(models.ChatMessages)).Filter("UserId", userId).Filter("MessageType", 1).Filter("IsUserMessage", false).Filter("IsProcessed", false).All(&messages)
	for _, msg := range messages {
		this.notifyChatMessages(userId, msg.Id)
		msg.IsProcessed = true
		//msg.KefuId = serviceId
		o.Update(&msg, "IsProcessed")
	}
}
func (this *ServerMgr) ChangeChannelInfo(data models.Channel) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(data); err != nil {
		logs.Error(err)
	} else {
		this.notifyChannelInfo(data.ChannelId)
	}
	return
}
func (this *ServerMgr) ChangeNotice(data models.Notice) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(data); err != nil {
		logs.Error(err)
	} else {
		this.notifyNotice()
	}
	return
}
func (this *ServerMgr) ChangeShowAgent(data models.AgentShow) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(data); err != nil {
		logs.Error(err)
	} else {
		this.notifyShowAgent(data.AgentId)
	}
	return
}
func (this *ServerMgr) notifyChatMessages(userId, msgId int) {
	fishServer.FishInstance.SendServiceMsg(userId, msgId)
}
func (this *ServerMgr) notifyChannelInfo(channelId int) {
	fishServer.FishInstance.ChangeChannelConfig(channelId, false)
}
func (this *ServerMgr) notifyNotice() {
	fishServer.FishInstance.ChangeNoticeConfig()
}
func (this *ServerMgr) notifyShowAgent(agentId int) {
	fishServer.FishInstance.ChangeAgentConfig(agentId, false)
}
