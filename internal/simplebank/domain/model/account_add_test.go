package model

import (
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_account_Add_GivenEmptyBalance_ThenBalanceShouldHoldOnlyAddedAmount(t *testing.T) {
	account := NewAccount("ABCD")

	firstAmount := money.New(100, "EUR")

	account.Add(*firstAmount)

	balance := account.Balance()
	require.Len(t, balance, 1)

	equals, err := balance[0].Equals(firstAmount)
	require.NoError(t, err)
	assert.True(t, equals)
}

func Test_account_Add_GivenNonEmptyBalance_ThenBalanceShouldHoldBothCurrencies(t *testing.T) {
	account := NewAccount("ABCD")

	usdAmount := money.New(100, "USD")
	eurAmount := money.New(100, "EUR")

	account.Add(*usdAmount)
	account.Add(*eurAmount)

	balance := account.Balance()
	require.Len(t, balance, 2)

	usdBalance, eurBalance := func() (money.Money, money.Money) {
		if balance[0].Currency() == usdAmount.Currency() {
			return balance[0], balance[1]
		}
		return balance[1], balance[0]
	}()

	{
		equals, err := usdBalance.Equals(usdAmount)
		require.NoError(t, err)
		assert.True(t, equals)
	}
	{
		equals, err := eurBalance.Equals(eurAmount)
		require.NoError(t, err)
		assert.True(t, equals)
	}
}

func Test_account_Add_GivenAlreadyHadCurrency_ThenCurrencyShouldBeAddedToExisting(t *testing.T) {
	account := NewAccount("ABCD")

	initialAmount := money.New(100, "EUR")
	anotherAmount := money.New(200, "EUR")

	account.Add(*initialAmount)
	account.Add(*anotherAmount)

	balance := account.Balance()
	require.Len(t, balance, 1)

	{
		equals, err := balance[0].Equals(
			money.New(300, "EUR"),
		)
		require.NoError(t, err)
		assert.True(t, equals)
	}
}
