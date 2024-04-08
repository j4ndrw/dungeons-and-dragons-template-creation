package validators

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ValidatorType func(str string) error

func NumberOfPlayers(str string) error {
	number, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	if number < 2 {
		return fmt.Errorf("The dungeon must have at least two players")
	}
	return nil
}

func CheckIfContains[S ~[]E, E comparable](options S) func(str E) error {
	return func(str E) error {
		if !slices.Contains(options, str) {
			var stringShell interface{} = str
			return fmt.Errorf("%s is not a valid option. Please choose one of the following: %v", stringShell.(string), options)
		}
		return nil
	}
}

func CheckNumeric(str string) error {
	number, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	if number <= 0 {
		return fmt.Errorf("%s is not a valid option. Please choose an integer greater than zero", str)
	}
	return nil
}

func CheckNumericAllowAllNumbers(str string) error {
	_, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	return nil
}

func CheckNumericWithModifier(str string) error {
	splitStr := strings.Split(str, " ")
	if len(splitStr) < 2 {
		return fmt.Errorf("%s is not a valid value. Please enter a numeric value and a modifier in the format `<value> (+modifier)` or `<value> (-modifier)`", str)
	}

	baseValue := splitStr[0]
	modifier := splitStr[1]

	if modifier[0] != '(' && modifier[len(modifier)-1] != ')' {
		return fmt.Errorf("%s is not a valid value. Please enter a numeric value and a modifier in the format `<value> (+modifier)` or `<value> (-modifier)`", str)
	}

	sign := string(modifier[1])
	modifierValue := modifier[2 : len(modifier)-1]

	err := CheckNumeric(baseValue)
	if err != nil {
		return err
	}

	if sign != "-" && sign != "+" {
		return fmt.Errorf("%s is not a valid value. Please enter a numeric value and a modifier in the format `<value> (+modifier)` or `<value> (-modifier)`", sign)
	}

	err = CheckNumericAllowAllNumbers(modifierValue)
	if err != nil {
		return err
	}

	return nil
}
