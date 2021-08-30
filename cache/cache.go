package cache

import (
	"lol/entity"
	"lol/repository/repo_hero"
	"lol/repository/repo_play"
)

var HeroList []entity.Hero
var PlayList []entity.Player

func init() {
	//查询出所有的英雄记录放在内存
	HeroList = repo_hero.GetAllHero()
	PlayList = repo_play.GetAllHero()
}

// GetHeroNameById 根据英雄id获取名字
func GetHeroNameById(id int64) string {
	for _, hero := range HeroList {
		if hero.Id == id {
			return hero.HeroName
		}
	}
	return ""
}

// GetPlayNameById 根据玩家id获取玩家游戏名称
func GetPlayNameById(id int64) string {
	for _, play := range PlayList {
		if play.Id == id {
			return play.GameName
		}
	}
	return ""
}
