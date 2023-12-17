package config

//importing all req lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {

	//write username, password, hostname and all to connect database
	d, err := gorm.Open("mysql", "username:password@tcp(host:port)/laundry_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
