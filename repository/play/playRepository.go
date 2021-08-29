package play

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var DB = db.DB

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
