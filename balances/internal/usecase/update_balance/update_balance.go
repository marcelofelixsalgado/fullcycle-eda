package updatebalance

import (
	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	AccountId string
	Amount    float64
}

type UpdateBalanceOutputDTO struct {
	AccountId string
	Amount    float64
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(balanceGateway gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) (*UpdateBalanceOutputDTO, error) {

	err := uc.BalanceGateway.Update(input.AccountId, input.Amount)
	if err != nil {
		return nil, err
	}

	output := &UpdateBalanceOutputDTO{
		AccountId: input.AccountId,
		Amount:    input.Amount,
	}
	return output, nil
}
