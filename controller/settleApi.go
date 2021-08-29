package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/entity"
	"lol/repository/match"
	"lol/repository/record"
	"lol/repository/season"
	"net/http"
	"time"
)

//  获取最后一次的匹配信息
func getLastMatch(c *gin.Context) {
	// 查询匹配表对应的最后一条记录
	// 如果匹配表是完成状态，需要上传
	// 不是的就查询显示 如果已经包含自己 等待他人上传
}

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
	insert := record.Insert(recordObj)
	c.JSON(http.StatusOK, entity.Result{
		Code:    100,
		Message: "保存成功",
		Data:    insert,
	})

}

// OpenNewMatch 开启新的一场对局
func OpenNewMatch(c *gin.Context) {
	// 开启关联匹配表 状态为 未结算
	currentSeason := season.GetCurrentSeason("class", "0")
	seasonMatch := match.GetCurrentSeasonMatch(currentSeason.Id)
	if !checkLifeStatus(c, seasonMatch) {
		return
	}
	match.Insert(entity.Match{
		SeasonId:   currentSeason.Id,
		LifeStatus: 0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
}

func checkLifeStatus(c *gin.Context, seasonMatch []entity.Match) bool {
	for _, e := range seasonMatch {
		if e.LifeStatus == 0 {
			c.JSON(http.StatusOK, entity.Result{
				Code:    101,
				Message: "有未结束的对局，不能开启",
				Data:    nil,
			})
			return false
		}
	}
	return true
}

// 结算新的一场对局
func settleNewMatch(c *gin.Context) {
	// 完成结算
	//计算出错 关闭场次
	// 前端关闭结算窗口
	// 展示最新战绩
}
