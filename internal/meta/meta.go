package meta

import (
	"fmt"
	"strconv"
)

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
	m.Initiative, _ = strconv.Atoi(*m.ParseTable__[Initiative].(*string))
	m.Alignment = *m.ParseTable__[Alignment].(*string)
	m.Lifestyle = *m.ParseTable__[Lifestyle].(*string)

	return m
}

func (m *Meta) ToString() string {
	var str string

	str += fmt.Sprintf("\tInitiative: %d\n", m.Initiative)
	str += fmt.Sprintf("\tAlignment: %s\n", m.Alignment)
	str += fmt.Sprintf("\tLifestyle: %s\n", m.Lifestyle)

	return str
}
