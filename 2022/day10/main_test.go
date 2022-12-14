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
			Stream: `noop
		addx 3
		addx -5`,
		},
		{
			Stream: `addx 15
			addx -11
			addx 6
			addx -3
			addx 5
			addx -1
			addx -8
			addx 13
			addx 4
			noop
			addx -1
			addx 5
			addx -1
			addx 5
			addx -1
			addx 5
			addx -1
			addx 5
			addx -1
			addx -35
			addx 1
			addx 24
			addx -19
			addx 1
			addx 16
			addx -11
			noop
			noop
			addx 21
			addx -15
			noop
			noop
			addx -3
			addx 9
			addx 1
			addx -3
			addx 8
			addx 1
			addx 5
			noop
			noop
			noop
			noop
			noop
			addx -36
			noop
			addx 1
			addx 7
			noop
			noop
			noop
			addx 2
			addx 6
			noop
			noop
			noop
			noop
			noop
			addx 1
			noop
			noop
			addx 7
			addx 1
			noop
			addx -13
			addx 13
			addx 7
			noop
			addx 1
			addx -33
			noop
			noop
			noop
			addx 2
			noop
			noop
			noop
			addx 8
			noop
			addx -1
			addx 2
			addx 1
			noop
			addx 17
			addx -9
			addx 1
			addx 1
			addx -3
			addx 11
			noop
			noop
			addx 1
			noop
			addx 1
			noop
			noop
			addx -13
			addx -19
			addx 1
			addx 3
			addx 26
			addx -30
			addx 12
			addx -1
			addx 3
			addx 1
			noop
			noop
			noop
			addx -9
			addx 18
			addx 1
			addx 2
			noop
			noop
			addx 9
			noop
			noop
			noop
			addx -1
			addx 2
			addx -37
			addx 1
			addx 3
			noop
			addx 15
			addx -21
			addx 22
			addx -6
			addx 1
			noop
			addx 2
			addx 1
			noop
			addx -10
			noop
			noop
			addx 20
			addx 1
			addx 2
			addx 2
			addx -6
			addx -11
			noop
			noop
			noop`,
			Expected: 13140,
		},
		{
			Stream:   stream,
			Expected: 12980,
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

/*
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
}*/

const stream = `noop
noop
addx 5
noop
noop
addx 6
addx 4
addx -4
addx 4
addx -6
addx 11
addx -1
addx 2
addx 4
addx 3
noop
addx 2
addx -30
addx 2
addx 33
noop
addx -37
noop
noop
noop
addx 3
addx 2
addx 5
addx 20
addx 7
addx -24
addx 2
noop
addx 7
addx -2
addx -6
addx 13
addx 3
addx -2
addx 2
noop
addx -5
addx 10
addx 5
addx -39
addx 1
addx 5
noop
addx 3
noop
addx -5
addx 10
addx -2
addx 2
noop
noop
addx 7
noop
noop
noop
noop
addx 3
noop
addx 3
addx 2
addx 8
addx -1
addx -20
addx 21
addx -38
addx 5
addx 2
noop
noop
noop
addx 8
noop
noop
addx -2
addx 2
addx -7
addx 14
addx 5
noop
noop
noop
addx -16
addx 17
addx 2
addx -12
addx 19
noop
noop
addx -37
noop
noop
noop
addx 3
addx 2
addx 2
addx 5
addx 20
addx -19
addx 2
noop
noop
noop
addx 5
addx 19
addx -12
addx 3
addx -2
addx 2
addx -18
addx 25
addx -14
addx -22
addx 1
noop
noop
noop
addx 3
addx 5
addx -4
addx 7
addx 4
noop
addx 1
noop
noop
addx 2
addx -6
addx 15
addx -1
addx 4
noop
noop
addx 1
addx 4
addx -33
noop
addx 21
noop`
