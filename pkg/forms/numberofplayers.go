package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/validators"
)

func NumberOfPlayers(numberOfPlayers *string) *huh.Form {
	return New(
		huh.NewGroup(
			huh.NewInput().
				Title("How many players are at the table?").
				Prompt("> ").
				Value(numberOfPlayers).
				Validate(validators.NumberOfPlayers),
		))
}
