package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
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
	service campaign.Service
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
