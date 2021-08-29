package repository

import (
	"lol/entity"
)

func AddRecord(record *entity.Record) {
	var _ = `INSERT INTO recording 
		( play_id, player_id, player_name, use_hero_id, use_hero_name, win, score, unit_price, Subtotal,create_time,sub_score) 
		VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?);`
}
