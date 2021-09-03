package repo_record

import (
	"lol/common"
	"lol/entity"
	"lol/expection"
	"lol/sys_sb"
	"strings"
)

var DB = sys_sb.DB

func GetRecordsByMatchIds(ids []int64) []entity.Record {
	var recordList []entity.Record
	db := DB.Where(`match_id in ( ? )`, ids).Find(&recordList)
	common.DealDbErrs(db)
	return recordList
}

func Insert(record entity.Record) entity.Record {
	err := DB.Save(&record).Error
	if err != nil {
		s := err.Error()
		contains := strings.Contains(s, "Error 1062: Duplicate entry")
		if contains {
			panic(expection.BizErr{
				Code: 420,
				Msg:  "请勿重复插入记录",
			})
		}
	}
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

func GetRecordBYMatchIdAndUserId(matchId int64, userId int64) entity.Record {
	var record entity.Record
	_ = DB.Where("match_id = ? and user_id = ?", matchId, userId).Find(&record).RecordNotFound()
	return record
}
