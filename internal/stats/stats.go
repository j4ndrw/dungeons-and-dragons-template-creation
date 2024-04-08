package stats

import (
	"fmt"
	"strconv"
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
	Wisdom       int
	Intelligence int
	Strength     int
	Dexterity    int
	Constitution int
	Charisma     int

	ParseTable__ map[string]interface{}
}

func (s *Stats) ParsePrimitives() *Stats {
	s.Wisdom, _ = strconv.Atoi(*s.ParseTable__[Wisdom].(*string))
	s.Intelligence, _ = strconv.Atoi(*s.ParseTable__[Intelligence].(*string))
	s.Strength, _ = strconv.Atoi(*s.ParseTable__[Strength].(*string))
	s.Dexterity, _ = strconv.Atoi(*s.ParseTable__[Dexterity].(*string))
	s.Constitution, _ = strconv.Atoi(*s.ParseTable__[Constitution].(*string))
	s.Charisma, _ = strconv.Atoi(*s.ParseTable__[Charisma].(*string))

	return s
}

func (s *Stats) ToString() string {
	var str string

	str += fmt.Sprintf("\tWisdom: %d\n", s.Wisdom)
	str += fmt.Sprintf("\tIntelligence: %d\n", s.Intelligence)
	str += fmt.Sprintf("\tStrength: %d\n", s.Strength)
	str += fmt.Sprintf("\tDexterity: %d\n", s.Dexterity)
	str += fmt.Sprintf("\tConstitution: %d\n", s.Constitution)
	str += fmt.Sprintf("\tCharisma: %d\n", s.Charisma)

	return str
}
