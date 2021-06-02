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
	CreateTime  string
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

var HeroList []Hero

var Db *sqlx.DB

func init() {
	log.Println("-----------------------------------")
	database, err := sqlx.Open("mysql", "root:admin1984@tcp(114.96.105.111:3306)/lol")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	log.Println(database, "database init success")

	//查询出所有的英雄记录放在内存
	getAllHero()
}

func getAllHero() {
	HeroList = GetAllHero()
}
