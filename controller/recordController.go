package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/db"
	"lol/entity"
	"lol/repository"
	"net/http"
	"strconv"
	"time"
)

func AddRecord(c *gin.Context) {

	var records []db.Record
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)

	err = json.Unmarshal(body, &records)
	common.DealErr(err)

	log.Println(records)
	for _, record := range records {
		log.Println(records)
		repository.AddRecord(&record)
	}
	c.JSON(http.StatusOK, records)
}

// JieSuan 结算游戏
func JieSuan(c *gin.Context) {

	var jies []Jie
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)

	err = json.Unmarshal(body, &jies)
	common.DealErr(err)

	var ids []string
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	for _, jy := range jies {
		ids = append(ids, strconv.FormatInt(jy.Hero, 10))
		var w int
		if jy.SubTotal > 0 {
			w = 1
		} else if jy.SubTotal == 0 {
			w = 0
		} else {
			w = -1
		}
		record := db.Record{
			UseHeroId:   jy.Hero,
			Score:       jy.Score,
			Win:         w,
			Subtotal:    jy.SubTotal,
			UseHeroName: getHeroNameById(repository.HeroList, jy.Hero),
			CreateTime:  currentTime,
		}
		repository.AddRecord(&record)
	}
	repository.DisableHero(ids)
	c.JSON(http.StatusOK, jies)
}

// 根据英雄id获取英雄名称
func getHeroNameById(list []entity.Hero, heroId int64) string {
	for _, one := range list {
		if one.Id == heroId {
			return one.HeroName
		}
	}
	return ""
}
