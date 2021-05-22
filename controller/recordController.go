package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"lol/common"
	"lol/db"
	"net/http"
	"strconv"
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
	for _, jy := range jies {
		ids = append(ids, strconv.FormatInt(jy.Hero, 10))
		var w int
		if jy.SubTotal>0 {
			w = 1
		}else if jy.SubTotal == 0 {
			w = 0
		}else {
			w = -1
		}
		record := db.Record{
			UseHeroId: jy.Hero,
			Score:     jy.Score,
			Win:       w,
			Subtotal: jy.SubTotal,
		}
		db.AddRecord(&record)
	}
	db.DisableHero(ids)
	c.JSON(http.StatusOK, jies)
}
