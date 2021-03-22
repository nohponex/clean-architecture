package terminal

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
)

type commandDispatcher struct {
	commands []command
}

func NewCommandDispatcher(commands ...command) *commandDispatcher {
	return &commandDispatcher{commands: commands}
}

func (d commandDispatcher) command(
	ctx context.Context,
	personID model.PersonID,
	commandParts []string,
) (handled bool, err error) {
	for _, cmd := range d.commands {
		h, err := cmd.command(ctx, personID, commandParts)
		if h == true {
			return true, err
		}
	}

	return false, nil
}
