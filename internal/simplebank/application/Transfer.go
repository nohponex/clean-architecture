package application

import (
	"context"
	"github.com/Rhymond/go-money"

	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type (
	Transfer interface {
		Transfer(
			ctx context.Context,
			person model.Person,
			from model.Account,
			to model.Account,
			amount money.Money,
		) error
	}
)
