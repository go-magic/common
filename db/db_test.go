package db

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type User struct {
	UserId int    `gorm:"AUTO_INCREMENT;primary_key"`
	Phone  string `gorm:"not null;unique_index;index:idx_name_phone"`
}

func init() {
	InitMysql("root:aU0FEQ^#ad1@/sell?charset=utf8&parseTime=True&loc=Local")
}

func TestCreateData(t *testing.T) {

}
