package controller

import (
	"github.com/gin-gonic/gin"
	"lol/cache"
	"lol/entity"
	"lol/repository/repo_match"
	"lol/repository/repo_record"
	"lol/repository/repo_season"
	"math"
	"net/http"
	"sort"
	"time"
)

// OpenNewMatch 开启新的一场对局
func OpenNewMatch(c *gin.Context) {
	// 开启关联匹配表 状态为 未结算
	gameType := c.GetString("gameType")
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason == (entity.Season{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "赛季还未开始，请开始赛季",
			Data:    nil,
		})
		return
	}

	seasonMatch := repo_match.GetMatchBySeasonAndStatus(currentSeason.Id, "0")

	//当前赛有未结束的匹配
	if len(seasonMatch) > 0 {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "有未结束的对局，不能开启新的",
			Data:    nil,
		})
		return
	}

	m := entity.Match{
		SeasonId:   currentSeason.Id,
		LifeStatus: 0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	match := repo_match.Insert(m)
	c.JSON(http.StatusOK, entity.Result{
		Code:    100,
		Message: "开启新对局",
		Data:    match,
	})
}

// GetLastMatch 获取最后一次的匹配信息
func GetLastMatch(c *gin.Context) {
	// 查询匹配表对应的最后一条记录
	gameType := c.GetString("gameType")
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason == (entity.Season{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    101,
			Message: "赛季还未开始，请开始赛季",
			Data:    nil,
		})
		return
	}
	match := repo_match.GetLastBySeasonAndStatus(currentSeason.Id, "1")
	if match == (entity.Match{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    100,
			Message: "无战绩",
			Data:    nil,
		})
		return
	}
	records := repo_record.GetByMatchId(match.Id)
	type rec struct {
		HeroName  string  `json:"hero_name,omitempty"`
		PlayName  string  `json:"play_name,omitempty"`
		Score     float64 `json:"score,omitempty"`
		UintPrice float64 `json:"unit_price,omitempty"`
		LastScore float64 `json:"last_score,omitempty"`
		Amount    float64 `json:"amount,omitempty"`
		IsWin     int64   `json:"is_win,omitempty"`
	}
	i := len(records)
	var resultList = make([]rec, i)

	sort.Slice(records, func(i, j int) bool {
		return records[i].Score < records[j].Score
	})

	for index, record := range records {
		resultList[index] = rec{
			HeroName:  cache.GetHeroNameById(record.HeroId),
			PlayName:  cache.GetPlayNameById(record.UserId),
			Score:     record.Score,
			UintPrice: record.UnitPrice,
			LastScore: record.LastScore,
			Amount:    record.Amount,
			IsWin:     record.IsWin,
		}
	}

	c.JSON(http.StatusOK, entity.Result{
		Code:    100,
		Message: "最新战绩",
		Data:    resultList,
	})
}

// CloseNewMatch 结算新的一场对局
func CloseNewMatch(c *gin.Context) {
	// 完成结算
	gameType := c.GetString("gameType")
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	match := repo_match.GetLastBySeasonId(currentSeason.Id)
	records := repo_record.GetByMatchId(match.Id)

	now := time.Now()
	//计算出错 关闭场次
	for index, record := range records {
		subtotal := 0.0
		for _, tmp := range records {
			subScore := record.Score - tmp.Score
			if subScore < 0 {
				subScore--
			}
			subtotal += math.Ceil(subScore)
		}
		records[index].LastScore = subtotal

		if subtotal > 0 {
			records[index].IsWin = 1
		} else if subtotal < 0 {
			records[index].IsWin = -1
		} else {
			records[index].IsWin = 0
		}
		records[index].Amount = subtotal * record.UnitPrice
		records[index].UpdateTime = now
	}

	for _, record := range records {
		repo_record.Update(record)
	}

	// 前端关闭结算窗口
	match.LifeStatus = 1
	match.UpdateTime = now
	repo_match.Update(match)
}
