package managers

import (
	"encoding/json"
	"fish/enums"
	"fish/models"
	"fish/payment"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type AdminMgr struct {
}

var AdminInstance = newAdmin()

func newAdmin() *AdminMgr {
	return new(AdminMgr)
}
func (this *AdminMgr) Register(account models.AdminAccount) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.Insert(&account); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) UpdateAdmin(account models.AdminAccount) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.QueryTable(account).Filter("Id", account.Id).Update(orm.Params{"Permissions": account.Permissions}); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) GetAdminById(id int) (account models.AdminAccount, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(account).Filter("Id", id).One(&account)
	return
}
func (this *AdminMgr) GetAdminByName(name string) (account models.AdminAccount, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(account).Filter("Name", name).One(&account)
	return
}
func (this *AdminMgr) ChangePwd(account models.AdminAccount, oldPwd, newPwd string) enums.ReturnCode {
	if account.Password != oldPwd {
		return enums.AGENT_OLDPASSWORD_ERROR
	}
	o := orm.NewOrm()
	account.Password = newPwd
	if _, err := o.Update(&account, "Password"); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) GetAdmins(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, admins []models.AdminAccount, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AdminAccount))
	var id string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("Id", id)
	}
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).All(&admins)
	}
	return
}
func (this *AdminMgr) GetPlayers(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, players []models.PlayerAccount, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.PlayerAccount))
	var id string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserId", id)
	}
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-UserId").All(&players)
	}
	return
}
func (this *AdminMgr) GetPlayerCashLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, cashLogs []models.PlayerCashLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.PlayerCashLog))
	var id, state string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			state = map_params["state"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserID", id)
	}
	if state != "-1" {
		rs = rs.Filter("State", state)
	}
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&cashLogs)
	return
}
func (this *AdminMgr) UpdatePlayerCashLogState(account models.AdminAccount, id int, state int) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.PlayerCashLog)).Filter("Id", id).Update(orm.Params{"State": state}); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) GetAllAgents(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, agents []models.AgentAccount, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentAccount))
	var id, name, mobile string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			name = map_params["name"]
			mobile = map_params["mobile"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("Id", id)
	}
	if len(name) > 0 {
		rs = rs.Filter("Name", name)
	}
	if len(mobile) > 0 {
		rs = rs.Filter("Mobile", mobile)
	}
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).All(&agents)
	return
}
func (this *AdminMgr) GetAgentCashLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, cashLogs []models.AgentCashLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentCashLog))
	var id, state string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			state = map_params["state"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("AgentId", id)
	}
	if state != "-1" {
		rs = rs.Filter("State", state)
	}
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&cashLogs)
	return
}
func (this *AdminMgr) UpdateAgentCashLogState(account models.AdminAccount, id int, state int) enums.ReturnCode {
	o := orm.NewOrm()
	o.Begin()
	var log models.AgentCashLog
	o.QueryTable(log).Filter("Id", id).One(&log)
	//拒绝兑换，需要将金币退还给代理
	if state == 2 {
		if _, err := o.QueryTable(new(models.AgentAccount)).Filter("Id", log.AgentId).Update(orm.Params{"Gold": orm.ColValue(orm.ColAdd, log.Gold)}); err != nil {
			o.Rollback()
			return enums.DB_ACTION_ERROR
		}
	}
	if _, err := o.QueryTable(log).Filter("Id", id).Update(orm.Params{"State": state}); err != nil {
		o.Rollback()
		return enums.DB_ACTION_ERROR
	}
	o.Commit()
	return enums.SUCCESS
}

func (this *AdminMgr) Online(search string) (total int64, result []orm.Params, err error) {
	o := orm.NewOrm()
	rs := o.Raw("SELECT StatDate,AllCount FROM onlinelog WHERE StatDate BETWEEN ? AND ?")
	if len(search) > 0 {
		var map_params map[string]string
		if err := json.Unmarshal([]byte(search), &map_params); err == nil {
			begin := map_params["begin"]
			end := map_params["end"]
			endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end+" 23:59:59", time.Local)
			rs = rs.SetArgs(begin, endTime)
		} else {
			logs.Error(err)
		}
	}
	total, err = rs.Values(&result)
	return
}
func (this *AdminMgr) RechargePlayer(id int, amount float64) enums.ReturnCode {
	_, err := PlayerInstance.GetPlayerById(id)
	if err != nil {
		return enums.PLAYER_NOT_FOUND
	}
	payOrder := payment.Create_order()
	SystemInstanse.PreRecharge(id, 105, 99, int(amount*100), payOrder)
	SystemInstanse.FinishRecharge(payOrder)
	return enums.SUCCESS
}

func (this *AdminMgr) GetRechargeLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, rechargeLogs []models.RechargeLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.RechargeLog))
	var id string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			begin := map_params["begin"]
			end := map_params["end"]
			endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end+" 23:59:59", time.Local)
			rs = rs.Filter("RechargeTime__gte", begin).Filter("RechargeTime__lte", endTime)
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserId", id)
	}
	rs = rs.Filter("Finished", 1)
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&rechargeLogs)
	}
	return
}
func (this *AdminMgr) GetPlayLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, playLogs []models.PlayLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.PlayLog))
	var id string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			begin := map_params["begin"]
			end := map_params["end"]
			endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end+" 23:59:59", time.Local)
			rs = rs.Filter("CreateTime__gte", begin).Filter("CreateTime__lte", endTime)
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserId", id)
	}
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&playLogs)
	}
	return
}
