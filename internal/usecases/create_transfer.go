package usecases

import (
	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/caiofsr/walletcore/internal/gateway"
	"github.com/caiofsr/walletcore/pkg/events"
)

type CreateTransferInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransferOutputDTO struct {
	ID string
}

type CreateTransferUseCase struct {
	TransferGateway gateway.TransferGateway
	AccountGateway  gateway.AccountGateway
	EventDispatcher events.EventDispatcherInterface
	TransferCreated events.EventInterface
}

func NewCreateTransferUseCase(
	tg gateway.TransferGateway,
	ag gateway.AccountGateway,
	ed events.EventDispatcherInterface,
	transferCreated events.EventInterface,
) *CreateTransferUseCase {
	return &CreateTransferUseCase{
		TransferGateway: tg,
		AccountGateway:  ag,
		EventDispatcher: ed,
		TransferCreated: transferCreated,
	}
}

func (uc *CreateTransferUseCase) Execute(input CreateTransferInputDTO) (*CreateTransferOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transfer, err := entity.NewTransfer(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransferGateway.Create(transfer)
	if err != nil {
		return nil, err
	}

	output := &CreateTransferOutputDTO{ID: transfer.ID}

	uc.TransferCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransferCreated)

	return output, nil
}
