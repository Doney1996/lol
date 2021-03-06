package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lol/cache"
	"lol/common"
	"lol/entity"
	"lol/expection"
	"lol/repository/repo_match"
	"lol/repository/repo_record"
	"lol/repository/repo_season"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// OpenNewMatch 开启新的一场对局
func OpenNewMatch(c *gin.Context) {
	// 开启关联匹配表 状态为 未结算
	gameType := c.GetString("gameType")
	type Request struct {
		PlayerNumber int64 `json:"player_number,omitempty"`
	}
	var request Request

	err := c.BindJSON(&request)
	common.DealErr(err)
	if request.PlayerNumber < 2 || request.PlayerNumber > 5 {
		panic(expection.BizErr{
			Code: 410,
			Msg:  "玩家数量不能小于2或大于5,当前玩家数量:" + strconv.FormatInt(request.PlayerNumber, 10),
		})
	}
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	if currentSeason == (entity.Season{}) {
		panic(expection.BizErr{
			Code: 410,
			Msg:  "有未结束的对局，不能开启新的",
		})
	}

	seasonMatch := repo_match.GetMatchBySeasonAndStatus(currentSeason.Id, "0")

	//当前赛有未结束的匹配
	if len(seasonMatch) > 0 {
		panic(expection.BizErr{
			Code: 410,
			Msg:  "有未结束的对局，不能开启新的",
		})
	}

	m := entity.Match{
		SeasonId:     currentSeason.Id,
		LifeStatus:   0,
		PlayerNumber: request.PlayerNumber,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	match := repo_match.Insert(m)
	c.JSON(http.StatusOK, entity.Result{
		Code:    200,
		Message: "开启新对局成功",
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
			Code:    500,
			Message: "赛季还未开始，请开始赛季",
			Data:    nil,
		})
		return
	}
	match := repo_match.GetLastBySeasonAndStatus(currentSeason.Id, "1")
	if match == (entity.Match{}) {
		c.JSON(http.StatusOK, entity.Result{
			Code:    200,
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
		Code:    200,
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
	if int64(len(records)) != match.PlayerNumber {
		msg := fmt.Sprintf("玩家数量错误，当前玩家数量：%d，已提交：%d", match.PlayerNumber, len(records))
		panic(expection.BizErr{
			Code: 410,
			Msg:  msg,
		})
	}

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
	c.JSON(http.StatusOK, entity.Result{
		Code:    200,
		Message: "对局结束",
		Data:    records,
	})
}

func NeedSubmitNew(c *gin.Context) {
	gameType := c.GetString("gameType")
	user := c.MustGet("claims").(common.CustomClaims)
	currentSeason := repo_season.GetCurrentSeason(gameType, "0")
	match := repo_match.GetLastBySeasonId(currentSeason.Id)
	mapData := map[string]interface{}{"need_submit": true}
	result := entity.Result{
		Code:    200,
		Message: "",
		Data:    mapData,
	}
	if match.LifeStatus != 0 {
		mapData["need_submit"] = false
		c.JSON(http.StatusOK, result)
		return
	}
	record := repo_record.GetRecordBYMatchIdAndUserId(match.Id, user.ID)
	if record != (entity.Record{}) {
		mapData["need_submit"] = true
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
