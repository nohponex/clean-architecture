package application

import (
	"context"
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/services"
)

var ErrAccessNotAllowed = errors.New("access to requested resource not allowed")
var ErrAccountNotFound = errors.New("account not found")
var ErrAccountAlreadyExists = errors.New("account already exists")

type Account interface {
	Open(
		ctx context.Context,
		personID model.PersonID,
		accountID model.AccountID,
	) error

	Balance(
		ctx context.Context,
		personID model.PersonID,
		accountID model.AccountID,
	) ([]money.Money, error)

	Withdraw(
		ctx context.Context,
		personID model.PersonID,
		accountID model.AccountID,
		amount money.Money,
	) error

	TopUp(
		ctx context.Context,
		personID model.PersonID,
		accountID model.AccountID,
		amount money.Money,
	) error
}

type accountUseCase struct {
	accountRepository repositories.AccountRepository
	accessService     services.AccessService
}

func NewAccount(
	accountRepository repositories.AccountRepository,
	accessService services.AccessService,
) Account {
	return &accountUseCase{
		accountRepository: accountRepository,
		accessService:     accessService,
	}
}

//ErrAccountAlreadyExists
func (u accountUseCase) Open(ctx context.Context, personID model.PersonID, accountID model.AccountID) error {
	_, found, err := u.accountRepository.Get(ctx, accountID)
	if err != nil {
		return err
	}
	if found {
		return ErrAccountAlreadyExists
	}

	account := model.NewAccount(accountID)
	//todo what about access to this personID?

	return u.accountRepository.Save(ctx, account)
}

//@throws ErrAccessNotAllowed
func (u accountUseCase) Balance(
	ctx context.Context,
	personID model.PersonID,
	accountID model.AccountID,
) ([]money.Money, error) {
	{
		//Perform access checks
		hasAccess, err := u.accessService.PersonHasAccessToAccount(
			ctx,
			personID,
			accountID,
		)
		if err != nil {
			return nil, err
		}
		if !hasAccess {
			return nil, ErrAccessNotAllowed
		}
	}

	account, found, err := u.accountRepository.Get(ctx, accountID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrAccountNotFound
	}

	return account.Balance(), nil
}

//@throws ErrAccessNotAllowed
func (u accountUseCase) Withdraw(
	ctx context.Context,
	personID model.PersonID,
	accountID model.AccountID,
	amount money.Money,
) error {
	{
		//Perform access checks
		hasAccess, err := u.accessService.PersonHasAccessToAccount(
			ctx,
			personID,
			accountID,
		)
		if err != nil {
			return err
		}
		if !hasAccess {
			return ErrAccessNotAllowed
		}
	}

	account, found, err := u.accountRepository.Get(ctx, accountID)
	if err != nil {
		return err
	}
	if !found {
		return ErrAccountNotFound
	}

	if err := account.Remove(amount); err != nil {
		return err
	}

	return u.accountRepository.Save(ctx, account)
}

//@throws ErrAccessNotAllowed
func (u accountUseCase) TopUp(
	ctx context.Context,
	personID model.PersonID,
	accountID model.AccountID,
	amount money.Money,
) error {
	{
		//Perform access checks
		hasAccess, err := u.accessService.PersonHasAccessToAccount(
			ctx,
			personID,
			accountID,
		)
		if err != nil {
			return err
		}
		if !hasAccess {
			return ErrAccessNotAllowed
		}
	}

	account, found, err := u.accountRepository.Get(ctx, accountID)
	if err != nil {
		return err
	}
	if !found {
		return ErrAccountNotFound
	}

	account.Add(amount)

	return u.accountRepository.Save(ctx, account)
}
