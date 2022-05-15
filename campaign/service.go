package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
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
