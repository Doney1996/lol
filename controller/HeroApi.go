package controller

import (
	"github.com/gin-gonic/gin"
	"lol/entity"
	"lol/repository/hero"
	"lol/repository/match"
	"lol/repository/record"
	"lol/repository/season"
	"net/http"
)

// GetHero 所有英雄
func GetHero(c *gin.Context) {
	heroList := hero.GetAllHero()
	c.JSON(http.StatusOK, heroList)
}

// GetHeroBySeason 赛季英雄状况
func GetHeroBySeason(c *gin.Context) {
	//获取当前赛季
	currentSeason := season.GetCurrentSeason("class", "0")
	if currentSeason == (entity.Season{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "赛季还未开始，请开始赛季",
			Data:    nil,
		})
		return
	}
	matchList := match.GetCurrentSeasonMatch(currentSeason.Id)
	var matchIds []int64
	for i, e := range matchList {
		matchIds[i] = e.Id
	}
	records := record.GetRecordByMatchIds(matchIds)

	//禁用对所有英雄
	var disableHeroIds []int64
	for i, e := range records {
		disableHeroIds[i] = e.Id
	}
	c.JSON(http.StatusOK, disableHeroIds)
}

// 英雄层级
func getHeroTier() {

}
