package record

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var DB = db.DB

func GetRecordByMatchIds(ids []int64) []entity.Record {
	var sql = `select id, user_id, hero_id, match_id, is_win, score, unit_price, last_score, amount, update_time, create_time
from record where match_id in (?)`
	var recordList []entity.Record
	//inStatus:=""
	//params:=make([]interface{},0)
	//for i:=0;i<len(ids);i++{
	//	if i==0{
	//		inStatus+="?"
	//	}else{
	//		inStatus+=",?"
	//	}
	//	params=append(params , ids[i])
	//}
	//sql = fmt.Sprintf(sql ,inStatus )
	//var recordList []entity.Record
	//DB.Select(&recordList,sql)
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
	db := DB.Save(record)
	common.DealDbErrs(db)
}

func GetByMatchId(matchId int64) []entity.Record {
	sql := `select * FROM record where match_id = ?`
	var records []entity.Record
	db := DB.Select(&records, sql, matchId)
	common.DealDbErrs(db)
	return records
}
