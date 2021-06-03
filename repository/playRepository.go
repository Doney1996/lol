package repository

import (
	"fmt"
	"lol/common"
	"lol/db"
	"lol/entity"
)

func AddPlayer(player *entity.Player) {
	var sql = `insert into player (name,game_name,username,password) value (?,?,?,?)`
	exec, err := db.Db.Exec(
		sql,
		player.Name,
		player.GameName,
		player.Username,
		player.Password)
	common.DealErr(err)

	affected, err := exec.RowsAffected()
	common.DealErr(err)
	if affected != 1 {
		panic(fmt.Sprintf("保存用户: %v 失败", player))
	}
}
func LoginCheck(username string, password string) (bool, *entity.Player, error) {
	var sql = `select * from player where username = ? and password = ?`
	var player entity.Player
	row := db.Db.QueryRowx(sql, username, password)
	err := row.StructScan(&player)
	return player.Id != 0, &player, err
}
