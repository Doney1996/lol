package controller

import (
	"github.com/gin-gonic/gin"
	"lol/cache"
	"lol/entity"
	"lol/repository/repo_match"
	"lol/repository/repo_record"
	"lol/repository/repo_season"
	"net/http"
)

// GetAllHero 所有英雄
func GetAllHero(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Result{
		Code:    200,
		Message: "",
		Data:    cache.HeroList,
	})
}

// GetHeroBySeason 当前赛季禁用的英雄
func GetHeroBySeason(c *gin.Context) {
	//获取当前赛季
	gameType := c.MustGet("gameType").(string)
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason == (entity.Season{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    410,
			Message: "赛季还未开始，请开始赛季",
			Data:    nil,
		})
		return
	}
	//当前赛季所有结束的对局
	matchList := repo_match.GetMatchBySeasonAndStatus(currentSeason.Id, "1")
	var matchIds []int64
	for _, e := range matchList {
		matchIds = append(matchIds, e.Id)
	}
	records := repo_record.GetRecordsByMatchIds(matchIds)

	//禁用的所有英雄的id
	var disableHeroIds []int64
	for _, e := range records {
		disableHeroIds = append(disableHeroIds, e.HeroId)
	}

	c.JSON(http.StatusOK, disableHeroIds)
}

// 英雄层级
func getHeroTier() {

}
