package abcd

import (
	"testing"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestABCDPattern(t *testing.T) {
	assert := test.NewAssert(t)

	var abcdCircuit RegexABCD

	assert.ProverFailed(&abcdCircuit, &RegexABCD{
		Input: [5]frontend.Variable{'a', 'a', 'c', 'd', 'd'},
	})

	assert.ProverSucceeded(&abcdCircuit, &RegexABCD{
		Input: [5]frontend.Variable{'a', 'b', 'c', 'c', 'd'},
	})
}
