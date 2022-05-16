package transaction

type TransactionFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type CreateTransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentUrl string `json:"payment_url"`
}

type CampaignTransactionFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type UserTransactionFormatter struct {
	ID        int                          `json:"id"`
	Amount    int                          `json:"name"`
	Status    string                       `json:"status"`
	CreatedAt string                       `json:"created_at"`
	Campaign  CampaignTransactionFormatter `json:"campaign"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt.String(),
	}
	return formatter

}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {
	formatter := []TransactionFormatter{}
	for _, transaction := range transactions {
		formatter = append(formatter, FormatTransaction(transaction))
	}
	return formatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	//Cek campaign images
	campaignImages := transaction.Campaign.CampaignImages
	imageUrl := ""
	if len(campaignImages) > 0 {
		imageUrl = campaignImages[0].FileName
	}

	campaignFormatter := CampaignTransactionFormatter{
		Name:     transaction.Campaign.Name,
		ImageUrl: imageUrl,
	}
	formatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt.String(),
		Campaign:  campaignFormatter,
	}
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	formatter := []UserTransactionFormatter{}
	for _, transaction := range transactions {
		formatter = append(formatter, FormatUserTransaction(transaction))
	}
	return formatter
}

func FormatCreateTransaction(transaction Transaction) CreateTransactionFormatter {
	return CreateTransactionFormatter{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentUrl: transaction.PaymentUrl,
	}
}
