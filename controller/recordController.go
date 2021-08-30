package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/entity"
	"lol/repository/repo_record"
	"lol/repository/repo_season"
	"net/http"
	"time"
)

// AddRecord 新增一条record
func AddRecord(c *gin.Context) {
	user := c.MustGet("claims").(common.CustomClaims)

	tmp := entity.Record{}
	err := c.BindJSON(&tmp)
	common.DealErr(err)

	recordObj := entity.Record{
		UserId:     user.ID,
		HeroId:     tmp.HeroId,
		MatchId:    tmp.MatchId,
		Score:      tmp.Score,
		UnitPrice:  tmp.UnitPrice,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}

	//检查对局的状态
	sea := repo_season.GetById(tmp.MatchId)
	if sea.LifeStatus != 0 {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "当前对局已结束",
			Data:    nil,
		})
		return
	}

	insert := repo_record.Insert(recordObj)
	c.JSON(http.StatusOK, entity.Result{
		Code:    100,
		Message: "保存成功",
		Data:    insert,
	})

}
