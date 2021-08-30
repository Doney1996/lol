package repo_season

import (
	"lol/common"
	"lol/entity"
	"lol/sys_sb"
)

var DB = sys_sb.DB

// GetCurrentSeason 根据赛季类型查看当前赛季
func GetCurrentSeason(seasonType string, lifeStatus string) entity.Season {
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

func GetById(id int64) entity.Season {
	var season entity.Season
	err := DB.Where("id = ?", id).Find(&season).Error
	common.DealErr(err)
	return season
}
