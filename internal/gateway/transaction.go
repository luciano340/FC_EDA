package gateway

import "github.com/luciano340/FC_EDA/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
