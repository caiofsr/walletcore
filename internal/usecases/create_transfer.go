package usecases

import (
	"context"

	"github.com/caiofsr/walletcore/internal/entity"
	"github.com/caiofsr/walletcore/internal/gateway"
	"github.com/caiofsr/walletcore/pkg/events"
	"github.com/caiofsr/walletcore/pkg/uow"
)

type CreateTransferInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransferOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransferUseCase struct {
	Uow             uow.UowInterface
	EventDispatcher events.EventDispatcherInterface
	TransferCreated events.EventInterface
}

func NewCreateTransferUseCase(
	uow uow.UowInterface,
	ed events.EventDispatcherInterface,
	transferCreated events.EventInterface,
) *CreateTransferUseCase {
	return &CreateTransferUseCase{
		Uow:             uow,
		EventDispatcher: ed,
		TransferCreated: transferCreated,
	}
}

func (uc *CreateTransferUseCase) Execute(ctx context.Context, input CreateTransferInputDTO) (*CreateTransferOutputDTO, error) {
	output := &CreateTransferOutputDTO{}

	err := uc.Uow.Do(ctx, func(_ *uow.UnitOfWork) error {
		accountRepository := uc.getAccountRepository(ctx)
		transferRepository := uc.getTransferRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}

		transfer, err := entity.NewTransfer(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = transferRepository.Create(transfer)
		if err != nil {
			return err
		}

		output.ID = transfer.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount

		return nil
	})
	if err != nil {
		return nil, err
	}

	uc.TransferCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransferCreated)

	return output, nil
}

func (uc *CreateTransferUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}

	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransferUseCase) getTransferRepository(ctx context.Context) gateway.TransferGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransferDB")
	if err != nil {
		panic(err)
	}

	return repo.(gateway.TransferGateway)
}
