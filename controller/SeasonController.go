package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/entity"
	"lol/repository/repo_season"
	"net/http"
	"strconv"
	"time"
)

func OpenNewSeason(c *gin.Context) {
	gameType := c.MustGet("gameType").(string)
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason.Id > 0 {
		c.JSON(http.StatusOK, entity.Result{
			Code:    500,
			Message: "当前赛季没有结束",
			Data:    nil,
		})
	} else {
		finishedSeason := repo_season.GetCurrentSeason("class", "1")

		if finishedSeason.Id < 1 {
			finishedSeason.SeasonName = "S0"
		}

		index, err := strconv.Atoi(string(finishedSeason.SeasonName[1]))
		common.DealErr(err)
		strconv.Itoa(index + 1)
		newSeason := entity.Season{
			SeasonName: "S" + strconv.Itoa(index+1),
			LifeStatus: 0,
			CreateTime: time.Now(),
			SeasonType: "class",
		}
		repo_season.InsertSeason(&newSeason)
		c.JSON(http.StatusOK, entity.Result{
			Code:    200,
			Message: "新赛季创建成功",
			Data:    nil,
		})
	}
}

func CloseSeason(c *gin.Context) {
	gameType := c.MustGet("gameType").(string)
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason.Id > 0 {
		repo_season.CloseSeason(currentSeason.Id)
		c.JSON(http.StatusOK, entity.Result{
			Code:    200,
			Message: "当前赛季已经成功结束",
			Data:    currentSeason.SeasonName,
		})
	}

}
