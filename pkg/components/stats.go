package components

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
)

func StatsInput(
	creature *creature.Creature,
	key string,
	options ...Option,
) *huh.Input {
	component := huh.NewInput().
		Title(key).
		Value(creature.Stats.ParseTable__[key].(*string))

	for _, option := range options {
		if option.Validator != nil {
			component = component.Validate(*option.Validator)
		}
		if option.Suggestions != nil {
			component = component.Suggestions(*option.Suggestions)
		}
		if option.Placeholder != nil {
			component = component.Placeholder(*option.Placeholder)
		}
	}

	return component
}
