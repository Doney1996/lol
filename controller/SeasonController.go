package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/entity"
	"lol/repository/season"
	"net/http"
	"strconv"
	"time"
)

func OpenNewSeason(c *gin.Context) {
	currentSeason := season.GetCurrentSeason("class", "0")
	if currentSeason.Id > 0 {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "当前赛季没有结束",
			Data:    nil,
		})
	} else {
		finishedSeason := season.GetCurrentSeason("class", "1")

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
		season.InsertSeason(&newSeason)
		c.JSON(http.StatusOK, entity.Result{
			Code:    100,
			Message: "新赛季创建成功",
			Data:    nil,
		})
	}
}

func CloseSeason(c *gin.Context) {
	currentSeason := season.GetCurrentSeason("class", "0")
	if currentSeason.Id > 0 {
		season.CloseSeason(currentSeason.Id)
		c.JSON(http.StatusOK, entity.Result{
			Code:    100,
			Message: "当前赛季已经成功结束",
			Data:    currentSeason.SeasonName,
		})
	}

}
