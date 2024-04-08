package common

import (
	"fmt"
	"strconv"
	"strings"
)

type EnhancedValue struct {
	Value    int
	Modifier int
}

func DecodeEnhancedValue(s string) EnhancedValue {
	// string is of format <number> (+<modifier>) or <number>(-<modifier>)
	// e.g. 14 (+3), 12 (-1), 10 (+0)

	splitStr := strings.Split(s, " ")
	baseValue, _ := strconv.Atoi(splitStr[0])
	modifier := splitStr[1]
	sign := string(modifier[1])
	modifierValue, _ := strconv.Atoi(modifier[2 : len(modifier)-1])

	if sign == "-" {
		modifierValue = modifierValue * (-1)
	}

	return EnhancedValue{baseValue, modifierValue}
}

func EncodeEnhancedValue(e EnhancedValue) string {
	var sign string
	if e.Modifier >= 0 {
		sign = "+"
	} else {
		sign = "-"
	}
	var modifier int
	if e.Modifier >= 0 {
		modifier = e.Modifier
	} else {
		modifier = e.Modifier * (-1)
	}
	return fmt.Sprintf("%d (%s%d)", e.Value, sign, modifier)
}
