package hero

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var DB = db.DB

func DisableHero(ids []string) {

}
func EnableAllHero() {

}
func AddHero(hero *entity.Hero) {

}

func GetAllHero() []entity.Hero {
	var heroes []entity.Hero
	db := DB.Select(&heroes, "select * from hero")
	common.DealDbErrs(db)
	return heroes
}

type showHero struct {
	*entity.Hero
	Count int64 `json:"count" db:"count"`
	Sum   int64 `json:"sum" db:"sum"`
}

func GetAllHeroListInfo() []entity.Hero {
	var sql = `select id,hero_name,hero_other_name,sort,disable	,tier,	img_position,	position from hero`
	var list []entity.Hero
	db := db.DB.Select(&list, sql)
	common.DealDbErrs(db)
	return list
}

// GetHeroTopList 获取排行榜
func GetHeroTopList() []showHero {
	var sql = `select 
       id, hero_name, hero_other_name,
       sort, disable, tier, img_position,
       position, count, sum
from (
         select hero.*, IFNULL(top.count, 0) as count, IFNULL(top.sum, 0) as sum
         from hero
                  left join (
             select use_hero_id, SUM(Subtotal) as sum, count(1) as count
             from recording
             group by use_hero_id) as top on hero.id = top.use_hero_id) as t

order by t.sum desc;`
	var list []showHero
	db := db.DB.Select(&list, sql)
	common.DealDbErrs(db)
	return list
}
