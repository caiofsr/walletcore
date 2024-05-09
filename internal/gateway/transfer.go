package gateway

import (
	"github.com/caiofsr/walletcore/internal/entity"
)

type TransferGateway interface {
	Create(transfer *entity.Transfer) error
}
