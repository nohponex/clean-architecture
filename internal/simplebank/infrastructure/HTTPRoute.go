package infrastructure

import (
	"github.com/gorilla/mux"
	"github.com/nohponex/clean-architecture/internal/simplebank/adapters/httphandler"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"net/http"
)

func NewHTTPRouteCollection(
	account application.Account,
	transfer application.Transfer,
) http.Handler {
	withdrawHandler := httphandler.NewWithdrawHandler(account)
	topUpHandler := httphandler.NewTopUpHandler(account)
	openHandler := httphandler.NewOpenHandler(account)

	r := mux.NewRouter()
	r.Handle("/account/{accountID}/withdraw/{amount}", withdrawHandler).Methods(http.MethodGet)
	r.Handle("/account/{accountID}/topup/{amount}", topUpHandler).Methods(http.MethodGet)
	r.Handle("/account/{accountID}/open", openHandler).Methods(http.MethodGet, http.MethodPost)

	return r
}
