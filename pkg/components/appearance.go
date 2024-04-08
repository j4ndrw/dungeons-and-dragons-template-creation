package components

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/validators"
)

func Appearance(
	creature *creature.Creature,
	key string,
	validatorFunctions ...validators.ValidatorType,
) *huh.Input {
	component := huh.NewInput().
		Title(key).
		Value(creature.Appearance.ParseTable__[key].(*string))

	for _, validatorFunction := range validatorFunctions {
		component = component.Validate(validatorFunction)
	}

	return component
}
