package domainservice

import (
	"context"
	"github.com/Rhymond/go-money"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_transfer(t *testing.T) {
	tests := []struct {
		name              string
		accountRepository repositories.AccountRepository
		amount            money.Money
		wantErr           bool
	}{
		{
			name:              "wantErr when ErrCannotTransferNegativeAmount",
			accountRepository: nil,
			amount:            *money.New(-500, "EUR"),
			wantErr:           true,
		},
		{
			name: "wantErr when 'to' account not found",
			accountRepository: func() repositories.AccountRepository {
				accountRepositoryMock := new(mocks.AccountRepository)

				accountFrom := model.NewAccount("from")

				{
					accountRepositoryMock.On(
						"Get",
						mock.Anything,
						model.AccountID("from"),
					).Return(
						accountFrom,
						true,
						nil,
					).Once()

					accountRepositoryMock.On(
						"Get",
						mock.Anything,
						model.AccountID("to"),
					).Return(
						nil,
						false,
						nil,
					).Once()
				}

				return accountRepositoryMock
			}(),
			amount:  *money.New(500, "EUR"),
			wantErr: true,
		},
		{
			name: "wantErr when from doesnt have sufficient amount",
			accountRepository: func() repositories.AccountRepository {
				accountRepositoryMock := new(mocks.AccountRepository)

				accountFrom := model.NewAccount("from")
				accountFrom.Add(*money.New(100, "EUR"))

				accountTo := model.NewAccount("to")

				{
					accountRepositoryMock.On(
						"Get",
						mock.Anything,
						model.AccountID("from"),
					).Return(
						accountFrom,
						true,
						nil,
					).Once()

					accountRepositoryMock.On(
						"Get",
						mock.Anything,
						model.AccountID("to"),
					).Return(
						accountTo,
						true,
						nil,
					).Once()
				}

				return accountRepositoryMock
			}(),
			amount:  *money.New(500, "EUR"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := transfer{
				accountRepository: tt.accountRepository,
			}
			if err := s.Transfer(context.Background(), "from", "to", tt.amount); (err != nil) != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_transfer_GivenFromAccountHasSufficientAmount(t *testing.T) {
	transferAmount := *money.New(600, "EUR")
	const fromID = "from"
	const toID = "to"

	accountRepositoryMock := new(mocks.AccountRepository)
	{
		accountFrom := model.NewAccount(fromID)
		accountFrom.Add(*money.New(1000, "EUR"))

		accountTo := model.NewAccount(toID)

		{
			accountRepositoryMock.On(
				"Get",
				mock.Anything,
				model.AccountID(fromID),
			).Return(
				accountFrom,
				true,
				nil,
			).Once()

			accountRepositoryMock.On(
				"Get",
				mock.Anything,
				model.AccountID(toID),
			).Return(
				accountTo,
				true,
				nil,
			).Once()

			accountRepositoryMock.On(
				"Save",
				mock.Anything,
				mock.MatchedBy(func(account model.Account) bool {
					if account.ID() != fromID {
						return false
					}

					return account.Balance()[0].Amount() == 400
				}),
			).Return(
				nil,
			).Once()

			accountRepositoryMock.On(
				"Save",
				mock.Anything,
				mock.MatchedBy(func(account model.Account) bool {
					if account.ID() != toID {
						return false
					}

					return account.Balance()[0].Amount() == transferAmount.Amount()
				}),
			).Return(
				nil,
			).Once()
		}
	}

	transfer := NewTransfer(accountRepositoryMock)
	err := transfer.Transfer(context.Background(), fromID, toID, transferAmount)
	require.NoError(t, err)

	accountRepositoryMock.AssertExpectations(t)
}
