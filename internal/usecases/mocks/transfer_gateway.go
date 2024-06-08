package mocks

import (
	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transfer) error {
	args := m.Called(transaction)
	return args.Error(0)
}
