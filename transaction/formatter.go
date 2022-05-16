package transaction

type TransactionFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
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
