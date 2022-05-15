package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(ID int) (Campaign, error)
	CreateCampaign(inputCampaign CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID == 0 {
		foundCampaigns, err := s.repository.FindAll()
		if err != nil {
			return foundCampaigns, err
		}
		return foundCampaigns, nil
	} else {
		foundCampaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return foundCampaigns, err
		}
		return foundCampaigns, nil
	}
}

func (s *service) GetCampaign(ID int) (Campaign, error) {
	foundCampaign, err := s.repository.FindByID(ID)
	if err != nil {
		return foundCampaign, err
	}
	return foundCampaign, nil

}

func (s *service) CreateCampaign(inputCampaign CreateCampaignInput) (Campaign, error) {
	newCampaign := Campaign{
		UserId:           inputCampaign.User.ID,
		Name:             inputCampaign.Name,
		ShortDescription: inputCampaign.ShortDescription,
		Description:      inputCampaign.Description,
		GoalAmount:       inputCampaign.GoalAmount,
		Perks:            inputCampaign.Perks,
	}

	stringSlug := fmt.Sprintf("%s %d", inputCampaign.Name, inputCampaign.User.ID)
	newSlug := slug.Make(stringSlug)

	newCampaign.Slug = newSlug

	insertedCampaign, err := s.repository.Create(newCampaign)
	if err != nil {
		return insertedCampaign, err
	}

	return insertedCampaign, err
}
