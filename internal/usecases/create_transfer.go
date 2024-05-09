package usecases

import (
	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/caiofsr/walletcore/internal/gateway"
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
}

func NewCreateTransferUseCase(tg gateway.TransferGateway, ag gateway.AccountGateway) *CreateTransferUseCase {
	return &CreateTransferUseCase{
		TransferGateway: tg,
		AccountGateway:  ag,
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

	return &CreateTransferOutputDTO{ID: transfer.ID}, nil
}
