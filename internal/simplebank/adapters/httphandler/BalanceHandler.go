package httphandler

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/gorilla/mux"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/thoas/go-funk"
	"net/http"
	"strings"
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

	{
		//very unorthodox

		list := funk.Map(
			balance,
			func(money money.Money) string {
				return fmt.Sprintf(`{"currency": "%s", "amount": %d}`, money.Currency().Code, money.Amount())
			},
		).([]string)

		json := fmt.Sprintf(
			"[%s]",
			strings.Join(list, ","),
		)

		w.Header().Add("Authorization", "application/json")
		_, err = w.Write([]byte(json))
		if err != nil {
			panic(err)
		}
	}
}
