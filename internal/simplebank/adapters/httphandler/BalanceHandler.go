package httphandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"net/http"
)

type balanceHandler struct {
	account application.Account
}

func NewBalanceHandler(
	account application.Account,
) http.Handler {
	return &balanceHandler{
		account: account,
	}
}

func (h balanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID := vars["accountID"]

	balance, err := h.account.Balance(
		r.Context(),
		getPersonID(r),
		model.AccountID(accountID),
	)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(balance)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Authorization", "application/json")
	_, err = w.Write(response)
	if err != nil {
		panic(err)
	}
}
