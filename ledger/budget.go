package ledger

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
)

type Budget struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Period   string `json:"period,omitempty"`
}

var budgets = make(map[string]Budget)

func GetBudgets() map[string]Budget {
	return maps.Clone(budgets)
}

func SetBudget(b Budget) error {
	if err := CheckValid(&b); err != nil {
		return fmt.Errorf("invalid budget: %w", err)
	}
	budgets[b.Category] = b
	return nil
}

func LoadBudgets(r io.Reader) error {
	var loadedBudgets []Budget

	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedBudgets); err != nil {
		return fmt.Errorf("failed to decode budgets: %w", err)
	}

	for _, budget := range loadedBudgets {
		if err := SetBudget(budget); err != nil {
			return fmt.Errorf(
				"failed to load budget for category %q: %w",
				budget.Category,
				err,
			)
		}
	}

	return nil
}
