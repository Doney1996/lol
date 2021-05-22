package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Record struct {
	Id          int64
	PlayId      int64
	PlayerId    int64
	PlayerName  string
	UseHeroId   int64
	UseHeroName string
	Win         int
	Score       float64
	UnitPrice   int
	Subtotal    float64
}

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

var Db *sqlx.DB

func init() {
	log.Println("-----------------------------------")
	database, err := sqlx.Open("mysql", "root:admin1984@tcp(119.45.13.75:3306)/lol")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	log.Println(database, "database init success")
}
