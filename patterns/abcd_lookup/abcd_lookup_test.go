package abcd_lookup

import (
	"testing"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestABCDPattern(t *testing.T) {
	assert := test.NewAssert(t)

	var abcdCircuit LookupABCD

	// Valid input: "abccd"
	validInput := [10]frontend.Variable{
		0, 0, // 'a'
		1, 0, // 'b'
		0, 1, // 'c'
		0, 1, // 'c'
		1, 1, // 'd'
	}

	assert.ProverSucceeded(&abcdCircuit, &LookupABCD{
		Input: validInput,
	})

	// Invalid input: "abcdd" (last character is not 'd')
	invalidInput := [10]frontend.Variable{
		0, 0, // 'a'
		1, 0, // 'b'
		0, 1, // 'c'
		0, 1, // 'c'
		0, 1, // 'd' (should be 1, 1 for 'd')
	}
	assert.ProverFailed(&abcdCircuit, &LookupABCD{
		Input: invalidInput,
	})
}
