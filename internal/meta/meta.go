package meta

import "fmt"

const (
	Initiative = "Initiative"
	Alignment  = "Alignment"
	Lifestyle  = "Lifestyle"
)

type Meta struct {
	Initiative int
	Alignment  string
	Lifestyle  string

	ParseTable__ map[string]interface{}
}

var MetaKeys []string = []string{
	Initiative,
	Alignment,
	Lifestyle,
}

func (m *Meta) ParsePrimitives() *Meta {
	m.Initiative = m.ParseTable__[Initiative].(int)
	m.Alignment = m.ParseTable__[Alignment].(string)
	m.Lifestyle = m.ParseTable__[Lifestyle].(string)

	return m
}

func (m *Meta) ToString() string {
	var str string

	str += fmt.Sprintf("Initiative: %d\n", m.Initiative)
	str += fmt.Sprintf("Alignment: %s\n", m.Alignment)
	str += fmt.Sprintf("Lifestyle: %s\n", m.Lifestyle)

	return str
}
