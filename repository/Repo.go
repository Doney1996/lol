package repository

import (
	"lol/entity"
	"lol/repository/hero"
)

var HeroList []entity.Hero

func init() {
	//查询出所有的英雄记录放在内存
	HeroList = hero.GetAllHero()
}
