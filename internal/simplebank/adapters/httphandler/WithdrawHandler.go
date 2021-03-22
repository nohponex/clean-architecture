package httphandler

import (
	"github.com/Rhymond/go-money"
	"github.com/gorilla/mux"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"net/http"
	"strconv"
)

type withdrawHandler struct {
	account application.Account
}

func NewWithdrawHandler(
	account application.Account,
) http.Handler {
	return &withdrawHandler{
		account: account,
	}
}

func (h withdrawHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID := vars["accountID"]
	currency := "EUR"
	amount := vars["amount"]
	amountAsInt, err := strconv.Atoi(amount)
	if err != nil {
		handleError(w, err)
		return
	}

	moneyAmount := money.New(int64(amountAsInt), currency)

	err = h.account.Withdraw(
		r.Context(),
		getPersonID(r),
		model.AccountID(accountID),
		*moneyAmount,
	)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
