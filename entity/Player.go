package entity

type Player struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	GameName string `json:"game_name" db:"game_name"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
}
