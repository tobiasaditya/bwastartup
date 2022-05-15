package campaign

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
