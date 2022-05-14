package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	userHandler := handler.NewUserHandler(userSevice)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.Login)

	router.Run()
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal()
	}
}
