package managers

import (
	"encoding/json"
	"fish/enums"
	"fish/models"
	"fish/payment"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"time"
)

type AgentMgr struct {
}

var AgentInstance = newAgent()

func newAgent() *AgentMgr {
	return new(AgentMgr)
}
func (this *AgentMgr) GetAgentInfo(agent models.AgentAccount) (result map[string]string) {
	var today_reg_player, today_reg_player_child, yesterday_reg_player, yesterday_reg_player_child, total_reg_player, total_reg_player_child int64
	o := orm.NewOrm()
	today_reg_player, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId", agent.Id).Filter("RsgLogTime__gte", time.Today()).Count()
	yesterday_reg_player, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId", agent.Id).Filter("RsgLogTime__gte", time.Yesterday()).Filter("RsgLogTime__lt", time.Today()).Count()
	total_reg_player, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId", agent.Id).Count()
	children := this.findAllChildrenId(agent.Id)
	if len(children) > 0 {
		today_reg_player_child, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId__in", children).Filter("RsgLogTime__gte", time.Today()).Count()
		yesterday_reg_player_child, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId__in", children).Filter("RsgLogTime__gte", time.Yesterday()).Filter("RsgLogTime__lt", time.Today()).Count()
		total_reg_player_child, _ = o.QueryTable(new(models.PlayerAccount)).Filter("AgentId__in", children).Count()
	} else {
		today_reg_player_child = 0
		yesterday_reg_player_child = 0
		total_reg_player_child = 0
	}
	var today_tax, yesterday_tax float64
	o.Raw("SELECT SUM(fee) FROM agent_fee_log WHERE agent_id=? AND log_time>=?", agent.Id, time.Today()).QueryRow(&today_tax)
	o.Raw("SELECT SUM(fee) FROM agent_fee_log WHERE agent_id=? AND log_time>=? AND log_time<?", agent.Id, time.Yesterday(), time.Today()).QueryRow(&yesterday_tax)
	//获取是刷新一下
	agent, _ = this.GetAgentById(agent.Id)
	result = make(map[string]string)
	result["rate"] = strconv.Itoa(agent.Rate)
	result["gold"] = strconv.FormatFloat(agent.Gold/100.0, 'f', 2, 64)
	result["today_reg_player"] = strconv.Itoa(int(today_reg_player))
	result["today_reg_player_child"] = strconv.Itoa(int(today_reg_player_child))
	result["yesterday_reg_player"] = strconv.Itoa(int(yesterday_reg_player))
	result["yesterday_reg_player_child"] = strconv.Itoa(int(yesterday_reg_player_child))
	result["total_reg_player"] = strconv.Itoa(int(total_reg_player))
	result["total_reg_player_child"] = strconv.Itoa(int(total_reg_player_child))
	result["today_tax"] = strconv.FormatFloat(today_tax/100, 'f', 4, 64)
	result["yesterday_tax"] = strconv.FormatFloat(yesterday_tax/100, 'f', 4, 64)
	result["total_tax"] = strconv.FormatFloat(agent.TotalFee/100, 'f', 4, 64)
	result["children"] = strconv.Itoa(agent.TotalChildrenImmediate)
	result["children_children"] = strconv.Itoa(agent.TotalChildrenOther)
	return
}
func (this *AgentMgr) GetAllAgents() (agents []models.AgentAccount, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(models.AgentAccount)).All(&agents)
	return
}
func (this *AgentMgr) GetAgentById(id int) (agent models.AgentAccount, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(agent).Filter("Id", id).One(&agent)
	return
}
func (this *AgentMgr) GetAgentByName(name string) (agent models.AgentAccount, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(agent).Filter("Name", name).One(&agent)
	return
}
func (this *AgentMgr) Register(agent models.AgentAccount) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.Insert(&agent); err != nil {
		return enums.DB_ACTION_ERROR
	}
	this.statistic_child(agent, true)
	return enums.SUCCESS
}
func (this *AgentMgr) ChangePwd(agent models.AgentAccount, oldPwd, newPwd string) enums.ReturnCode {
	if agent.Password != oldPwd {
		return enums.AGENT_OLDPASSWORD_ERROR
	}
	o := orm.NewOrm()
	agent.Password = newPwd
	if _, err := o.Update(&agent, "Password"); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AgentMgr) GetChildren(agent models.AgentAccount, pageSize, pageIndex int, searchParams string) (total int64, agents []models.AgentAccount, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentAccount)).Filter("ParentId", agent.Id)
	var id, name, mobile, start, end string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			name = map_params["name"]
			mobile = map_params["mobile"]
			start = map_params["start"]
			end = map_params["end"]
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
	if len(start) > 0 || len(end) > 0 {
		rs = rs.Filter("RegisterTime__gte", start).Filter("RegisterTime_lt", end)
	}
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).All(&agents)
	return
}
func (this *AgentMgr) GetPlayerAgent(playerId int) (agentId int) {
	o := orm.NewOrm()
	var user2Agent models.User2Agent
	if err := o.QueryTable(user2Agent).Filter("UserId", playerId).One(&user2Agent); err != nil {
		return 0
	}
	return user2Agent.QmAgentId
}
func (this *AgentMgr) GetPlayers(agent models.AgentAccount, pageSize, pageIndex int, searchParams string) (total int64, players []models.PlayerAccount, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.PlayerAccount)).Filter("AgentId", agent.Id)
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
func (this *AgentMgr) GetTax(agent models.AgentAccount, pageSize, pageIndex int, searchParams string) (total int64, feeLogs []models.AgentFeeLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentFeeLog)).Filter("AgentId", agent.Id)
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&feeLogs)
	}
	return
}
func (this *AgentMgr) GetCashLog(agent models.AgentAccount, pageSize, pageIndex int, searchParams string) (total int64, feeLogs []models.AgentCashLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentCashLog)).Filter("AgentId", agent.Id)
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
		rs = rs.Filter("Id", id)
	}
	if state != "-1" {
		rs = rs.Filter("State", state)
	}
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&feeLogs)
	}
	return
}
func (this *AgentMgr) GetBankInfos(account models.AgentAccount) (infos []models.BankCardInfo, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(models.BankCardInfo)).Filter("AgentId", account.Id).All(&infos)
	return
}
func (this *AgentMgr) GetBankInfoById(id int) (info models.BankCardInfo) {
	o := orm.NewOrm()
	o.QueryTable(info).Filter("Id", id).One(&info)
	return
}
func (this *AgentMgr) AddBankInfo(info models.BankCardInfo) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.Insert(&info); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AgentMgr) UpdateBankInfo(info models.BankCardInfo) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.Update(&info); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AgentMgr) CashApply(account models.AgentAccount, amount, bankInfoId int) enums.ReturnCode {
	changeGold := amount * 100
	bankInfo := this.GetBankInfoById(bankInfoId)
	o := orm.NewOrm()
	o.Begin()
	o.QueryTable(account).Filter("Id", account.Id).Update(orm.Params{"Gold": orm.ColValue(orm.ColAdd, -changeGold)})
	o.QueryTable(account).Filter("Id", account.Id).One(&account)
	if account.Gold < 0 {
		o.Rollback()
		return enums.AGENT_NOT_ENOUGH_GOLD
	}
	cashLog := models.AgentCashLog{}
	cashLog.AgentId = account.Id
	cashLog.Gold = changeGold
	cashLog.TxType = bankInfo.CashType
	cashLog.OrderId = payment.Create_order()
	if enums.CashType(bankInfo.CashType) == enums.CASH_TYPE_ALIPAY {
		cashLog.Alipay = bankInfo.BankCardNo
		cashLog.AlipayName = bankInfo.RealName
	} else {
		cashLog.BankCardTypeId = bankInfo.BankType
		cashLog.BankCardNo = bankInfo.BankCardNo
		cashLog.RealName = bankInfo.RealName
	}
	cashLog.WithdrawalsLogTime = time.Now()
	if _, err := o.Insert(&cashLog); err != nil {
		o.Rollback()
		return enums.DB_ACTION_ERROR
	}
	o.Commit()
	return enums.SUCCESS
}

func (this *AgentMgr) ChangeChildRate(account models.AgentAccount, id, rate int) enums.ReturnCode {
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.AgentAccount)).Filter("Id", id).Update(orm.Params{"Rate": rate}); err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}

func (this *AgentMgr) FeeToAgent(playLog models.PlayLog) (err error) {
	var feeRate = 0
	switch enums.GameType(playLog.GameId) {
	default:
	case enums.GAME_CATCH_FISH:
		feeRate = 3
		break
	case enums.GAME_GOLDEN_FLOWER, enums.GAME_DOUDIZHU, enums.GAME_NIUNIU_M, enums.GAME_KING_QUEE, enums.GAME_NIUNIU_Z, enums.GAME_DRAGON_TIGER:
		feeRate = 5
		break
	}
	if feeRate > 0 {
		agentId := this.GetPlayerAgent(playLog.UserId)
		if agentId > 0 {
			if agent, err := this.GetAgentById(agentId); err == nil {
				this.statistic_fee(playLog, "", agent, models.AgentAccount{}, int64(feeRate))
			}
		}
	}
	return nil
}
func (this *AgentMgr) statistic_fee(playLog models.PlayLog, playerName string, agent, child models.AgentAccount, gameFeeRate int64) {
	var realGoldChange int64 //玩家真实输赢
	if playLog.GoldChange > 0 {
		realGoldChange = playLog.GoldChange * 100 / (100 - gameFeeRate)
	} else {
		realGoldChange = playLog.GoldChange
	}
	//总税收
	tax := float64(realGoldChange) * float64(gameFeeRate) / 100 / 2 //输赢各一半
	//真实税收
	realTax := math.Abs(tax)
	//代理应得比例
	agentFeeRate := agent.Rate - child.Rate
	//玩家输赢按代理权重计算
	winLose := float64(realGoldChange) * float64(agentFeeRate) / 100
	//代理收益
	fee := realTax * float64(agentFeeRate) / 100
	logData := models.AgentFeeLog{
		LogId:      playLog.Id,
		UserId:     playLog.UserId,
		UserName:   playerName,
		AgentId:    agent.Id,
		AgentChild: child.Id,
		GameId:     playLog.GameId,
		RoomId:     playLog.RoomId,
		Tax:        realTax,
		Rate:       agent.Rate,
		ChildRate:  child.Rate,
		Fee:        fee,
		WinLose:    winLose,
		LogTime:    playLog.CreateTime,
	}
	o := orm.NewOrm()
	o.Begin()
	if winLose > 0 {
		if _, err := o.QueryTable(agent).Filter("Id", agent.Id).Update(orm.Params{
			"Gold":      orm.ColValue(orm.ColAdd, fee),
			"TotalTax":  orm.ColValue(orm.ColAdd, realTax),
			"TotalFee":  orm.ColValue(orm.ColAdd, fee),
			"PlayerWin": orm.ColValue(orm.ColAdd, winLose),
		}); err != nil {
			o.Rollback()
			logs.Error(err)
		}
	} else {
		if _, err := o.QueryTable(agent).Filter("Id", agent.Id).Update(orm.Params{
			"Gold":       orm.ColValue(orm.ColAdd, fee),
			"TotalTax":   orm.ColValue(orm.ColAdd, realTax),
			"TotalFee":   orm.ColValue(orm.ColAdd, fee),
			"PlayerLose": orm.ColValue(orm.ColAdd, winLose),
		}); err != nil {
			o.Rollback()
			logs.Error(err)
		}
	}
	if _, err := o.Insert(&logData); err != nil {
		o.Rollback()
		logs.Error(err)
	}
	o.Commit()
	if agent.ParentId > 0 {
		parent, err := this.GetAgentById(agent.ParentId)
		if err != nil {
			return
		}
		this.statistic_fee(playLog, playerName, parent, agent, gameFeeRate)
	}
}
func (this *AgentMgr) Statistic_player(agentId int, isImmediate bool) (err error) {
	o := orm.NewOrm()
	if agent, err := this.GetAgentById(agentId); err == nil {
		if isImmediate {
			o.QueryTable(agent).Filter("Id", agent.Id).Update(orm.Params{"TotalPlayersImmediate": orm.ColValue(orm.ColAdd, 1)})
		} else {
			o.QueryTable(agent).Filter("Id", agent.Id).Update(orm.Params{"TotalPlayersOther": orm.ColValue(orm.ColAdd, 1)})
		}
		if agent.ParentId > 0 {
			this.Statistic_player(agent.ParentId, false)
		}
	}
	return
}
func (this *AgentMgr) statistic_child(account models.AgentAccount, isImmediate bool) {
	if account.ParentId > 0 {
		var parent models.AgentAccount
		o := orm.NewOrm()
		if isImmediate {
			o.QueryTable(account).Filter("Id", account.ParentId).Update(orm.Params{"TotalChildrenImmediate": orm.ColValue(orm.ColAdd, 1)})
		} else {
			o.QueryTable(account).Filter("Id", account.ParentId).Update(orm.Params{"TotalChildrenOther": orm.ColValue(orm.ColAdd, 1)})
		}
		if err := o.QueryTable(account).Filter("Id", account.ParentId).One(&parent); err == nil {
			this.statistic_child(parent, false)
		}
	}
}

func (this *AgentMgr) findAllChildrenId(id int) (ids []int) {
	children := this.findChildrenId(id)
	if len(children) > 0 {
		ids = append(ids, children...)
		for _, id := range ids {
			ids = append(ids, this.findAllChildrenId(id)...)
		}
	}
	return
}
func (this *AgentMgr) findChildrenId(id int) (ids []int) {
	o := orm.NewOrm()
	var children []models.AgentAccount
	o.QueryTable(new(models.AgentAccount)).Filter("ParentId", id).All(&children)
	for _, child := range children {
		ids = append(ids, child.Id)
	}
	return ids
}
