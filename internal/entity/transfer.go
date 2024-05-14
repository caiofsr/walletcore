package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransfer(accountFrom, accountTo *Account, amount float64) (*Transfer, error) {
	transfer := &Transfer{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := transfer.Validate()
	if err != nil {
		return nil, err
	}

	transfer.Commit()

	return transfer, nil
}

func (t *Transfer) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transfer) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient fund")
	}

	if t.AccountFrom.ID == t.AccountTo.ID {
		return errors.New("is not possible to transfer to same account")
	}

	return nil
}
