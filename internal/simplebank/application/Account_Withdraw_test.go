package application

import (
	"context"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	repositoriesMocks "github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories/mocks"
	servicesMocks "github.com/nohponex/clean-architecture/internal/simplebank/domain/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_Withdraw_Given_PersonDoesntHaveAccess_ThenShouldFail(t *testing.T) {
	accessServiceMock := getAnAllowedNoneAccessServiceStub()

	useCase := NewAccount(
		nil,
		accessServiceMock,
	)

	err := useCase.Withdraw(
		context.Background(),
		"abcd",
		"ABCD",
		*money.New(50, "EUR"),
	)
	assert.ErrorIs(t, err, ErrAccessNotAllowed)
}

func Test_Withdraw_GivenAccountNotFound_ThenShouldFail(t *testing.T) {
	accountRepositoryStub := new(repositoriesMocks.AccountRepository)
	{
		accountRepositoryStub.On(
			"Get",
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
			false,
			nil,
		).Once()
	}

	useCase := NewAccount(
		accountRepositoryStub,
		getAnAllowedAllAccessServiceStub(),
	)

	err := useCase.Withdraw(
		context.Background(),
		"abcd",
		"SomeAccountID",
		*money.New(500, "EUR"),
	)
	assert.ErrorIs(t, err, ErrAccountNotFound)

	accountRepositoryStub.AssertExpectations(t)
}

func Test_Withdraw_Given_ThenShouldWithdrawAndSave(t *testing.T) {
	accountRepositoryMock := new(repositoriesMocks.AccountRepository)

	const fakeAccountID = model.AccountID("ABCD")

	{
		fakeAccount := model.NewAccount(fakeAccountID)

		accountRepositoryMock.On(
			"Get",
			mock.Anything,
			fakeAccountID,
		).Return(
			fakeAccount,
			true,
			nil,
		).Once()

		accountRepositoryMock.On(
			"Save",
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		).Once()
	}

	useCase := NewAccount(
		accountRepositoryMock,
		getAnAllowedAllAccessServiceStub(),
	)

	amount := *money.New(50, "EUR")

	err := useCase.Withdraw(
		context.Background(),
		"abcd",
		fakeAccountID,
		amount,
	)
	assert.NoError(t, err)

	accountRepositoryMock.AssertExpectations(t)
}

func getAnAllowedAllAccessServiceStub() *servicesMocks.AccessService {
	stub := new(servicesMocks.AccessService)

	stub.On(
		"PersonHasAccessToAccount",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		true,
		nil,
	)

	return stub
}

func getAnAllowedNoneAccessServiceStub() *servicesMocks.AccessService {
	stub := new(servicesMocks.AccessService)

	stub.On(
		"PersonHasAccessToAccount",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		false,
		nil,
	)

	return stub
}
