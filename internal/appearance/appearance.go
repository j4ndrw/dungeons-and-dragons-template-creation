package appearance

import "fmt"

const (
	Weight = "Weight"
	Height = "Height"
	Skin   = "Skin"
	Hair   = "Hair"
)

var Appearances []string = []string{
	Weight, Height, Skin, Hair,
}

type Appearance struct {
	Weight int
	Height int
	Skin   string
	Hair   string

	ParseTable__ map[string]interface{}
}

func (a *Appearance) ParsePrimitives() *Appearance {
	a.Weight = a.ParseTable__[Weight].(int)
	a.Height = a.ParseTable__[Height].(int)
	a.Skin = a.ParseTable__[Skin].(string)
	a.Hair = a.ParseTable__[Hair].(string)

	return a
}

func (a *Appearance) ToString() string {
	var str string
	str += fmt.Sprintf("Weight: %d\n", a.Weight)
	str += fmt.Sprintf("Height: %d\n", a.Height)
	str += fmt.Sprintf("Skin: %s\n", a.Skin)
	str += fmt.Sprintf("Hair: %s\n", a.Hair)

	return str
}
