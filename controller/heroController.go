package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/entity"
	"lol/repository"
	"net/http"
)

// DisableHero 关闭指定英雄
func DisableHero(c *gin.Context) {
	var ids []string
	_ = c.BindJSON(&ids)
	repository.DisableHero(ids)
	c.JSON(http.StatusOK, "success")
}

// EnableAllHero 关闭所有英雄
func EnableAllHero(c *gin.Context) {
	repository.EnableAllHero()
	c.JSON(http.StatusOK, "success")
}

// EnableAllHeroById 根据id激活英雄
func EnableAllHeroById(c *gin.Context) {
	repository.EnableAllHero()
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
		repository.AddHero(&record)
	}
	c.JSON(http.StatusOK, heroes)
}
func GetAllHero(c *gin.Context) {
	list := repository.GetHeroList()
	c.JSON(http.StatusOK, list)
}

type Jie struct {
	Score    float64
	Hero     int64
	SubTotal float64
}
