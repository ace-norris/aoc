package main

import (
	"testing"
)

func Test_exercise1(t *testing.T) {
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
		out := exercise1(c.Stream)

		// assert
		if out != c.Expected {
			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, out)
		}
	}
}

func Test_exercise2(t *testing.T) {
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
		out := exercise2(c.Stream)

		// assert
		if out != c.Expected {
			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, out)
		}
	}
}

const stream = ``
