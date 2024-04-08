package stats

import "fmt"

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
	s.Wisdom = s.ParseTable__[Wisdom].(int)
	s.Intelligence = s.ParseTable__[Intelligence].(int)
	s.Strength = s.ParseTable__[Strength].(int)
	s.Dexterity = s.ParseTable__[Dexterity].(int)
	s.Constitution = s.ParseTable__[Constitution].(int)
	s.Charisma = s.ParseTable__[Charisma].(int)

	return s
}

func (s *Stats) ToString() string {
	var str string

	str += fmt.Sprintf("Wisdom: %d\n", s.Wisdom)
	str += fmt.Sprintf("Intelligence: %d\n", s.Intelligence)
	str += fmt.Sprintf("Strength: %d\n", s.Strength)
	str += fmt.Sprintf("Dexterity: %d\n", s.Dexterity)
	str += fmt.Sprintf("Constitution: %d\n", s.Constitution)
	str += fmt.Sprintf("Charisma: %d\n", s.Charisma)

	return str
}
