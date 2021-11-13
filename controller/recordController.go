package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/db"
	"lol/rpc"
	"net/http"
	"strconv"
	"time"
)

// GetInfoFromWeGame 从WeGame的接口获取战绩
func GetInfoFromWeGame(c *gin.Context) {
	players := db.GetAllPlayer()
	var uids []int64
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)
	err = json.Unmarshal(body, &uids)
	common.DealErr(err)

	var openIds []*string
	for _, uid := range uids {
		for _, player := range players {
			if player.Id == uid {
				openIds = append(openIds, player.OpenId)
			}
		}
	}

	var result []rpc.Info
	for _, openId := range openIds {
		post := rpc.Post(*openId, "")
		result = append(result, post.Battles[0])
	}
	c.JSON(http.StatusOK, result)
}

func AddRecord(c *gin.Context) {

	var records []db.Record
	body, err := ioutil.ReadAll(c.Request.Body)
	common.DealErr(err)

	err = json.Unmarshal(body, &records)
	common.DealErr(err)

	for _, record := range records {
		log.Println(records)
		db.AddRecord(&record)
	}
	c.JSON(http.StatusOK, records)
}

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
			UseHeroName: getHeroNameById(db.HeroList, jy.Hero),
			CreateTime:  currentTime,
		}
		db.AddRecord(&record)
	}
	db.DisableHero(ids)
	c.JSON(http.StatusOK, jies)
}

// 根据英雄id获取英雄名称
func getHeroNameById(list []db.Hero, heroId int64) string {
	for _, one := range list {
		if one.Id == heroId {
			return one.HeroName
		}
	}
	return ""
}
