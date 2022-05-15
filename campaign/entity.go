package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
