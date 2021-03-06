package terminal

import (
	"bufio"
	"context"
	"fmt"
	"github.com/nohponex/clean-architecture/internal/simplebank/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"os"
	"strings"
)

func TerminalApplication(
	ctx context.Context,
	account application.Account,
) {
	dispatcher := NewCommandDispatcher(
		&open{account: account},
		&balance{account: account},
		&withdraw{account: account},
		&topUp{account: account},
	)

	reader := bufio.NewReader(os.Stdin)

	personID := func() model.PersonID {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("This is simplebank")
		fmt.Print("Enter your Identity: ")
		personRaw, _ := reader.ReadString('\n')
		personRaw = strings.TrimSpace(personRaw)
		fmt.Printf("Welcome %q\n", personRaw)

		return model.PersonID(personRaw)
	}()

	{
		fmt.Println("Available commands:")
		for _, c := range dispatcher.commands {
			fmt.Println(c.help())
		}
		fmt.Println()
	}

	for {
		fmt.Println("Command: ")
		command, _ := reader.ReadString('\n')
		commandParts := strings.Split(
			strings.TrimSpace(command),
			" ",
		)

		handled, err := dispatcher.command(ctx, personID, commandParts)
		if !handled {
			fmt.Println("Error: Unknown command")
		}

		if err != nil {
			fmt.Println("Error: " + err.Error())
		}

		fmt.Println()
	}
}
