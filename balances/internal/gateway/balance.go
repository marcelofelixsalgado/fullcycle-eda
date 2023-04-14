package gateway

import "github.com.br/marcelofelixsalgado/fullcycle-eda/balances/internal/entity"

type BalanceGateway interface {
	Get(accountId string) (*entity.Balance, error)
	Insert(balance *entity.Balance) error
	Update(accountId string, amount float64) error
}
