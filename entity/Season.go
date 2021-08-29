package entity

import "time"

type Season struct {
	Id         int64     `json:"id" db:"id"`
	SeasonName string    `json:"season_name" db:"season_name"`
	LifeStatus int64     `json:"life_status" db:"life_status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	SeasonType string    `json:"season_type" db:"season_type"`
}
