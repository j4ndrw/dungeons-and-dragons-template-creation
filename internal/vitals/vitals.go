package vitals

import "fmt"

const (
	HitPoints  = "HitPoints"
	ArmorClass = "ArmorClass"
	Speed      = "Speed"
)

var VitalKeys []string = []string{
	HitPoints,
	ArmorClass,
	Speed,
}

type Vitals struct {
	HitPoints  int
	ArmorClass int
	Speed      int

	ParseTable__ map[string]interface{}
}

func (v *Vitals) ParsePrimitives() *Vitals {
	v.HitPoints = v.ParseTable__[HitPoints].(int)
	v.ArmorClass = v.ParseTable__[ArmorClass].(int)
	v.Speed = v.ParseTable__[Speed].(int)

	return v
}

func (v *Vitals) ToString() string {
	var str string

	str += fmt.Sprintf("HitPoints: %d\n", v.HitPoints)
	str += fmt.Sprintf("ArmorClass: %d\n", v.ArmorClass)
	str += fmt.Sprintf("Speed: %d\n", v.Speed)

	return str
}
