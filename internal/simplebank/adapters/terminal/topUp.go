package terminal

import (
	"context"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"strconv"
	"strings"
)

type topUp struct {
	account application.Account
}

func (c topUp) command(
	ctx context.Context,
	personID model.PersonID,
	commandParts []string,
) (handled bool, err error) {
	if strings.ToLower(commandParts[0]) != "topup" || len(commandParts) != 3 {
		return false, nil
	}

	handled = true

	amountAsInt, err := strconv.Atoi(commandParts[2])
	if err != nil {
		return handled, err
	}

	m := money.New(int64(amountAsInt), "EUR")
	if err := c.account.TopUp(ctx, personID, model.AccountID(commandParts[1]), *m); err != nil {
		return handled, err
	}

	return handled, nil
}
