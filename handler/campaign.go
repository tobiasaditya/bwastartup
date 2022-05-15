package handler

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Fitur
// 2.List campaign (filter by user, atau kalau no filter, tampilin semua)
//	tangkap input filter
//  ke service, menentukan apakah dapetin semua campaign atau filter by user_id
// repository akses ke db

type campaignHandler struct {
	service     campaign.Service
	authService auth.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	foundCampaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get campaigns", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns := campaign.FormatCampaigns(foundCampaigns)

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	input := campaign.GetCampaignDetailInput{}
	c.ShouldBindUri(&input)

	foundCampaign, err := h.service.GetCampaign(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign := campaign.FormatDetailCampaign(foundCampaign)

	response := helper.APIResponse("Success get detail campaign", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	createdCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to create new campaign", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign := campaign.FormatCampaign(createdCampaign)

	response := helper.APIResponse("Success create new campaign", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)

}
