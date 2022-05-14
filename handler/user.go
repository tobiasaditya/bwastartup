package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

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
