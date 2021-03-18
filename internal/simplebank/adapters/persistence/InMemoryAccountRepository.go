package persistence

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories"
)

type inMemoryAccountRepository struct {
	memory map[model.AccountID]model.Account
}

func NewInMemoryAccountRepository() repositories.AccountRepository {
	return &inMemoryAccountRepository{
		memory: map[model.AccountID]model.Account{},
	}
}

func (i *inMemoryAccountRepository) Get(
	ctx context.Context,
	id model.AccountID,
) (account model.Account, found bool, err error) {
	account, found = i.memory[id]

	return account, found, nil
}

func (i *inMemoryAccountRepository) Save(
	ctx context.Context,
	account model.Account,
) error {
	i.memory[account.ID()] = account
	return nil
}
