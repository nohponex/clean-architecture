package application

import (
	"context"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/domainservice"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/services"

	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type (
	Transfer interface {
		Transfer(
			ctx context.Context,
			personID model.PersonID,
			from model.AccountID,
			to model.AccountID,
			amount money.Money,
		) error
	}
)

type transfer struct {
	transferService domainservice.Transfer
	accessService   services.AccessService
}

func NewTransfer(
	transferService domainservice.Transfer,
	accessService services.AccessService,
) Transfer {
	return &transfer{transferService: transferService, accessService: accessService}
}

func (u *transfer) Transfer(
	ctx context.Context,
	personID model.PersonID,
	from model.AccountID,
	to model.AccountID,
	amount money.Money,
) error {
	{
		hasAccess, err := u.accessService.PersonHasAccessToAccount(
			ctx,
			personID,
			from,
		)
		if err != nil {
			return err
		}
		if !hasAccess {
			return ErrAccessNotAllowed
		}
	}

	return u.transferService.Transfer(
		ctx,
		from,
		to,
		amount,
	)
}
