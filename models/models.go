package models

import (
	"blog-backend/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetUp() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", config.Conf.Database.User, config.Conf.Database.Password, config.Conf.Database.Host, config.Conf.Database.Port, config.Conf.Database.DBName)
	db, err = gorm.Open(config.Conf.Database.Type, dsn)
	if err != nil {
		log.Fatalf("Connect database fail, %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
