package usecases

import (
	"testing"

	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/caiofsr/walletcore/internal/event"
	"github.com/caiofsr/walletcore/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransferGatewayMock struct {
	mock.Mock
}

func (m *TransferGatewayMock) Create(transfer *entity.Transfer) error {
	args := m.Called(transfer)
	return args.Error(0)
}

func TestCreateTransferUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john@doe.com")
	account1, _ := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Mary Doe", "mary@doe.com")
	account2, _ := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	mockTransfer := &TransferGatewayMock{}
	mockTransfer.On("Create", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransferCreated()

	inputDTO := CreateTransferInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        float64(100),
	}

	uc := NewCreateTransferUseCase(mockTransfer, mockAccount, dispatcher, event)
	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockTransfer.AssertExpectations(t)
	mockTransfer.AssertNumberOfCalls(t, "Create", 1)
}
