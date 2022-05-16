package transaction

import (
	"bwastartup/campaign"
	"errors"
	"time"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository: repository, campaignRepository: campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	foundCampaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if foundCampaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("User can't access this campaign transactions")
	}

	transactions, err := s.repository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	newTransaction := Transaction{
		CampaignID: input.CampaignID,
		UserID:     input.User.ID,
		Amount:     input.Amount,
		Status:     "PENDING",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	insertedTransaction, err := s.repository.Create(newTransaction)
	if err != nil {
		return insertedTransaction, err
	}
	return insertedTransaction, err
}
