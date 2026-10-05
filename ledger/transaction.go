package main

import (
	"errors"
	"slices"
)

var ErrBudgetExceeded = errors.New("budget exceeded")

type Transaction struct {
	ID          int
	Amount      int
	Category    string
	Description string
	Date        string
}

var transactions = make([]Transaction, 0)

func AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("transaction amount cannot be zero")
	}

	if budget, ok := budgets[tx.Category]; ok {
		currentSpent := sumByCategory(tx.Category)

		if currentSpent+tx.Amount > budget.Limit {
			return ErrBudgetExceeded
		}
	}

	if tx.ID == 0 {
		tx.ID = len(transactions) + 1
	}

	transactions = append(transactions, tx)

	return nil
}

func ListTransactions() []Transaction {
	return slices.Clone(transactions)
}

func sumByCategory(category string) int {
	var total int

	for _, tx := range transactions {
		if tx.Category == category {
			total += tx.Amount
		}
	}

	return total
}