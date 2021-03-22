package main

import (
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/persistence"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
	"net/http"
	"time"
)

func main() {
	accountRepository := persistence.NewInMemoryAccountRepository()
	accessService := adapters.NewPersonPrefixAccountAccessService()

	accountUseCase := application.NewAccount(accountRepository, accessService)

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
