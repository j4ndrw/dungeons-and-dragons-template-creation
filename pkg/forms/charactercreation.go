package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/alignment"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/appearance"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/lifestyle"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/meta"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/stats"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/vitals"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/components"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/validators"
)

func CharacterCreation(creature *creature.Creature) *huh.Form {
	modifierPlaceholder := "Format: <number> (+<modifier>) or <number>(-<modifier>). E.g. 14 (+3)"

	nameGroup := huh.
		NewGroup(
			huh.
				NewInput().
				Title("Enter the character's name").
				Value(&creature.Name),
		).
		Title("What is this creature's name?")

	appearanceGroup := huh.
		NewGroup(
			components.AppearanceInput(
				creature,
				appearance.Weight,
				components.WithValidator(validators.CheckNumeric),
			),
			components.AppearanceInput(
				creature,
				appearance.Height,
				components.WithValidator(validators.CheckNumeric),
			),
			components.AppearanceInput(
				creature,
				appearance.Skin,
			),
			components.AppearanceInput(
				creature,
				appearance.Hair,
			),
		).
		Title("Describe this creature.")

	statsGroup := huh.
		NewGroup(
			components.StatsInput(
				creature,
				stats.Wisdom,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
			components.StatsInput(
				creature,
				stats.Strength,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
			components.StatsInput(
				creature,
				stats.Dexterity,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
			components.StatsInput(
				creature,
				stats.Constitution,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
			components.StatsInput(
				creature,
				stats.Intelligence,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
			components.StatsInput(
				creature,
				stats.Charisma,
				components.WithValidator(validators.CheckNumericWithModifier),
				components.WithPlaceholder(modifierPlaceholder),
			),
		).Title("What are the stats of this creature?")

	metaGroup := huh.
		NewGroup(
			components.MetaConfirm(
				creature,
				meta.Initiative,
				components.WithAffirmative("+1"),
				components.WithNegative("-1"),
			),
			components.MetaSelect(creature, meta.Lifestyle, lifestyle.Lifestyles),
			components.MetaSelect(creature, meta.Alignment, alignment.Alignments),
		).Title("What is this creature's meta?")

	vitalsGroup := huh.
		NewGroup(
			components.VitalsInput(
				creature,
				vitals.HitPoints,
				components.WithValidator(validators.CheckNumeric),
			),
			components.VitalsInput(
				creature,
				vitals.ArmorClass,
				components.WithValidator(validators.CheckNumeric),
			),
			components.VitalsInput(
				creature,
				vitals.Speed,
				components.WithValidator(validators.CheckNumeric),
			),
		)

	return New(
		nameGroup,
		appearanceGroup,
		statsGroup,
		metaGroup,
		vitalsGroup,
	)
}
