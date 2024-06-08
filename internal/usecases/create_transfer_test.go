package usecases

import (
	"context"
	"testing"

	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/caiofsr/walletcore/internal/event"
	"github.com/caiofsr/walletcore/internal/usecases/mocks"
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

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransferCreated()
	ctx := context.Background()

	inputDTO := CreateTransferInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        float64(100),
	}

	uc := NewCreateTransferUseCase(mockUow, dispatcher, event)
	output, err := uc.Execute(ctx, inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
