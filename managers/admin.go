package managers

import (
	"encoding/json"
	"fish/enums"
	"fish/models"
	"fish/payment"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AdminMgr struct {
}

var AdminInstance = newAdmin()

func newAdmin() *AdminMgr {
	return new(AdminMgr)
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
