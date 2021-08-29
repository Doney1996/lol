package entity

type Hero struct {
	Id            int64  `json:"id" db:"id"`
	HeroName      string `json:"heroName" db:"hero_name"`
	HeroOtherName string `json:"heroOtherName" db:"hero_other_name"`
	Sort          int64  `json:"sort" db:"sort"`
	ImgPosition   int64  `json:"imgPosition" db:"img_position"`
	Position      string `json:"position" db:"position"`
}
