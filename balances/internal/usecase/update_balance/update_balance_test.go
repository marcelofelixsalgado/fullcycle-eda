package updatebalance_test

import (
	"testing"

	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/entity"
	updatebalance "github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/usecase/update_balance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceRepositoryMock struct {
	mock.Mock
}

func (m *BalanceRepositoryMock) Insert(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceRepositoryMock) Get(id string) (*entity.Balance, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *BalanceRepositoryMock) Update(accountId string, amount float64) error {
	args := m.Called(accountId, amount)
	return args.Error(0)
}

func TestUpdateBalanceUseCase_Execute(t *testing.T) {
	m := &BalanceRepositoryMock{}
	m.On("Update", mock.Anything, mock.Anything).Return(nil)
	uc := updatebalance.NewUpdateBalanceUseCase(m)

	var amount float64 = 100.0

	output, err := uc.Execute(updatebalance.UpdateBalanceInputDTO{
		AccountId: "11",
		Amount:    amount,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "11", output.AccountId)
	assert.Equal(t, amount, output.Amount)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Update", 1)
}
