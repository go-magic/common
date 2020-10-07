package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mysql struct {
	db *gorm.DB
}

var my *mysql

func InitMysql(path string) error {
	my = new(mysql)
	var err error
	my.db, err = gorm.Open("mysql", path)
	if err != nil {
		return err
	}
	my.db.SingularTable(true)
	return nil
}
