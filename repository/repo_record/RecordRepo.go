package repo_record

import (
	"fmt"
	"lol/common"
	"lol/entity"
	"lol/sys_sb"
)

var DB = sys_sb.DB

func GetRecordsByMatchIds(ids []int64) []entity.Record {
	var sql = `select id, user_id, hero_id, match_id, is_win, score, unit_price, last_score, amount, update_time, create_time
			from record where  match_id in (?)`
	var recordList []entity.Record
	inStatus := ""
	params := make([]interface{}, 0)
	for i := 0; i < len(ids); i++ {
		if i == 0 {
			inStatus += "?"
		} else {
			inStatus += ",?"
		}
		params = append(params, ids[i])
	}
	sql = fmt.Sprintf(sql, inStatus)
	db := DB.Select(&recordList, sql, ids)
	common.DealDbErrs(db)
	return recordList
}

func Insert(record entity.Record) entity.Record {
	db := DB.Save(&record)
	common.DealDbErrs(db)
	return record
}

func Update(record entity.Record) {
	db := DB.Save(&record)
	common.DealDbErrs(db)
}

func GetByMatchId(matchId int64) []entity.Record {
	var records []entity.Record
	db := DB.Where("match_id = ?", matchId).Find(&records)
	common.DealDbErrs(db)
	return records
}
