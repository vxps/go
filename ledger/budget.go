package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Budget struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Period   string `json:"period,omitempty"`
}

var budgets = make(map[string]Budget)

func SetBudget(b Budget) {
	budgets[b.Category] = b
}

func LoadBudgets(r io.Reader) error {
	var loadedBudgets []Budget

	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedBudgets); err != nil {
		return fmt.Errorf("failed to decode budgets: %w", err)
	}

	for _, budget := range loadedBudgets {
		if budget.Category == "" {
			return fmt.Errorf("budget category cannot be empty")
		}

		if budget.Limit <= 0 {
			return fmt.Errorf(
				"budget limit for category %q must be positive",
				budget.Category,
			)
		}

		SetBudget(budget)
	}

	return nil
}