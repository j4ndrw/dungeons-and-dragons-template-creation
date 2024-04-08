package vitals

import (
	"fmt"
	"strconv"
)

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
	v.HitPoints, _ = strconv.Atoi(*v.ParseTable__[HitPoints].(*string))
	v.ArmorClass, _ = strconv.Atoi(*v.ParseTable__[ArmorClass].(*string))
	v.Speed, _ = strconv.Atoi(*v.ParseTable__[Speed].(*string))

	return v
}

func (v *Vitals) ToString() string {
	var str string

	str += fmt.Sprintf("\tHitPoints: %d\n", v.HitPoints)
	str += fmt.Sprintf("\tArmorClass: %d\n", v.ArmorClass)
	str += fmt.Sprintf("\tSpeed: %d\n", v.Speed)

	return str
}
