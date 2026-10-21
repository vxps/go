package mapper

import (
	"gateway/models"
	"ledger"
)

func ToTransactionDomain(r models.CreateTransactionRequest) ledger.Transaction {
	return ledger.Transaction{
		Amount:      r.Amount,
		Category:    r.Category,
		Description: r.Description,
		Date:        r.Date,
	}
}

func ToBudgetDomain(r models.CreateBudgetRequest) ledger.Budget {
	return ledger.Budget{
		Category: r.Category,
		Limit:    r.Limit,
		Period:   r.Period,
	}
}

func ToTransactionResponse(tx ledger.Transaction) models.TransactionResponse {
	return models.TransactionResponse{
		ID:          tx.ID,
		Amount:      tx.Amount,
		Category:    tx.Category,
		Description: tx.Description,
		Date:        tx.Date,
	}
}

func ToBudgetResponse(b ledger.Budget) models.BudgetResponse {
	return models.BudgetResponse{
		Category: b.Category,
		Limit:    b.Limit,
	}
}
