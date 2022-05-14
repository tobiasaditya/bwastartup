package main

import (
	"bwastartup/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "tobiasaditya:!D3papepe@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database success")

	userRepository := user.NewRepository(db)
	userSevice := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Testing Service"
	userInput.Email = "testing.service@gmail.com"
	userInput.Occupation = "Existential Crisis"
	userInput.Password = "Owner123"

	userSevice.RegisterUser(userInput)
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal()
	}
}
