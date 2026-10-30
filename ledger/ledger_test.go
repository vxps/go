package ledger

import (
	"errors"
	"fmt"
	"testing"
)

func TestTransactions(t *testing.T) {
	var tests = []struct {
		name    string
		tx      Transaction
		wantErr bool
	}{
		{
			"test",
			Transaction{
				Amount:   100,
				Category: "food",
				Date:     "2023-01-01",
			},
			false,
		},
		{
			"2 test",
			Transaction{
				Amount:   0,
				Category: "food",
				Date:     "2023-01-01",
			},
			true,
		},
		{
			"3 test",
			Transaction{
				Amount:   -100,
				Category: "food",
				Date:     "2023-01-01",
			},
			true,
		},
		{
			"4 test",
			Transaction{
				Amount:   100,
				Category: "",
				Date:     "2023-01-01",
			},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.tx.Validate()

			if (err != nil) != test.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}

}

func TestBudgets(t *testing.T) {
	var tests = []struct {
		name    string
		budget  Budget
		wantErr bool
	}{
		{
			"test 1",
			Budget{
				Category: "food",
				Limit:    1000,
				Period:   "monthly",
			},
			false,
		},
		{
			"test 2",
			Budget{
				Category: "food",
				Limit:    0,
				Period:   "monthly",
			},
			true,
		},
		{
			"test 3",
			Budget{
				Category: "food",
				Limit:    -1000,
				Period:   "monthly",
			},
			true,
		},
		{
			"test 4",
			Budget{
				Category: "",
				Limit:    1000,
				Period:   "monthly",
			},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.budget.Validate()
			if (err != nil) != test.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestAddTransaction(t *testing.T) {
	Reset()
	t.Cleanup(Reset)
	err := SetBudget(Budget{
		Category: "food",
		Limit:    5000,
		Period:   "monthly",
	})
	if err != nil {
		t.Fatalf("SetBudget() failed: %v", err)
	}

	t.Run("add transaction 1 - 3000", func(t *testing.T) {
		_, err := AddTransaction(Transaction{
			Amount:   3000,
			Category: "food",
			Date:     "2023-01-01"})

		if err != nil {
			t.Errorf("AddTransaction() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("add transaction 2 - 3000", func(t *testing.T) {
		_, err := AddTransaction(Transaction{
			Amount:   3000,
			Category: "food",
			Date:     "2023-01-01"})
		if !errors.Is(err, ErrBudgetExceeded) {
			t.Errorf("AddTransaction() another error = %v, wantErr %v", err, true)
		}
		fmt.Printf("%v", err)
	})

	t.Run("test list transactions", func(t *testing.T) {
		tr := ListTransactions()
		if len(tr) != 1 {
			t.Error("transaction count mismatch.")
		}
	})
}
