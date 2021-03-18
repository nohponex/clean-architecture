package model

import (
	"errors"
	"github.com/Rhymond/go-money"
)

type (
	AccountID string
	Account   interface {
		Add(amount money.Money)
		Remove(amount money.Money) error
		Balance() []money.Money
	}
	account struct {
		ID       AccountID
		Username string
		money    map[*money.Currency]*money.Money
	}
)

func NewAccount(ID AccountID) Account {
	return &account{
		ID:    ID,
		money: map[*money.Currency]*money.Money{},
	}
}

func (a account) Balance() []money.Money {
	list := []money.Money{}

	for _, m := range a.money {
		list = append(list, *m)
	}

	return list
}

func (a *account) Add(amount money.Money) {
	existing, exists := a.money[amount.Currency()]
	if !exists {
		a.money[amount.Currency()] = &amount
	} else {
		a.money[amount.Currency()], _ = existing.Add(&amount)
	}

}

func (a *account) Remove(amount money.Money) error {
	existing, exists := a.money[amount.Currency()]
	if !exists {
		return errors.New("account doesn't have requested currency")
	}

	if isLessThan, _ := existing.LessThan(&amount); isLessThan {
		return errors.New("requested amount is greater than current balance")
	}

	a.money[amount.Currency()], _ = existing.Subtract(&amount)
	return nil
}
