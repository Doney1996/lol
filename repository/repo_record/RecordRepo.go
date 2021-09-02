package repo_record

import (
	"lol/common"
	"lol/entity"
	"lol/sys_sb"
)

var DB = sys_sb.DB

func GetRecordsByMatchIds(ids []int64) []entity.Record {
	var recordList []entity.Record
	db := DB.Where(`match_id in ( ? )`, ids).Find(&recordList)
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
