package common

import "github.com/jinzhu/gorm"

func DealErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DealDbErrs(db *gorm.DB) {
	errors := db.GetErrors()
	if len(errors) > 0 {
		db.Rollback()
		panic(errors[0])
	}
}
