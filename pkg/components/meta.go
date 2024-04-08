package components

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
)

func MetaInput(
	creature *creature.Creature,
	key string,
	options ...Option,
) *huh.Input {
	component := huh.NewInput().
		Title(key).
		Value(creature.Meta.ParseTable__[key].(*string))

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

func MetaSelect(
	creature *creature.Creature,
	key string,
	selectables []string,
	options ...Option,
) *huh.Select[string] {
	var selectOptions []huh.Option[string]
	for _, selectable := range selectables {
		selectOptions = append(selectOptions, huh.NewOption(selectable, selectable))
	}
	component := huh.NewSelect[string]().
		Title(key).
		Options(selectOptions...).
		Value(creature.Meta.ParseTable__[key].(*string))

	for _, option := range options {
		if option.Validator != nil {
			component = component.Validate(*option.Validator)
		}
	}

	return component

}
