package model

import (
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_remove_GivenAccountNotHavingCurrency_ShouldFail(t *testing.T) {
	account := NewAccount("ABCD")

	err := account.Remove(*money.New(500, "EUR"))
	require.Error(t, err)
}

func Test_remove_GivenAccountHavingCurrencyButNotEnoughMoney_ShouldFail(t *testing.T) {
	account := NewAccount("ABCD")

	account.Add(*money.New(100, "EUR"))

	err := account.Remove(*money.New(500, "EUR"))
	require.Error(t, err)
}

func Test_remove_ShouldSubstract(t *testing.T) {
	account := NewAccount("ABCD")

	account.Add(*money.New(125, "EUR"))

	err := account.Remove(*money.New(50, "EUR"))
	require.NoError(t, err)

	balance := account.Balance()
	require.Len(t, balance, 1)

	{
		equals, err := balance[0].Equals(money.New(75, "EUR"))
		require.NoError(t, err)
		assert.True(t, equals)
	}
}
