package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/entity"
	"lol/expection"
	"lol/repository/repo_match"
	"lol/repository/repo_record"
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
	match := repo_match.GetById(tmp.MatchId)
	if match.LifeStatus != 0 {
		panic(expection.BizErr{
			Code: 410,
			Msg:  "当前对局已结束",
		})
	}

	//如何避免重复提交呢
	insert := repo_record.Insert(recordObj)
	c.JSON(http.StatusOK, entity.Result{
		Code:    200,
		Message: "保存成功",
		Data:    insert,
	})

}
