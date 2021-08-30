package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/entity"
	"lol/repository/repo_hero"
	"net/http"
)

// DisableHero 关闭指定英雄
func DisableHero(c *gin.Context) {
	var ids []string
	_ = c.BindJSON(&ids)
	repo_hero.DisableHero(ids)
	c.JSON(http.StatusOK, "success")
}

// EnableAllHero 关闭所有英雄
func EnableAllHero(c *gin.Context) {
	repo_hero.EnableAllHero()
	c.JSON(http.StatusOK, "success")
}

// EnableAllHeroById 根据id激活英雄
func EnableAllHeroById(c *gin.Context) {
	repo_hero.EnableAllHero()
	c.JSON(http.StatusOK, "success")
}

func AddHero(c *gin.Context) {
	var heroes []entity.Hero
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)

	err = json.Unmarshal(body, &heroes)
	common.DealErr(err)

	log.Println(heroes)
	for _, record := range heroes {
		log.Println(heroes)
		repo_hero.AddHero(&record)
	}
	c.JSON(http.StatusOK, heroes)
}
func GetAllHeroTop(c *gin.Context) {
	list := repo_hero.GetHeroTopList()
	c.JSON(http.StatusOK, list)
}

func GetAllHeroInfoList(c *gin.Context) {
	list := repo_hero.GetAllHeroListInfo()
	c.JSON(http.StatusOK, list)
}

type Jie struct {
	Score    float64
	Hero     int64
	SubTotal float64
}
