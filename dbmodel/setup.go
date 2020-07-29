package dbmodel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDataBase() {
	fmt.Println("GOING TO CONNECT DB .....")
	database, err := gorm.Open("postgres", "host=<> port=<> user=<> dbname=<> password=<> sslmode=disable")
	//defer database.Close()

	if err != nil {
		fmt.Println("ERROR WHILE CONNECTING DB: ",err.Error())
		panic("Failed to connect to database!")
	}
	fmt.Println("CONNECTING DB SUCCESS")
	database.AutoMigrate(&TrainResource{})

	DB = database
}
