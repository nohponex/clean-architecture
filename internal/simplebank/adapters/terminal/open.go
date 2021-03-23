package terminal

import (
	"context"
	"fmt"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"strings"
)

const openCommandName = "open"

type open struct {
	account application.Account
}

func (open) help() string {
	return fmt.Sprintf("%s {accountID}", openCommandName)
}

func (c open) command(
	ctx context.Context,
	personID model.PersonID,
	commandParts []string,
) (handled bool, err error) {
	if strings.ToLower(commandParts[0]) != openCommandName || len(commandParts) != 2 {
		return false, nil
	}

	handled = true

	if err := c.account.Open(ctx, personID, model.AccountID(commandParts[1])); err != nil {
		return handled, err
	}

	return handled, nil
}
