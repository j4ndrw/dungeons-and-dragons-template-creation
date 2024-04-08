package utils

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func MapSlice[TIn, TOut any](in []TIn, mapper func(TIn) TOut) []TOut {
	out := make([]TOut, len(in))
	for i := range in {
		out[i] = mapper(in[i])
	}
	return out
}

func AsAnyPtr[T any](x T) interface{} { return &x }
