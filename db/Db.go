package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"log"
	"lol/common"
	"lol/entity"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:admin1984@tcp(114.96.105.111:3306)/lol_dev?parseTime=true")
	common.DealErr(err)
	db.SingularTable(true)
	db.AutoMigrate(&entity.Record{})
	db.AutoMigrate(&entity.Match{})
	db.AutoMigrate(&entity.Season{})
	db.AutoMigrate(&entity.Player{})
	db.AutoMigrate(&entity.Hero{})
	DB = db
}

func original() *sqlx.DB {
	log.Println("------------int database-------------")
	database, err := sqlx.Open("mysql", "root:admin1984@tcp(114.96.105.111:3306)/lol_dev?parseTime=true")
	if err != nil {
		panic("open mysql failed." + err.Error())
		return nil
	}
	return database
}
