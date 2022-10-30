package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"lr25v4_back/config"
)

var database *gorm.DB

func Init() *gorm.DB {
	database, err = gorm.Open(
		mysql.Open("lr21v4"), &gorm.Config{})
	if err != nil {
		fmt.Println("something wrong with db")
	}


}

func GetDB() *gorm.DB {

}
