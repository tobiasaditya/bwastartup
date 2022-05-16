package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(ID int) (Campaign, error)
	CreateCampaign(inputCampaign CreateCampaignInput) (Campaign, error)
	UpdateCampaign(ID int, updateCampaign CreateCampaignInput) (Campaign, error)
	UploadCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) UpdateCampaign(ID int, updateCampaign CreateCampaignInput) (Campaign, error) {
	foundCampaign, err := s.repository.FindByID(ID)

	if err != nil {
		return foundCampaign, err
	}

	//Validasi user, usernya ga boleh beda dengan yang buat campaign
	if foundCampaign.UserId != updateCampaign.User.ID {
		return foundCampaign, errors.New("User can't update this campaign")
	}

	foundCampaign.Name = updateCampaign.Name
	foundCampaign.Description = updateCampaign.Description
	foundCampaign.ShortDescription = updateCampaign.ShortDescription
	foundCampaign.GoalAmount = updateCampaign.GoalAmount
	foundCampaign.Perks = updateCampaign.Perks
	foundCampaign.UpdatedAt = time.Now()

	updatedCampaign, err := s.repository.Update(foundCampaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}

func (s *service) UploadCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	foundCampaign, err := s.repository.FindByID(input.CampaignID)

	if err != nil {
		return CampaignImage{}, err
	}

	//Validasi user, usernya ga boleh beda dengan yang buat campaign
	if foundCampaign.UserId != input.User.ID {
		return CampaignImage{}, errors.New("User can't update this campaign")
	}

	//Handling jika input isPrimary true, cari ke db, images dengan campaignID tersebut, dibikin false
	if input.IsPrimary {
		_, err := s.repository.UpdateImageAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	// Convert isPrimary jadi int
	intIsPrimary := 0
	if input.IsPrimary {
		intIsPrimary = 1
	}
	newCampaignImage := CampaignImage{
		CampaignId: input.CampaignID,
		FileName:   fileLocation,
		IsPrimary:  intIsPrimary,
	}

	insertedImage, err := s.repository.CreateImage(newCampaignImage)
	if err != nil {
		return insertedImage, err
	}
	return insertedImage, nil
}
