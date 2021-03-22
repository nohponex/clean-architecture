package httphandler

import (
	"fmt"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"net/http"
	"strings"
)

func getPersonID(r *http.Request) model.PersonID {
	authorization := r.Header.Get("Authorization")

	if len(authorization) == 0 {
		return ""
	}

	var personRaw string
	n, err := fmt.Fscanf(strings.NewReader(authorization), "Person %s", &personRaw)
	if err != nil || n == 0 {
		return ""
	}
	return model.PersonID(personRaw)
}
