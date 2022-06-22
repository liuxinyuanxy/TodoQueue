package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	connectDatabase()
	err := db.AutoMigrate(&User{}, &Todo{}, &TodoDone{}, &Template{})
	if err != nil {
		logrus.Fatal(err)
	}
}

func connectDatabase() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("psql")
	dbArgs := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", loginInfo["user"], loginInfo["password"], loginInfo["host"], loginInfo["port"], loginInfo["dbname"])
	var err error
	db, err = gorm.Open(postgres.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
