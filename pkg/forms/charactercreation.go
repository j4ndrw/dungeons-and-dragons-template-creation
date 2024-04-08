package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/appearance"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/meta"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/stats"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/vitals"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/components"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/validators"
)

func CharacterCreation(creature *creature.Creature) *huh.Form {
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
			components.Appearance(creature, appearance.Weight, validators.CheckNumeric),
			components.Appearance(creature, appearance.Height, validators.CheckNumeric),
			components.Appearance(creature, appearance.Skin),
			components.Appearance(creature, appearance.Hair),
		).
		Title("Describe this creature.")

	statsGroup := huh.
		NewGroup(
			components.Stats(creature, stats.Wisdom, validators.CheckNumeric),
			components.Stats(creature, stats.Strength, validators.CheckNumeric),
			components.Stats(creature, stats.Dexterity, validators.CheckNumeric),
			components.Stats(creature, stats.Constitution, validators.CheckNumeric),
			components.Stats(creature, stats.Intelligence, validators.CheckNumeric),
			components.Stats(creature, stats.Charisma, validators.CheckNumeric),
		).Title("What are the stats of this creature?")

	metaGroup := huh.
		NewGroup(
			components.Meta(creature, meta.Initiative, validators.CheckNumeric),
			components.Meta(creature, meta.Lifestyle),
			components.Meta(creature, meta.Alignment),
		).Title("What is this creature's meta?")

	vitalsGroup := huh.
		NewGroup(
			components.Vitals(creature, vitals.HitPoints, validators.CheckNumeric),
			components.Vitals(creature, vitals.ArmorClass, validators.CheckNumeric),
			components.Vitals(creature, vitals.Speed, validators.CheckNumeric),
		)

	return New(
		nameGroup,
		appearanceGroup,
		statsGroup,
		metaGroup,
		vitalsGroup,
	)
}
