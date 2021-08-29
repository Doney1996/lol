package entity

import "time"

type Record struct {
	Id         int64     `json:"id" db:"id"`
	UserId     int64     `json:"user_id" db:"user_id"`
	HeroId     int64     `json:"hero_id" db:"hero_id"`
	MatchId    int64     `json:"match_id" db:"match_id"`
	IsWin      int64     `json:"is_win" db:"is_win"`
	Score      float64   `json:"score" db:"score"`
	UnitPrice  float64   `json:"unit_price" db:"unit_price"`
	LastScore  float64   `json:"last_score" db:"last_score"`
	Amount     float64   `json:"amount" db:"amount"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
