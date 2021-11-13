package db

import (
	"log"
	"lol/common"
)

func DisableHero(ids []string) {
	var sql = `update hero set  disable = '1' where id = ? `
	for _, id := range ids {
		_, err := Db.Exec(sql, id)
		common.DealErr(err)
	}
}
func EnableAllHero() {
	_, err := Db.Exec("update hero set  disable = '0' ")
	common.DealErr(err)
}
func AddHero(hero *Hero) {
	var sql = `INSERT INTO hero (hero_name, hero_other_name, sort, disable, tier, img_position, position)
				VALUES (?, ?, ?, ?, ?, ?, ?);`
	_, err := Db.Exec(sql,
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

func GetAllHero() []Hero {
	var heros []Hero
	err := Db.Select(&heros, "select * from hero")
	common.DealErr(err)
	return heros
}

func GetAllPlayer() []Player {
	var players []Player
	err := Db.Select(&players, "select * from player")
	common.DealErr(err)
	return players
}

func GetHeroList() []map[string]interface{} {
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
	query, err := Db.Query(sql)
	common.DealErr(err)

	list := make([]map[string]interface{}, 0)
	columns, _ := query.Columns()
	log.Println(columns)
	for query.Next() {
		one := make(map[string]interface{})
		var id int64
		var heroName string
		var heroOtherName string
		var sort int64
		var disable bool
		var tier int64
		var imgPosition int64
		var position int64
		var count int64
		var sum int64
		err = query.Scan(&id,
			&heroName,
			&heroOtherName,
			&sort,
			&disable,
			&tier,
			&imgPosition,
			&position,
			&count,
			&sum,
		)
		one["id"] = id
		one["heroName"] = heroName
		one["heroOtherName"] = heroOtherName
		one["sort"] = sort
		one["disable"] = disable
		one["tier"] = tier
		one["imgPosition"] = imgPosition
		one["position"] = position
		one["count"] = count
		one["sum"] = sum
		common.DealErr(err)

		list = append(list, one)
	}
	return list
}
