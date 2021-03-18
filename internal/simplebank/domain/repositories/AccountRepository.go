package repositories

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type AccountRepository interface {
	Get (ctx context.Context, id model.AccountID) (account model.Account, found bool, err error)
	Save (ctx context.Context, account model.Account) error
}
