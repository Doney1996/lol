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
}
