package creature

import (
	"fmt"

	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/appearance"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/meta"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/stats"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/vitals"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/utils"
)

type Creature struct {
	Name       string
	Stats      *stats.Stats
	Appearance *appearance.Appearance
	Meta       *meta.Meta
	Vitals     *vitals.Vitals
}

func Empty() *Creature {
	return &Creature{
		Name: "",
		Stats: &stats.Stats{
			ParseTable__: map[string]interface{}{
				stats.Wisdom:       utils.AsAnyPtr("0"),
				stats.Intelligence: utils.AsAnyPtr("0"),
				stats.Strength:     utils.AsAnyPtr("0"),
				stats.Dexterity:    utils.AsAnyPtr("0"),
				stats.Constitution: utils.AsAnyPtr("0"),
				stats.Charisma:     utils.AsAnyPtr("0"),
			},
		},
		Appearance: &appearance.Appearance{
			ParseTable__: map[string]interface{}{
				appearance.Height: utils.AsAnyPtr("0"),
				appearance.Weight: utils.AsAnyPtr("0"),
				appearance.Skin:   utils.AsAnyPtr(""),
				appearance.Hair:   utils.AsAnyPtr(""),
			},
		},
		Vitals: &vitals.Vitals{
			ParseTable__: map[string]interface{}{
				vitals.HitPoints:  utils.AsAnyPtr("0"),
				vitals.ArmorClass: utils.AsAnyPtr("0"),
				vitals.Speed:      utils.AsAnyPtr("0"),
			},
		},
		Meta: &meta.Meta{
			ParseTable__: map[string]interface{}{
				meta.Initiative: utils.AsAnyPtr("0"),
				meta.Alignment:  utils.AsAnyPtr(""),
				meta.Lifestyle:  utils.AsAnyPtr(""),
			},
		},
	}
}

func New(
	name string,
	stats *stats.Stats,
	appearance *appearance.Appearance,
	meta *meta.Meta,
	vitals *vitals.Vitals,
) *Creature {
	return &Creature{
		Name:       name,
		Stats:      stats,
		Appearance: appearance,
		Meta:       meta,
		Vitals:     vitals,
	}
}

func (c *Creature) Build() *Creature {
	c.Stats = c.Stats.ParsePrimitives()
	c.Meta = c.Meta.ParsePrimitives()
	c.Vitals = c.Vitals.ParsePrimitives()
	c.Appearance = c.Appearance.ParsePrimitives()

	return c
}

func (c *Creature) ToString() string {
	var str string

	str += fmt.Sprintf("Name: %s\n", c.Name)
	str += fmt.Sprintf("Stats:\n%s", c.Stats.ToString())
	str += fmt.Sprintf("Meta:\n%s", c.Meta.ToString())
	str += fmt.Sprintf("Vitals:\n%s", c.Vitals.ToString())
	str += fmt.Sprintf("Appearance:\n%s", c.Appearance.ToString())

	return str
}
