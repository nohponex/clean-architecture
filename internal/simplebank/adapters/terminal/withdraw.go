package terminal

import (
	"context"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"strconv"
	"strings"
)

const withdrawCommandName = "withdraw"

type withdraw struct {
	account application.Account
}

func (withdraw) help() string {
	return fmt.Sprintf("%s {accountID} {amount}", withdrawCommandName)
}

func (c withdraw) command(
	ctx context.Context,
	personID model.PersonID,
	commandParts []string,
) (handled bool, err error) {
	if strings.ToLower(commandParts[0]) != withdrawCommandName || len(commandParts) != 3 {
		return false, nil
	}

	handled = true

	amountAsInt, err := strconv.Atoi(commandParts[2])
	if err != nil {
		return handled, err
	}

	m := money.New(int64(amountAsInt), "EUR")
	if err := c.account.Withdraw(ctx, personID, model.AccountID(commandParts[1]), *m); err != nil {
		return handled, err
	}

	return handled, nil
}
