package abcd

import (
	"errors"

	"github.com/consensys/gnark/frontend"
)

// RegexABCD represents the circuit for the regex pattern a(b|c)*d
type RegexABCD struct {
	Input [5]frontend.Variable `gnark:"secret"` // The input string, encoded as a number
}

// define constructs the circuit for the regex pattern
func (circuit *RegexABCD) Define(api frontend.API) error {
	// Ensure the input has at least the minimum length to match the pattern (2 characters: "ad")
	if len(circuit.Input) < 2 {
		return errors.New("input string too short")
	}
	// Check the first character is 'a'
	api.AssertIsEqual(circuit.Input[0], 'a')
	// Check the last character is 'd'
	api.AssertIsEqual(circuit.Input[len(circuit.Input)-1], 'd')
	// Ensure the middle characters are either 'b' or 'c'
	for i := 1; i < len(circuit.Input)-1; i++ {
		isB := api.IsZero(api.Sub(circuit.Input[i], 'b')) // 1 if character is 'b', 0 otherwise
		isC := api.IsZero(api.Sub(circuit.Input[i], 'c')) // 1 if character is 'c', 0 otherwise

		// Ensure that only one of isB or isC is true
		onlyOneTrue := api.Xor(isB, isC)
		api.AssertIsEqual(onlyOneTrue, 1)
	}
	return nil
}
