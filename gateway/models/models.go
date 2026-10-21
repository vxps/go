package models

type CreateTransactionRequest struct {
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type TransactionResponse struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type CreateBudgetRequest struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Period   string `json:"period"`
}

type BudgetResponse struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
}
