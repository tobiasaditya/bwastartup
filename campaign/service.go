package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(ID int) (Campaign, error)
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
