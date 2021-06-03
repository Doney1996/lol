package repository

import (
	"lol/common"
	"lol/db"
)

func AddRecord(record *db.Record) {
	var sql = `INSERT INTO recording 
		( play_id, player_id, player_name, use_hero_id, use_hero_name, win, score, unit_price, Subtotal,create_time,sub_score) 
		VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?);`
	_, err := db.Db.Exec(sql,
		record.PlayId,
		record.PlayerId,
		record.PlayerName,
		record.UseHeroId,
		record.UseHeroName,
		record.Win,
		record.Score,
		record.UnitPrice,
		record.Subtotal,
		record.CreateTime,
		record.Subtotal/10,
	)
	common.DealErr(err)
}
