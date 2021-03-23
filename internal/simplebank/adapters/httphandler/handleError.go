package httphandler

import (
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch err {
	case application.ErrAccessNotAllowed:
		w.WriteHeader(http.StatusForbidden)
	case application.ErrAccountNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	_, _ = w.Write([]byte(err.Error()))
}
