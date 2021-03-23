package httphandler

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/application/mocks"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_balanceHandler_ServeHTTP(t *testing.T) {
	accountStub := new(mocks.Account)
	{
		accountStub.On(
			"Balance",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			[]money.Money{
				*money.New(500, "EUR"),
				*money.New(1000, "USD"),
			},
			nil,
		)
	}

	res := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		"/account/ABCD",
		nil,
	)

	handler := NewBalanceHandler(accountStub)
	handler.ServeHTTP(res, req)

	response, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
}
