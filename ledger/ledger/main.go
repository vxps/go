package main

import (
	"bufio"
	"errors"
	"fmt"
	"ledger"
	"log"
	"os"
)

func main() {
	fmt.Println("Ledger service started")

	file, err := os.Open("budgets.json")
	if err != nil {
		log.Fatalf("failed to open budgets file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	if err := ledger.LoadBudgets(reader); err != nil {
		log.Fatalf("failed to load budgets: %v", err)
	}

	budgets := ledger.GetBudgets()

	fmt.Println("Budgets loaded:")
	for category, budget := range budgets {

		if err := ledger.CheckValid(&budget); err != nil {
			fmt.Printf("Invalid budget: %v\n", err)
			continue
		}

		fmt.Printf(
			"%s: limit=%d, period=%s\n",
			category,
			budget.Limit,
			budget.Period,
		)
	}

	fmt.Println()

	testTransactions := []ledger.Transaction{
		{
			Amount:      1500,
			Category:    "food",
			Description: "Lunch",
			Date:        "2026-10-05",
		},
		{
			Amount:      2000,
			Category:    "food",
			Description: "Groceries",
			Date:        "2026-10-05",
		},
		{
			Amount:      2000,
			Category:    "food",
			Description: "Restaurant",
			Date:        "2026-10-05",
		},
		{
			Amount:      1000,
			Category:    "transport",
			Description: "Taxi",
			Date:        "2026-10-05",
		},
	}

	for _, tx := range testTransactions {

		err2 := ledger.CheckValid(&tx)
		fmt.Print(err2)

		_, err := ledger.AddTransaction(tx)

		if errors.Is(err, ledger.ErrBudgetExceeded) {
			fmt.Printf(
				"Transaction rejected: category=%s amount=%d: %v\n",
				tx.Category,
				tx.Amount,
				err,
			)
			continue
		}

		if err != nil {
			fmt.Printf(
				"Failed to add transaction: %v\n",
				err,
			)
			continue
		}

		fmt.Printf(
			"Transaction added: category=%s amount=%d\n",
			tx.Category,
			tx.Amount,
		)
	}

	fmt.Println()
	fmt.Println("Saved transactions:")

	for _, tx := range ledger.ListTransactions() {
		fmt.Printf(
			"ID=%d Amount=%d Category=%s Description=%s Date=%s\n",
			tx.ID,
			tx.Amount,
			tx.Category,
			tx.Description,
			tx.Date,
		)
	}
}
