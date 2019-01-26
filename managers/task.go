package managers

import (
	"fish/configs"
	"fish/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"strconv"
)

type TaskMgr struct {
}

var TaskInstance = newTask()
var AgentFeeBlock = false
var SmsBlock = false
var agentBindBlock = false

func newTask() *TaskMgr {
	return new(TaskMgr)
}

func (this *TaskMgr) Init() {
	agentFee := toolbox.NewTask("agentFee", "0/1 * * * * *", this.agentFee)
	smsSend := toolbox.NewTask("smsSend", "0/1 * * * * *", this.smsSend)
	agentPlayerStatistic := toolbox.NewTask("agentPlayerStatistic", "0/1 * * * * *", this.agentBind)
	toolbox.AddTask("agentFee", agentFee)
	toolbox.AddTask("smsSend", smsSend)
	toolbox.AddTask("agentPlayerStatistic", agentPlayerStatistic)
	if configs.TaskEnabled {
		toolbox.StartTask()
	}
}
func (this *TaskMgr) agentFee() (err error) {
	if !AgentFeeBlock {
		AgentFeeBlock = true
		o := orm.NewOrm()
		var playLogs []models.PlayLog
		var lastLoadId = configs.Data["last_play_log"]
		_, err = o.QueryTable(new(models.PlayLog)).Filter("Id__gt", lastLoadId).All(&playLogs)
		if err != nil {
			logs.Error(err)
		}
		for _, playLog := range playLogs {
			if err = AgentInstance.FeeToAgent(playLog); err != nil {
				logs.Error("fee failed,log_id:%d,reason:%s", playLog.Id, err.Error())
			}
			lastLoadId = strconv.Itoa(playLog.Id)
		}
		if lastLoadId != configs.Data["last_play_log"] {
			SystemInstance.ChangeConfig("data::last_play_log", lastLoadId)
		}
		AgentFeeBlock = false
	}
	return nil
}
func (this *TaskMgr) smsSend() (err error) {
	if !SmsBlock {
		SmsBlock = true
		o := orm.NewOrm()
		var smsLogs []models.SmsLog
		var lastSmsId = configs.Data["last_sms_log"]
		_, err = o.QueryTable(new(models.SmsLog)).Filter("Id__gt", lastSmsId).All(&smsLogs)
		if err != nil {
			logs.Error(err)
		}
		for _, smsLog := range smsLogs {
			if err = SystemInstance.SendSms(smsLog.PhoneNumber, smsLog.Text); err != nil {
				logs.Error("send sms failed,log_id:%d,reason:%s", smsLog.Id, err.Error())
			}
			lastSmsId = strconv.Itoa(smsLog.Id)
		}
		if lastSmsId != configs.Data["last_sms_log"] {
			SystemInstance.ChangeConfig("data::last_sms_log", lastSmsId)
		}
		SmsBlock = false
	}
	return nil
}
func (this *TaskMgr) agentBind() (err error) {
	if !agentBindBlock {
		agentBindBlock = true
		o := orm.NewOrm()
		var bindLogs []models.User2Agent
		lastBindId := configs.Data["last_bind_agent"]
		_, err = o.QueryTable(new(models.User2Agent)).Filter("Id__gt", lastBindId).All(&bindLogs)
		if err != nil {
			logs.Error(err)
		}
		for _, bindLog := range bindLogs {
			if err = AgentInstance.Statistic_player(bindLog.QmAgentId, true); err != nil {
				logs.Error("math agent's player error,bind_id:%d,reason:%s", bindLog.Id, err.Error())
			}
			lastBindId = strconv.Itoa(bindLog.Id)
		}
		if lastBindId != configs.Data["last_bind_agent"] {
			SystemInstance.ChangeConfig("data::last_bind_agent", lastBindId)
		}
		agentBindBlock = false
	}
	return nil
}
func (this *TaskMgr) playerTax() (err error) {
	return
}
func (this *TaskMgr) playerCash() (err error) {
	return
}
