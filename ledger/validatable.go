package main

import (
	"errors"
	"time"
)

var ErrAmountLessZero = errors.New("amount must be >0")
var ErrCategoryEmpty = errors.New("category must not be empty")
var ErrDateInvalid = errors.New("date must be in YYYY-MM-DD format")
var ErrLimitNegative = errors.New("limit must be a positive value")

type Validatable interface {
	Validate() error
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return ErrAmountLessZero
	}
	if t.Category == "" {
		return ErrCategoryEmpty
	}
	_, err := time.Parse("2006-01-02", t.Date)
	if err != nil {
		return ErrDateInvalid
	}

	return nil
}

func (b *Budget) Validate() error {
	if b.Limit <= 0 {
		return ErrLimitNegative
	}
	if b.Category == "" {
		return ErrCategoryEmpty
	}

	return nil
}

func CheckValid(v Validatable) error {
	return v.Validate()
}
