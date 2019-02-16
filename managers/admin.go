package managers

import (
	"encoding/json"
	"fish/enums"
	"fish/models"
	"fish/payment"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
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
func (this *AdminMgr) GetAgentApply(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, cashLogs []models.AgentApply, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentApply))
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
func (this *AdminMgr) RechargePlayer(id int, amount float64, rechargeType int) enums.ReturnCode {
	_, err := PlayerInstance.GetPlayerById(id)
	if err != nil {
		return enums.PLAYER_NOT_FOUND
	}
	payOrder := payment.Create_order()
	SystemInstance.PreRecharge(enums.PAY_CHANNEL_SYSTEM, id, 105, rechargeType, int(amount*100), payOrder)
	SystemInstance.FinishRecharge(payOrder)
	return enums.SUCCESS
}
func (this *AdminMgr) UpdateChatMessage(id int) enums.ReturnCode {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.ChatMessages)).Filter("Id", id).Update(orm.Params{"IsProcessed": true})
	if err != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) AddCustomServiceMsg(account models.AdminAccount, userId int, content string) enums.ReturnCode {
	o := orm.NewOrm()
	data := models.ChatMessages{
		CreationTime: time.Now(),
		UserId:       userId,
		MessageType:  1,
		Message:      content,
		KefuId:       account.Id,
	}
	o.Insert(&data)
	ServerInstance.SendMsg(userId)
	return enums.SUCCESS
}

func (this *AdminMgr) GetChatMessages(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, cashLogs []models.ChatMessages, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.ChatMessages))
	var id, msgType string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			msgType = map_params["msg_type"]
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserId", id)
	}
	if len(msgType) > 0 {
		rs = rs.Filter("MessageType", msgType)
	}
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&cashLogs)
	return
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
func (this *AdminMgr) StatisticsPlay(begin, end time.Time) (total int64, result []models.StatisticsPlay, err error) {
	o := orm.NewOrm()
	end = end.AddDate(0, 0, 1).Add(-1)
	o.Raw("SELECT COUNT(DISTINCT UserId) FROM playlog WHERE CreateTime BETWEEN ? AND ?").SetArgs(begin, end).QueryRow(&total)
	o.Raw("SELECT GameId,COUNT(DISTINCT GameLogId) as PlayTimes,COUNT(DISTINCT UserId) as PlayPlayers,SUM(GoldChange) as WinOrLose from playlog where CreateTime BETWEEN ? AND ? GROUP BY GameId").SetArgs(begin, end).QueryRows(&result)
	return
}

func (this *AdminMgr) GetRechargeLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, rechargeLogs []models.RechargeLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.RechargeLog))
	var id, agent_id, rechargeType, finished string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			agent_id = map_params["agent_id"]
			finished = map_params["finished"]
			rechargeType = map_params["recharge_type"]
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
	if len(agent_id) > 0 {
		rs = rs.Filter("AgentId", agent_id)
	}
	if rechargeType != "0" {
		if rechargeType == "100" {
			rs = rs.Filter("RechargeType__lt", rechargeType)
		} else {
			rs = rs.Filter("RechargeType", rechargeType)
		}
	}
	rs = rs.Filter("Finished", finished)
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&rechargeLogs)
	}
	return
}
func (this *AdminMgr) GetPlayLogs(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, playLogs []models.PlayLog, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.PlayLog))
	var id, gameType string
	if len(searchParams) > 0 {
		var map_params map[string]string
		if err = json.Unmarshal([]byte(searchParams), &map_params); err == nil {
			id = map_params["id"]
			gameType = map_params["game_type"]
			begin := map_params["begin"]
			end := map_params["end"]
			beginTime, _ := time.ParseInLocation("2006-01-02T15:04", begin, time.Local)
			endTime, _ := time.ParseInLocation("2006-01-02T15:04", end, time.Local)
			rs = rs.Filter("CreateTime__gte", beginTime).Filter("CreateTime__lte", endTime)
		} else {
			logs.Error(err)
			return
		}
	}
	if len(id) > 0 {
		rs = rs.Filter("UserId", id)
	}
	if gameType != "0" {
		rs = rs.Filter("GameId", gameType)
	}
	total, _ = rs.Count()
	if total > 0 {
		_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).OrderBy("-Id").All(&playLogs)
	}
	return
}

func (this *AdminMgr) GetSystemWin(begin, end time.Time) (sysWin, sysFee int) {
	o := orm.NewOrm()
	var playResult, pumpResutl []orm.Params
	o.Raw("SELECT GameId,SUM(GoldChange) AS Gold FROM v_playlog WHERE UserId<10000000 AND CreateTime BETWEEN ? AND ? GROUP BY GameId;", begin, end).Values(&playResult)
	o.Raw("SELECT GameId,SUM(Pumping) AS Gold FROM v_pumplogdetail WHERE UserId<10000000 AND CreateTime BETWEEN ? AND ? GROUP BY GameId;", begin, end).Values(&pumpResutl)
	var catchFishWin, goldenFlowerWin, douDiZuWin, niuNiuMWin, kingQueenWin, niuNiuZWin, dragonTigerWin int
	var catchFishFee, goldenFlowerFee, douDiZuFee, niuNiuMFee, kingQueenFee, niuNiuZFee, dragonTigerFee int
	for _, item := range playResult {
		gameId, _ := strconv.Atoi(item["GameId"].(string))
		switch enums.GameType(gameId) {
		case enums.GAME_CATCH_FISH:
			catchFishWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_GOLDEN_FLOWER:
			goldenFlowerWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DOUDIZHU:
			douDiZuWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_M:
			niuNiuMWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_KING_QUEE:
			kingQueenWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_Z:
			niuNiuZWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DRAGON_TIGER:
			dragonTigerWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		}
	}
	for _, item := range pumpResutl {
		gameId, _ := strconv.Atoi(item["GameId"].(string))
		switch enums.GameType(gameId) {
		case enums.GAME_CATCH_FISH:
			catchFishFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_GOLDEN_FLOWER:
			goldenFlowerFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DOUDIZHU:
			douDiZuFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_M:
			niuNiuMFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_KING_QUEE:
			kingQueenFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_Z:
			niuNiuZFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DRAGON_TIGER:
			dragonTigerFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		}
	}
	goldenFlowerWin += goldenFlowerFee
	douDiZuWin += douDiZuFee
	niuNiuMWin += niuNiuMFee
	kingQueenWin += kingQueenFee
	niuNiuZWin += niuNiuZFee
	dragonTigerWin += dragonTigerFee
	sysWin = -(catchFishWin + goldenFlowerWin + douDiZuWin + niuNiuMWin + kingQueenWin + niuNiuZWin + dragonTigerWin)
	//logs.Info(fmt.Sprintf("noAgent#catchFishWin:%d,goldenFlowerWin:%d,douDiZuWin:%d,niuNiuMWin:%d,kingQueenWin:%d,niuNiuZWin:%d,dragonTigerWin:%d",catchFishWin,goldenFlowerWin,douDiZuWin,niuNiuMWin,kingQueenWin,niuNiuZWin,dragonTigerWin))
	sysFee = catchFishFee + goldenFlowerFee + douDiZuFee + niuNiuMFee + kingQueenFee + niuNiuZFee + dragonTigerFee
	//logs.Info(fmt.Sprintf("noAgent#catchFishFee:%d,goldenFlowerFee:%d,douDiZuFee:%d,niuNiuMFee:%d,kingQueenFee:%d,niuNiuZFee:%d,dragonTigerFee:%d",catchFishFee,goldenFlowerFee,douDiZuFee,niuNiuMFee,kingQueenFee,niuNiuZFee,dragonTigerFee))
	return
}
func (this *AdminMgr) GetHasAgentSystemWin(begin, end time.Time) (sysWin, sysFee int) {
	o := orm.NewOrm()
	var playResult, pumpResutl []orm.Params
	o.Raw("SELECT GameId,SUM(GoldChange) AS Gold FROM v_playlog WHERE UserId<10000000 AND AgentId>0 AND CreateTime BETWEEN ? AND ? GROUP BY GameId;", begin, end).Values(&playResult)
	o.Raw("SELECT GameId,SUM(Pumping) AS Gold FROM v_pumplogdetail WHERE UserId<10000000 AND AgentId>0 AND CreateTime BETWEEN ? AND ? GROUP BY GameId;", begin, end).Values(&pumpResutl)
	var catchFishWin, goldenFlowerWin, douDiZuWin, niuNiuMWin, kingQueenWin, niuNiuZWin, dragonTigerWin int
	var catchFishFee, goldenFlowerFee, douDiZuFee, niuNiuMFee, kingQueenFee, niuNiuZFee, dragonTigerFee int
	for _, item := range playResult {
		gameId, _ := strconv.Atoi(item["GameId"].(string))
		switch enums.GameType(gameId) {
		case enums.GAME_CATCH_FISH:
			catchFishWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_GOLDEN_FLOWER:
			goldenFlowerWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DOUDIZHU:
			douDiZuWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_M:
			niuNiuMWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_KING_QUEE:
			kingQueenWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_Z:
			niuNiuZWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DRAGON_TIGER:
			dragonTigerWin, _ = strconv.Atoi(item["Gold"].(string))
			break
		}
	}
	for _, item := range pumpResutl {
		gameId, _ := strconv.Atoi(item["GameId"].(string))
		switch enums.GameType(gameId) {
		case enums.GAME_CATCH_FISH:
			catchFishFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_GOLDEN_FLOWER:
			goldenFlowerFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DOUDIZHU:
			douDiZuFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_M:
			niuNiuMFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_KING_QUEE:
			kingQueenFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_NIUNIU_Z:
			niuNiuZFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		case enums.GAME_DRAGON_TIGER:
			dragonTigerFee, _ = strconv.Atoi(item["Gold"].(string))
			break
		}
	}
	goldenFlowerWin += goldenFlowerFee
	douDiZuWin += douDiZuFee
	niuNiuMWin += niuNiuMFee
	kingQueenWin += kingQueenFee
	niuNiuZWin += niuNiuZFee
	dragonTigerWin += dragonTigerFee
	sysWin = -(catchFishWin + goldenFlowerWin + douDiZuWin + niuNiuMWin + kingQueenWin + niuNiuZWin + dragonTigerWin)
	//logs.Info(fmt.Sprintf("hasAgent#catchFishWin:%d,goldenFlowerWin:%d,douDiZuWin:%d,niuNiuMWin:%d,kingQueenWin:%d,niuNiuZWin:%d,dragonTigerWin:%d",catchFishWin,goldenFlowerWin,douDiZuWin,niuNiuMWin,kingQueenWin,niuNiuZWin,dragonTigerWin))
	sysFee = catchFishFee + goldenFlowerFee + douDiZuFee + niuNiuMFee + kingQueenFee + niuNiuZFee + dragonTigerFee
	//logs.Info(fmt.Sprintf("hasAgent#catchFishFee:%d,goldenFlowerFee:%d,douDiZuFee:%d,niuNiuMFee:%d,kingQueenFee:%d,niuNiuZFee:%d,dragonTigerFee:%d",catchFishFee,goldenFlowerFee,douDiZuFee,niuNiuMFee,kingQueenFee,niuNiuZFee,dragonTigerFee))
	return
}
func (this *AdminMgr) GetAgentFee(begin, end time.Time) (fee float64) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(tax) FROM (SELECT log_id,tax FROM agent_fee_log WHERE log_time BETWEEN ? AND ? GROUP BY log_id) AS T0;", begin, end).QueryRow(&fee)
	return
}
func (this *AdminMgr) GetAgentTax(begin, end time.Time) (tax float64) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(fee) FROM agent_fee_log WHERE log_time BETWEEN ? AND ?", begin, end).QueryRow(&tax)
	return
}
func (this *AdminMgr) GetSystemAward(begin, end time.Time) (award int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(GoldChange) FROM goldchangelog WHERE ChangeType=9 AND ChangeTime BETWEEN ? AND ?", begin, end).QueryRow(&award)
	return
}
func (this *AdminMgr) GetPlayerRecharge(begin, end time.Time) (gold int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(GoldChange) FROM rechargelog WHERE Finished=1 AND RechargeType<100 AND RechargeTime BETWEEN ? AND ?", begin, end).QueryRow(&gold)
	return
}
func (this *AdminMgr) GetPlayerRechargeAward(begin, end time.Time) (gold int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(GoldChange) FROM rechargelog WHERE Finished=1 AND RechargeType=101 AND RechargeTime BETWEEN ? AND ?", begin, end).QueryRow(&gold)
	return
}
func (this *AdminMgr) GetPlayerRechargeForTest(begin, end time.Time) (gold int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(GoldChange) FROM rechargelog WHERE Finished=1 AND RechargeType=102 AND RechargeTime BETWEEN ? AND ?", begin, end).QueryRow(&gold)
	return
}
func (this *AdminMgr) GetPlayerCash(begin, end time.Time) (gold int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(Gold) FROM withdrawalslog WHERE State=3 AND WithdrawalsLogTime BETWEEN ? AND ?", begin, end).QueryRow(&gold)
	return
}
func (this *AdminMgr) GetSystemPunish(begin, end time.Time) (punish int) {
	o := orm.NewOrm()
	o.Raw("SELECT -SUM(GoldChange) FROM goldchangelog WHERE ChangeType=99 AND ChangeTime BETWEEN ? AND ?", begin, end).QueryRow(&punish)
	return
}
func (this *AdminMgr) GetPlayerRemain() (gold int) {
	o := orm.NewOrm()
	o.Raw("SELECT SUM(GlobalNum)+SUM(BankNum) FROM v_playeraccount").QueryRow(&gold)
	return
}

func (this *AdminMgr) IsNewMessage() (hasCashApply, hasChat int64) {
	o := orm.NewOrm()
	hasCashApply, _ = o.QueryTable(new(models.PlayerCashLog)).Filter("State", 0).Count()
	hasChat, _ = o.QueryTable(new(models.ChatMessages)).Filter("IsProcessed", false).Filter("IsUserMessage", true).Count()
	return
}

func (this *AdminMgr) GetAllChannels(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, channels []models.Channel, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.Channel))
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).All(&channels)
	return
}
func (this *AdminMgr) GetChannelInfo(id int) (channel models.Channel, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(channel).Filter("Id", id).One(&channel)
	return
}
func (this *AdminMgr) ChangeChannelInfo(data models.Channel) enums.ReturnCode {
	if ServerInstance.ChangeChannelInfo(data) != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) GetAllShowAgents(account models.AdminAccount, pageSize, pageIndex int, searchParams string) (total int64, agents []models.AgentShow, err error) {
	o := orm.NewOrm()
	rs := o.QueryTable(new(models.AgentShow))
	total, _ = rs.Count()
	_, err = rs.Limit(pageSize, (pageIndex-1)*pageSize).All(&agents)
	return
}
func (this *AdminMgr) GetShowAgentInfo(id int) (agent models.AgentShow, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(agent).Filter("AgentId", id).One(&agent)
	return
}
func (this *AdminMgr) ChangeShowAgentInfo(data models.AgentShow) enums.ReturnCode {
	if ServerInstance.ChangeShowAgent(data) != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
func (this *AdminMgr) GetNotice() {

}
func (this *AdminMgr) ChangeNotice(data models.Notice) enums.ReturnCode {
	if ServerInstance.ChangeNotice(data) != nil {
		return enums.DB_ACTION_ERROR
	}
	return enums.SUCCESS
}
