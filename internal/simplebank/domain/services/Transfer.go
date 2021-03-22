package services

import (
	"context"
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories"
)

type Transfer interface {
	Transfer(
		ctx context.Context,
		from model.AccountID,
		to model.AccountID,
		amount money.Money,
	) error
}

type transfer struct {
	accountRepository repositories.AccountRepository
}

func NewTransfer(accountRepository repositories.AccountRepository) Transfer {
	return &transfer{accountRepository: accountRepository}
}

func (s transfer) Transfer(
	ctx context.Context,
	from model.AccountID,
	to model.AccountID,
	amount money.Money,
) error {
	fromAccount, found, err := s.accountRepository.Get(ctx, from)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("from account not found")
	}

	toAccount, found, err := s.accountRepository.Get(ctx, to)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("to account not found")
	}

	err = fromAccount.Remove(amount)
	if err != nil {
		return err
	}

	toAccount.Add(amount)
	if err := s.accountRepository.Save(ctx, fromAccount); err != nil {
		return err
	}
	if err := s.accountRepository.Save(ctx, toAccount); err != nil {
		//todo need to revert fromAccount
		return err
	}

	return nil
}
