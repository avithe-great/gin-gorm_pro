package main

import (
	"book-store/Config"
	"book-store/Models"
	"book-store/Routers"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/collbook?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
