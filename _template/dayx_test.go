package dayx

import (
	"testing"
)

func Test_Exercise1(t *testing.T) {
	// arrange
	cs := []struct {
		Stream   string
		Expected int
	}{
		{
			Stream:   stream,
			Expected: 0,
		},
	}

	// act
	for i, c := range cs {
		out := Exercise1(c.Stream)

		// assert
		if out != c.Expected {
			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, out)
		}
	}
}

func Test_Exercise2(t *testing.T) {
	// arrange
	cs := []struct {
		Stream   string
		Expected int
	}{
		{
			Stream:   stream,
			Expected: 0,
		},
	}

	// act
	for i, c := range cs {
		out := Exercise2(c.Stream)

		// assert
		if out != c.Expected {
			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, out)
		}
	}
}

const stream = ``
