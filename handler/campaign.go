package handler

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
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

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	campaignID := campaign.GetCampaignDetailInput{}
	updateCampaignData := campaign.CreateCampaignInput{}

	err := c.ShouldBindUri(&campaignID)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	err = c.ShouldBindJSON(&updateCampaignData)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Add user
	currentUser := c.MustGet("currentUser").(user.User)
	updateCampaignData.User = currentUser
	fmt.Println("campaign id", campaignID)
	updatedCampaign, err := h.service.UpdateCampaign(campaignID.ID, updateCampaignData)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaign := campaign.FormatCampaign(updatedCampaign)
	response := helper.APIResponse("Success update existing campaign", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	input := campaign.CreateCampaignImageInput{}
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "failed", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save campaign image", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.UploadCampaignImage(input, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save campaign image", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success uploading campaign image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
