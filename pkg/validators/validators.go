package validators

import (
	"fmt"
	"slices"
	"strconv"
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
	if number < 0 {
		return fmt.Errorf("%s is not a valid option. Please choose an integer greater than zero", str)
	}
	return nil
}
