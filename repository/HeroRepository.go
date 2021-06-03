package repository

import (
	"lol/common"
	"lol/db"
	"lol/entity"
)

var HeroList []entity.Hero

func DisableHero(ids []string) {
	var sql = `update hero set  disable = '1' where id = ? `
	for _, id := range ids {
		_, err := db.Db.Exec(sql, id)
		common.DealErr(err)
	}
}
func EnableAllHero() {
	_, err := db.Db.Exec("update hero set  disable = '0' ")
	common.DealErr(err)
}
func AddHero(hero *entity.Hero) {
	var sql = `INSERT INTO hero (hero_name, hero_other_name, sort, disable, tier, img_position, position)
				VALUES (?, ?, ?, ?, ?, ?, ?);`
	_, err := db.Db.Exec(sql,
		hero.HeroName,
		hero.HeroOtherName,
		hero.Sort,
		hero.Disable,
		hero.Tier,
		hero.ImgPosition,
		hero.Position,
	)
	common.DealErr(err)
}

func GetAllHero() []entity.Hero {
	var heroes []entity.Hero
	err := db.Db.Select(&heroes, "select * from hero")
	common.DealErr(err)
	return heroes
}

type showHero struct {
	Id            int64  `json:"id" db:"id"`
	HeroName      string `json:"heroName" db:"hero_name"`
	HeroOtherName string `json:"heroOtherName" db:"hero_other_name"`
	Sort          int64  `json:"sort" db:"sort"`
	Disable       bool   `json:"disable" db:"disable"`
	Tier          int64  `json:"tier" db:"tier"`
	ImgPosition   int64  `json:"imgPosition" db:"img_position"`
	Position      string `json:"position" db:"position"`
	Count         int64  `json:"count" db:"count"`
	Sum           int64  `json:"sum" db:"sum"`
}

func GetHeroList() []showHero {
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
	err := db.Db.Select(&list, sql)
	common.DealErr(err)
	return list
}

//查询出所有的英雄记录放在内存
func init() {
	HeroList = GetAllHero()
}
