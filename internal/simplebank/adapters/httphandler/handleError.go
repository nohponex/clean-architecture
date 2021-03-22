package httphandler

import "net/http"

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err.Error()))
}
