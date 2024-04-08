package stats

import (
	"fmt"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/common"
)

const (
	Wisdom       = "Wisdom"
	Intelligence = "Intelligence"
	Strength     = "Strength"
	Dexterity    = "Dexterity"
	Constitution = "Constitution"
	Charisma     = "Charisma"
)

var ValidStats []string = []string{
	Wisdom,
	Intelligence,
	Strength,
	Dexterity,
	Constitution,
	Charisma,
}

type Stats struct {
	Wisdom       common.EnhancedValue
	Intelligence common.EnhancedValue
	Strength     common.EnhancedValue
	Dexterity    common.EnhancedValue
	Constitution common.EnhancedValue
	Charisma     common.EnhancedValue

	ParseTable__ map[string]interface{}
}

func (s *Stats) ParsePrimitives() *Stats {
	s.Wisdom = common.DecodeEnhancedValue(*s.ParseTable__[Wisdom].(*string))
	s.Intelligence = common.DecodeEnhancedValue(*s.ParseTable__[Intelligence].(*string))
	s.Strength = common.DecodeEnhancedValue(*s.ParseTable__[Strength].(*string))
	s.Dexterity = common.DecodeEnhancedValue(*s.ParseTable__[Dexterity].(*string))
	s.Constitution = common.DecodeEnhancedValue(*s.ParseTable__[Constitution].(*string))
	s.Charisma = common.DecodeEnhancedValue(*s.ParseTable__[Charisma].(*string))

	return s
}

func (s *Stats) ToString() string {
	var str string

	str += fmt.Sprintf("\tWisdom: %s\n", common.EncodeEnhancedValue(s.Wisdom))
	str += fmt.Sprintf("\tIntelligence: %s\n", common.EncodeEnhancedValue(s.Intelligence))
	str += fmt.Sprintf("\tStrength: %s\n", common.EncodeEnhancedValue(s.Strength))
	str += fmt.Sprintf("\tDexterity: %s\n", common.EncodeEnhancedValue(s.Dexterity))
	str += fmt.Sprintf("\tConstitution: %s\n", common.EncodeEnhancedValue(s.Constitution))
	str += fmt.Sprintf("\tCharisma: %s\n", common.EncodeEnhancedValue(s.Charisma))

	return str
}
