package season

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var DB = db.DB

// GetCurrentSeason 根据赛季类型查看当前赛季
func GetCurrentSeason(seasonType string, lifeStatus string) entity.Season {
	//var sql = `select id, season_name, life_status, create_time ,season_type
	//			from season where season_type = ? and life_status = ? order by id desc limit 1 `
	//var season []entity.Season
	//DB.Select(&season, sql, seasonType, lifeStatus)
	//if len(season) > 0 {
	//	return season[0]
	//}else {
	//	return entity.Season{}
	//}
	var season entity.Season
	db := DB.Where("season_type = ? and life_status = ? ", seasonType, lifeStatus).Last(&season)
	if db.RecordNotFound() {
		return season
	}
	common.DealDbErrs(db)
	return season
}

// InsertSeason 新增一个赛季
func InsertSeason(season *entity.Season) *entity.Season {
	db := DB.Save(season)
	common.DealDbErrs(db)
	return season
}

func CloseSeason(id int64) {
	db := DB.Model(&entity.Season{}).Update("life_status", "1").Where("id", id)
	common.DealDbErrs(db)
}
