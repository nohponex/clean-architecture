package adapters

import (
	"context"
	"errors"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/services"
	"strings"
)

type personPrefixAccountAccessService struct {
}

func NewPersonPrefixAccountAccessService() services.AccessService {
	return &personPrefixAccountAccessService{}
}

func (p personPrefixAccountAccessService) PersonHasAccessToAccount(
	ctx context.Context,
	person model.PersonID,
	account model.AccountID,
) (bool, error) {
	if len(person) == 0 || len(account) == 0 {
		return false, errors.New("empty IDs provided")
	}

	accountStartsWithPerson := strings.HasPrefix(string(account), string(person))
	return accountStartsWithPerson, nil
}
