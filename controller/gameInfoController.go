package controller

import (
	"github.com/gin-gonic/gin"
	"lol/common"
	"lol/db"
	"net/http"
)

type result struct {
	Name     string  `db:"use_hero_name" json:"name"`
	Score    float64 `db:"score" json:"score"`
	SubScore float64 `db:"sub_score" json:"sub_score"`
	SubTotal int64   `db:"Subtotal" json:"sub_total"`
}

func GetRecentResult(c *gin.Context) {
	sql := `select use_hero_name,score,sub_score,Subtotal
from recording
where create_time = (select create_time from recording order by create_time desc limit 1)`

	var list []result
	db := db.DB.Select(&list, sql)
	common.DealDbErrs(db)
	c.JSON(http.StatusOK, list)
}
