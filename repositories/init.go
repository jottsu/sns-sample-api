package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf(
		"%v:%v@tcp(db)/sns_sample?parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
	))
	if err != nil {
		panic(err)
	}
}

func Begin() (*gorm.DB, error) {
	tx := DB.Begin()
	return tx, tx.Error
}

func Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}
