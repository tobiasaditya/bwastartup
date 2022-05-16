package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type DetailCampaignUser struct {
	Name           string `json:"name"`
	AvatarFileName string `json:"avatar_url"`
}

type DetailCampaignImage struct {
	FileName  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}
type DetailCampaignFormatter struct {
	ID               int                   `json:"id"`
	UserID           int                   `json:"user_id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	ImageURL         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	Description      string                `json:"description"`
	Perks            []string              `json:"perks"`
	User             DetailCampaignUser    `json:"user"`
	CampaignImages   []DetailCampaignImage `json:"images"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	format := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		format.ImageURL = campaign.CampaignImages[0].FileName
	}
	return format
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	formatters := []CampaignFormatter{} //inisialisasi gini biar balikannya ttp list kosong, bukan nil

	for _, campaign := range campaigns {
		formatters = append(formatters, FormatCampaign(campaign))
	}

	return formatters
}

func FormatDetailCampaignImage(campaignImage CampaignImage) DetailCampaignImage {
	isPrimary := false
	if campaignImage.IsPrimary == 1 {
		isPrimary = true
	}
	imageFormatter := DetailCampaignImage{
		FileName:  campaignImage.FileName,
		IsPrimary: isPrimary,
	}
	return imageFormatter
}

func FormatDetailCampaignImages(campaignImages []CampaignImage) []DetailCampaignImage {
	imageFormatters := []DetailCampaignImage{}

	for _, campaignImage := range campaignImages {
		imageFormatters = append(imageFormatters, FormatDetailCampaignImage(campaignImage))
	}
	return imageFormatters
}

func FormatDetailCampaign(campaign Campaign) DetailCampaignFormatter {
	formatter := DetailCampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Description:      campaign.Description,
	}

	perksFormatter := []string{}
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perksFormatter = append(perksFormatter, strings.TrimSpace(perk))
	}

	formatter.Perks = perksFormatter

	userFormatter := DetailCampaignUser{
		Name:           campaign.User.Name,
		AvatarFileName: campaign.User.AvatarFileName,
	}
	formatter.User = userFormatter

	imageFormatters := FormatDetailCampaignImages(campaign.CampaignImages)
	formatter.CampaignImages = imageFormatters

	if len(imageFormatters) > 0 {
		for _, image := range imageFormatters {
			if image.IsPrimary {
				formatter.ImageURL = image.FileName
			}
		}
	}

	return formatter
}
