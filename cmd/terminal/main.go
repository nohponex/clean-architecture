package main

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/persistence"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/terminal"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
)

func main() {
	ctx := context.Background()

	accountRepository := persistence.NewInMemoryAccountRepository()
	accessService := adapters.NewPersonPrefixAccountAccessService()

	accountUseCase := application.NewAccount(accountRepository, accessService)

	terminal.TerminalApplication(ctx, accountUseCase)
}
