package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/db"
	"net/http"
)

func DisableHero(c *gin.Context) {
	var ids []string
	_ = c.BindJSON(&ids)
	db.DisableHero(ids)
	c.JSON(http.StatusOK, "success")
}

func EnableAllHero(c *gin.Context) {
	db.EnableAllHero()
	c.JSON(http.StatusOK, "success")
}

func AddHero(c *gin.Context) {
	var heros []db.Hero
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)

	err = json.Unmarshal(body, &heros)
	common.DealErr(err)

	log.Println(heros)
	for _, record := range heros {
		log.Println(heros)
		db.AddHero(&record)
	}
	c.JSON(http.StatusOK, heros)
}
func GetAllHero(c *gin.Context) {
	list := db.GetHeroList()
	c.JSON(http.StatusOK, list)
}

type Jie struct {
	Score    float64
	Hero     int64
	SubTotal float64
}

