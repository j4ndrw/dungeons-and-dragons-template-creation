package meta

import (
	"fmt"

	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/common"
)

const (
	Initiative = "Initiative"
	Alignment  = "Alignment"
	Lifestyle  = "Lifestyle"
)

type Meta struct {
	Initiative common.EnhancedValue
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
	m.Initiative = common.DecodeEnhancedValue(*m.ParseTable__[Initiative].(*string))
	m.Alignment = *m.ParseTable__[Alignment].(*string)
	m.Lifestyle = *m.ParseTable__[Lifestyle].(*string)

	return m
}

func (m *Meta) ToString() string {
	var str string

	str += fmt.Sprintf("\tInitiative: %s\n", common.EncodeEnhancedValue(m.Initiative))
	str += fmt.Sprintf("\tAlignment: %s\n", m.Alignment)
	str += fmt.Sprintf("\tLifestyle: %s\n", m.Lifestyle)

	return str
}
