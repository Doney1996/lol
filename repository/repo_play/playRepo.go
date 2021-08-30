package repo_play

import (
	"lol/common"
	"lol/entity"
	"lol/sys_sb"
)

var DB = sys_sb.DB

func AddPlayer(player *entity.Player) {
	db := DB.Save(player)
	common.DealDbErrs(db)
}
func LoginCheck(username string, password string) (bool, *entity.Player) {
	var player entity.Player
	db := DB.Where("username = ? and password = ?", username, password).Find(&player)
	common.DealDbErrs(db)
	return player.Id != 0, &player
}

// GetAllHero 所有的玩家
func GetAllHero() []entity.Player {
	var plays []entity.Player
	db := DB.Find(&plays)
	common.DealDbErrs(db)
	return plays
}
