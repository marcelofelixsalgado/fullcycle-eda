package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	AccountId string
	Amount    float64
	UpdatedAt time.Time
}

func NewBalance(accountId string, amount float64) (*Balance, error) {
	balance := &Balance{
		ID:        uuid.New().String(),
		AccountId: accountId,
		Amount:    amount,
		UpdatedAt: time.Now(),
	}
	err := balance.Validate()
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *Balance) Validate() error {
	if b.AccountId == "" {
		return errors.New("account id is required")
	}
	return nil
}
