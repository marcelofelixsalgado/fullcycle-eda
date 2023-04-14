package getbalance_test

import (
	"testing"

	"github.com.br/marcelofelixsalgado/fullcycle-eda/balances/internal/entity"
	getbalance "github.com.br/marcelofelixsalgado/fullcycle-eda/balances/internal/usecase/get_balance"
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

func TestGetBalanceUseCase_Execute(t *testing.T) {
	var amount float64 = 1000.0
	balance, _ := entity.NewBalance("11", amount)

	m := &BalanceRepositoryMock{}
	m.On("Get", mock.Anything).Return(balance, nil)
	uc := getbalance.NewGetBalanceUseCase(m)

	output, err := uc.Execute(getbalance.GetBalanceInputDTO{
		AccountId: "11",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "11", output.AccountId)
	assert.Equal(t, output.Amount, amount)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Get", 1)
}
