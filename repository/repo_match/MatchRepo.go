package repo_match

import (
	"lol/common"
	"lol/entity"
	"lol/expection"
	"lol/sys_sb"
)

var DB = sys_sb.DB

// GetMatchBySeasonAndStatus 当前赛季的所有对局
func GetMatchBySeasonAndStatus(seasonId int64, lifeStatus string) []entity.Match {
	var matchList []entity.Match
	db := DB.Where(" season_id = ? and life_status = ? ", seasonId, lifeStatus).Find(&matchList)
	common.DealDbErrs(db)
	return matchList
}

// GetLastBySeasonAndStatus 获取最后一条记录
func GetLastBySeasonAndStatus(seasonId int64, lifeStatus string) entity.Match {
	var match entity.Match
	if DB.Where(" season_id = ? and life_status = ? ", seasonId, lifeStatus).Last(&match).RecordNotFound() {
		return match
	}
	return match
}

func Insert(match entity.Match) entity.Match {
	db := DB.Save(&match)
	common.DealDbErrs(db)
	return match
}

func Update(match entity.Match) {
	db := DB.Save(match)
	common.DealDbErrs(db)
}

func GetById(id int64) entity.Match {
	var match entity.Match
	db := DB.Where("id = ? ", id).Find(&match)
	common.DealDbErrs(db)
	return match
}

// GetLastBySeasonId 根据赛季类型获取最后一场
func GetLastBySeasonId(seasonId int64) entity.Match {
	var match entity.Match
	if DB.Where("season_id = ? and life_status = '0' ", seasonId).Find(&match).RecordNotFound() {
		panic(expection.BizErr{
			Code: 410,
			Msg:  "没有正在进行的对局，重新开始",
		})
	} else {
		return match
	}

}
