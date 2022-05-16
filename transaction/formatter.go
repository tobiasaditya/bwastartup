package transaction

type TransactionFormatter struct {
	ID        int
	Name      string
	Amount    int
	CreatedAt string
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
