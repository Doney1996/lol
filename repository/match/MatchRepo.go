package match

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var DB = db.DB

// GetCurrentSeasonMatch 当前赛季的所有对局
func GetCurrentSeasonMatch(seasonId int64) []entity.Match {
	var matchList []entity.Match
	var sql = `select * from game_match where season_id = ? and life_status = '1'`
	db := DB.Select(&matchList, sql, seasonId)
	common.DealDbErrs(db)
	return matchList
}

func Insert(match entity.Match) {
	//sql := `INSERT INTO game_match
	//	(season_id, life_status, create_time, update_time)
	//	VALUES (:season_id, :life_status, :create_time, :update_time) ;`
	db := DB.Save(match)
	common.DealDbErrs(db)
}

func Update(match entity.Match) {
	//sql :=`UPDATE game_match
	//SET season_id = :season_id, life_status = :life_status, create_time = :create_time, update_time = :update_time WHERE
	//id = :id;`
	//result, err := DB.NamedExec(sql, match)
	//affected, err := result.RowsAffected()
	//common.DealErr(err)
	//if affected == 0 {
	//	panic("更新match失败")
	//}
	db := DB.Save(match)
	common.DealDbErrs(db)
}

func GetById(id int64) entity.Match {
	var match entity.Match
	db := DB.Where("id = ? ", id).Find(&match)
	common.DealDbErrs(db)
	return match
}
