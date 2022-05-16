package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	authService := auth.NewJwtService()

	userRepository := user.NewRepository(db)
	userSevice := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userSevice, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.Login)
	api.POST("/user/email", userHandler.CheckEmail)
	api.POST("/user/avatar", authMiddleware(authService, userSevice), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userSevice), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userSevice), campaignHandler.UpdateCampaign)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header, Authorization Bearer tokentokentoken
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		headerToken := strings.Split(authHeader, " ")

		if len(headerToken) == 2 {
			tokenString = headerToken[1]
		}

		// Validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Kalau valid, decode token, buat dapetin claim/payload dari jwt token
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		//Get user by Id
		foundUser, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println(foundUser.ID)
		c.Set("currentUser", foundUser)
	}

}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal()
	}
}
