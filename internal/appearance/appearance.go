package appearance

import (
	"fmt"
	"strconv"
)

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
	a.Weight, _ = strconv.Atoi(*a.ParseTable__[Weight].(*string))
	a.Height, _ = strconv.Atoi(*a.ParseTable__[Height].(*string))
	a.Skin = *a.ParseTable__[Skin].(*string)
	a.Hair = *a.ParseTable__[Hair].(*string)

	return a
}

func (a *Appearance) ToString() string {
	var str string
	str += fmt.Sprintf("\tWeight: %d\n", a.Weight)
	str += fmt.Sprintf("\tHeight: %d\n", a.Height)
	str += fmt.Sprintf("\tSkin: %s\n", a.Skin)
	str += fmt.Sprintf("\tHair: %s\n", a.Hair)

	return str
}
