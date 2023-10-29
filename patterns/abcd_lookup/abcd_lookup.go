package abcd_lookup

import (
	"errors"

	"github.com/consensys/gnark/frontend"
)

type LookupABCD struct {
	Input [10]frontend.Variable `gnark:",public"`
}

func (circuit *LookupABCD) Define(api frontend.API) error {
	// Define the valid characters using 2 bits
	a := []frontend.Variable{0, 0}
	b := []frontend.Variable{1, 0}
	c := []frontend.Variable{0, 1}
	d := []frontend.Variable{1, 1}

	// Ensure the input has at least the minimum length to match the pattern (2 characters: "ad")
	if len(circuit.Input) < 2 {
		return errors.New("input string too short")
	}

	// Check the first character is 'a'
	api.AssertIsEqual(circuit.Input[0], a[0])
	api.AssertIsEqual(circuit.Input[1], a[1])

	// Check the last character is 'd'
	api.AssertIsEqual(circuit.Input[len(circuit.Input)-2], d[0])
	api.AssertIsEqual(circuit.Input[len(circuit.Input)-1], d[1])

	// Ensure the middle characters are either 'b' or 'c'
	for i := 2; i < len(circuit.Input)-2; i += 2 {
		// Ensure the character is one of 'b' or 'c'
		isB := api.Lookup2(circuit.Input[i], circuit.Input[i+1], b[0], b[1], c[0], c[1])
		isC := api.Lookup2(circuit.Input[i], circuit.Input[i+1], c[0], c[1], b[0], b[1])
		api.AssertIsEqual(api.Add(isB, isC), 1)
	}

	return nil
}
