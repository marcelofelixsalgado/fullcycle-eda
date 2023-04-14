package getbalance

import (
	"github.com.br/marcelofelixsalgado/fullcycle-eda/balances/internal/gateway"
)

type GetBalanceInputDTO struct {
	AccountId string
}

type GetBalanceOutputDTO struct {
	AccountId string
	Amount    float64
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(balanceGateway gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *GetBalanceUseCase) Execute(input GetBalanceInputDTO) (*GetBalanceOutputDTO, error) {

	balance, err := uc.BalanceGateway.Get(input.AccountId)
	if err != nil {
		return nil, err
	}

	output := &GetBalanceOutputDTO{
		AccountId: balance.AccountId,
		Amount:    balance.Amount,
	}
	return output, nil
}
