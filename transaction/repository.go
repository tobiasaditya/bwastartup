package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindByCampaignID(campaignID int) ([]Transaction, error) {
	campaigns := []Transaction{}

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}
