package storage

import (
	"awesomeProject/server_DB_gl_er_handler/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "postgres://yeieycvd:xxxxxxxxxxxxxxx@trumpet.db.elephantsql.com/yeieycvd"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err.Error())
		panic("AutoMigrate problem")
		return
	}
}
