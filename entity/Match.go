package entity

import (
	"time"
)

type Match struct {
	Id           int64
	SeasonId     int64
	LifeStatus   int64
	PlayerNumber int64
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (Match) TableName() string {
	return "game_match"
}
