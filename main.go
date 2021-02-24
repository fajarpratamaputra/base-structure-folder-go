//main.go
package main

import (
	"fmt"
	"mvc-api/Config"
	"mvc-api/Models"
	"mvc-api/Routes"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	// the jwt middleware
	r := Routes.SetupRouter()

	//running
	port := os.Getenv("PORT")

	r.Run(":" + port)
}
