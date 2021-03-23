package main

import (
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/messaging"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/persistence"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
	"net/http"
	"time"
)

func main() {
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

	accountUseCase := application.NewAccount(accountRepository, accessService, messenger)

	router := infrastructure.NewHTTPRouteCollection(accountUseCase, nil)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
