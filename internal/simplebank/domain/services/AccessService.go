package services

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type AccessService interface {
	PersonHasAccessToAccount(
		ctx context.Context,
		person model.PersonID,
		account model.AccountID,
	) (bool, error)
}
