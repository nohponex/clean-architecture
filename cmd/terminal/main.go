package main

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/messaging"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/persistence"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/terminal"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
)

func main() {
	ctx := context.Background()

	config, err := infrastructure.NewConfigFromEnvironmental()
	if err != nil {
		panic(err)
	}

	accountRepository := persistence.NewInMemoryAccountRepository()
	accessService := adapters.NewPersonPrefixAccountAccessService()

	messenger, err := messaging.NewRabbitMQMessaging(config)
	if err != nil {
		panic(err)
	}

	accountUseCase := application.NewAccount(
		accountRepository,
		accessService,
		messenger,
	)

	terminal.TerminalApplication(ctx, accountUseCase)
}
