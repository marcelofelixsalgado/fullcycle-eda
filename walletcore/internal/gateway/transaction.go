package gateway

import "github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
