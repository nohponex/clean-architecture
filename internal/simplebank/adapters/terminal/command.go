package terminal

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type command interface {
	command(
		ctx context.Context,
		personID model.PersonID,
		commandParts []string,
	) (handled bool, err error)

	help() string
}
