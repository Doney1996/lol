package common

import (
	"bytes"
	"encoding/gob"
	"github.com/jinzhu/gorm"
)

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

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
