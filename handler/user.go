package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
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
	formatter := user.FormatUser(newUser, "tes")

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
	formatter := user.FormatUser(foundUser, "tes")

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
