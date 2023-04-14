package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewBalance(t *testing.T) {
	var amount float64 = 1000.0
	balance, err := NewBalance("11", amount)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, "11", balance.AccountId)
	assert.Equal(t, amount, balance.Amount)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	var amount float64 = 1000.0
	balance, err := NewBalance("", amount)
	assert.NotNil(t, err)
	assert.Nil(t, balance)
}
