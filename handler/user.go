package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// Terima Input User
	// Map Input User jadi User
	// Simpen User ke service register

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to generate token", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// Terima input User (email dan password)
	// Mapping input login ke struct
	// Parsing ke service untuk cari user dengan email
	// Jika ketemu, baru verify password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	foundUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to login", http.StatusUnauthorized, "failed", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	token, err := h.authService.GenerateToken(foundUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to generate token", http.StatusUnauthorized, "failed", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	formatter := user.FormatUser(foundUser, token)

	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	// Terima input email
	// Mapping input login ke struct
	// Parsing ke service untuk cari user dengan email
	// Jika ketemu, balikin error duplikat email ditemukan
	var input user.EmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvaiable, err := h.userService.CheckEmailUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to check email", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// formatter := user.FormatUser(foundUser, "tes")
	metaMessage := ""
	if isEmailAvaiable {
		metaMessage = "Email is avaiable"
	} else {
		metaMessage = "Email is unavaiable (duplicate)"
	}
	formatter := gin.H{"is_avaiable": isEmailAvaiable}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// Input dari user, bukan json tapi form data
	// Simpan gambar ke folder (images)
	// Service panggil repo user
	// JWT (sementara hardcode), seakan" user yang login punya id = 15
	// Repo update data user, simpen filename gambar tadi
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 15 //Hardcode ID
	// Ngubah nama file yang disimpen, ditambahain userId yang unik
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success uploading avatar", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
