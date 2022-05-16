package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
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
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository: repository, campaignRepository: campaignRepository, paymentService: paymentService}
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

	//Midtrans
	paymentTransaction := payment.Transaction{
		ID:     insertedTransaction.ID,
		Amount: insertedTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return insertedTransaction, err
	}
	insertedTransaction.PaymentUrl = paymentURL

	updatedTransaction, err := s.repository.Update(insertedTransaction)

	return updatedTransaction, err
}
