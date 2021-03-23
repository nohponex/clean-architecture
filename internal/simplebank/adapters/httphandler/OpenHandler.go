package httphandler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"net/http"
)

type openHandler struct {
	account application.Account
}

func NewOpenHandler(
	account application.Account,
) http.Handler {
	return &openHandler{
		account: account,
	}
}

func (h openHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID := vars["accountID"]

	err := h.account.Open(
		r.Context(),
		getPersonID(r),
		model.AccountID(accountID),
	)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Add("Location", fmt.Sprintf("/account/%s", accountID))
	w.WriteHeader(http.StatusCreated)
}
