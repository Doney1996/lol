package entity

type Hero struct {
	Id            int64  `db:"id"`
	HeroName      string `db:"hero_name"`
	HeroOtherName string `db:"hero_other_name"`
	Sort          int64  `db:"sort"`
	Disable       bool   `db:"disable"`
	Tier          int64  `db:"tier"`
	ImgPosition   int64  `db:"img_position"`
	Position      string `db:"position"`
}
