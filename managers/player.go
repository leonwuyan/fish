package managers

import (
	"fish/models"
	"github.com/astaxie/beego/orm"
)

type PlayerMgr struct {
}

var PlayerInstance = newPlayer()

func newPlayer() *PlayerMgr {
	return new(PlayerMgr)
}

func (this *PlayerMgr) GetPlayerById(id int) (player models.PlayerAccount, err error) {
	o := orm.NewOrm()
	o.QueryTable(new(models.PlayerAccount)).Filter("UserId", id).One(&player)
	return
}
func (this *PlayerMgr) GetPlayersByIds(ids []int) {

}
