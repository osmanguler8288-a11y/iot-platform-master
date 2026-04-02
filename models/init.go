package models

import (
	"iot-platform-master/define"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := define.MySqlDSN + "/iot-platform?charset=utf8mb4&parseTime=True&loc=Local" //后面是数据库的名称

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[db:err]", err)
	}
	db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})

	DB = db
}
