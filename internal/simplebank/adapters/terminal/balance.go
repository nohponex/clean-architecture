package terminal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"strings"
)

const balanceCommandName = "balance"

type balance struct {
	account application.Account
}

func (balance) help() string {
	return fmt.Sprintf("%s {accountID}", balanceCommandName)
}

func (c balance) command(
	ctx context.Context,
	personID model.PersonID,
	commandParts []string,
) (handled bool, err error) {
	if strings.ToLower(commandParts[0]) != balanceCommandName || len(commandParts) != 2 {
		return false, nil
	}

	err = func() error {
		balance, err := c.account.Balance(ctx, personID, model.AccountID(commandParts[1]))
		if err != nil {
			return err
		}
		fmt.Println("balance:")

		JSON, err := json.Marshal(balance)
		if err != nil {
			return err
		}

		fmt.Println(string(JSON))

		return nil
	}()

	return true, err
}
