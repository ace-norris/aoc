package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_exercise1(t *testing.T) {
	// arrange
	cs := []struct {
		Stream   string
		Expected int
	}{
		{
			Stream: `R 4
						U 4
						L 3
						D 1
						R 4
						D 1
						L 5
						R 2`,
			Expected: 13,
		},
		{
			Stream:   stream,
			Expected: 6190,
		},
	}

	// act
	for i, c := range cs {
		count, vis := exercise1(c.Stream)
		if count < 100 {
			fmt.Println(vis)
		}

		// assert
		if count != c.Expected {

			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, count)
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
}
*/

func Test_nextTailPosition(t *testing.T) {
	// arrange
	cs := []struct {
		NextHeadPosition    position
		CurrentHeadPosition position
		CurrentTailPosition position
		Expected            position
	}{
		// don't move T
		{
			// . . .
			// . S .
			// . . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(1, 1),
			Expected:            newPosition(1, 1),
		},

		{
			// . T .
			// . H .
			// . . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(2, 1),
			CurrentTailPosition: newPosition(2, 1),
			Expected:            newPosition(2, 1),
		},
		{
			// . . .
			// . H .
			// . T .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(0, 1),
			CurrentTailPosition: newPosition(0, 1),
			Expected:            newPosition(0, 1),
		},
		{
			// . . T
			// . H .
			// . . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(2, 2),
			CurrentTailPosition: newPosition(2, 2),
			Expected:            newPosition(2, 2),
		},
		{
			// . . .
			// . H .
			// . . T
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(0, 2),
			CurrentTailPosition: newPosition(0, 2),
			Expected:            newPosition(0, 2),
		},
		{
			// T . .
			// . H .
			// . . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(2, 0),
			CurrentTailPosition: newPosition(2, 0),
			Expected:            newPosition(2, 0),
		},
		{
			// . . .
			// . H .
			// T . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(0, 0),
			CurrentTailPosition: newPosition(0, 0),
			Expected:            newPosition(0, 0),
		},
		{
			// . . .
			// . H T
			// . . .
			NextHeadPosition:    newPosition(1, 1),
			CurrentHeadPosition: newPosition(1, 2),
			CurrentTailPosition: newPosition(1, 2),
			Expected:            newPosition(1, 2),
		},
		{
			// . . .
			// T H .
			// . . .
			NextHeadPosition:    newPosition(1, 0),
			CurrentHeadPosition: newPosition(1, 0),
			CurrentTailPosition: newPosition(1, 0),
			Expected:            newPosition(1, 0),
		},

		// move T to previous H
		{
			// . H .
			// . # .
			// . T .
			NextHeadPosition:    newPosition(2, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(0, 1),
			Expected:            newPosition(1, 1),
		},
		{
			// . H .
			// . # .
			// T . .
			NextHeadPosition:    newPosition(2, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(0, 0),
			Expected:            newPosition(1, 1),
		},
		{
			// . . .
			// . # H
			// T . .
			NextHeadPosition:    newPosition(1, 2),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(0, 0),
			Expected:            newPosition(1, 1),
		},
		{
			// . . .
			// T # H
			// . . .
			NextHeadPosition:    newPosition(1, 2),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(1, 0),
			Expected:            newPosition(1, 1),
		},
		{
			// T . .
			// . # H
			// . . .
			NextHeadPosition:    newPosition(1, 2),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(2, 0),
			Expected:            newPosition(1, 1),
		},
		{
			// T . .
			// . # .
			// . H .
			NextHeadPosition:    newPosition(0, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(2, 0),
			Expected:            newPosition(1, 1),
		},
		{
			// . T .
			// . # .
			// . H .
			NextHeadPosition:    newPosition(0, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(2, 1),
			Expected:            newPosition(1, 1),
		},
		{
			// . . T
			// . # .
			// . H .
			NextHeadPosition:    newPosition(0, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(2, 2),
			Expected:            newPosition(1, 1),
		},
		{
			// . . T
			// H # .
			// . . .
			NextHeadPosition:    newPosition(1, 0),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(2, 2),
			Expected:            newPosition(1, 1),
		},
		{
			// . . .
			// H # T
			// . . .
			NextHeadPosition:    newPosition(1, 0),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(1, 2),
			Expected:            newPosition(1, 1),
		},
		{
			// . . .
			// H # .
			// . . T
			NextHeadPosition:    newPosition(1, 0),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(0, 2),
			Expected:            newPosition(1, 1),
		},
		{
			// . H .
			// . # .
			// . . T
			NextHeadPosition:    newPosition(2, 1),
			CurrentHeadPosition: newPosition(1, 1),
			CurrentTailPosition: newPosition(0, 2),
			Expected:            newPosition(1, 1),
		},
	}

	// act
	for i, c := range cs {
		g := newGrid()
		out := g.nextTailPosition(c.CurrentHeadPosition, c.CurrentTailPosition, c.NextHeadPosition)

		// assert
		if !reflect.DeepEqual(out, c.Expected) {
			t.Errorf("Case: %v, Expected: %v, Actual: %v", i, c.Expected, out)
		}
	}
}

const stream = `R 2
D 2
U 1
D 2
U 2
R 2
U 2
R 2
L 2
D 2
R 1
D 2
R 1
L 1
R 2
U 1
R 1
L 1
U 1
L 2
R 2
L 2
R 2
L 2
R 2
D 2
U 2
L 1
D 1
R 2
L 1
R 2
D 1
R 1
D 2
R 2
D 1
R 1
D 1
L 2
R 1
U 2
D 1
R 1
L 2
D 2
U 2
L 2
U 2
L 2
R 1
D 1
L 1
R 2
L 2
R 1
L 2
D 2
L 1
U 2
D 1
L 1
D 2
U 1
L 1
R 2
D 2
L 2
R 2
U 2
L 2
R 2
L 2
R 2
L 2
R 1
L 2
U 1
D 1
R 2
L 1
R 1
U 2
R 1
U 1
D 1
U 1
L 1
R 2
U 1
L 1
R 2
L 1
R 2
U 1
L 1
D 1
U 2
R 2
L 1
D 2
R 2
D 2
R 1
L 2
D 1
U 2
D 1
L 2
U 1
R 1
L 1
U 2
R 3
U 1
R 1
U 3
R 1
D 2
R 2
U 3
R 1
U 2
L 2
D 2
U 3
L 3
D 2
U 1
L 3
D 2
R 2
U 1
L 1
D 3
L 2
U 2
L 3
U 1
R 1
U 3
R 3
U 1
D 1
U 2
L 2
D 1
U 2
R 2
U 3
R 2
L 2
R 1
D 1
U 3
R 1
U 3
R 3
U 3
L 1
D 2
R 3
L 2
R 2
D 1
L 2
D 1
R 1
D 2
U 3
D 2
L 2
D 2
L 1
D 2
U 3
D 3
R 3
L 3
D 2
U 2
D 1
R 3
D 2
L 3
D 3
L 1
D 1
U 3
L 3
R 3
U 3
R 2
L 2
D 2
R 3
U 1
D 3
R 2
U 2
R 1
U 3
L 1
U 2
D 2
U 3
L 1
U 3
R 1
U 2
R 2
L 3
U 2
D 3
L 2
R 3
D 2
U 3
D 1
U 2
L 3
U 1
R 4
D 1
R 4
D 2
R 3
L 3
U 4
D 1
L 4
U 2
D 1
U 1
L 3
D 2
U 4
D 2
U 3
R 2
L 4
U 1
D 3
U 2
L 3
D 1
R 2
D 1
R 1
D 3
U 4
D 3
L 3
R 1
U 1
D 2
L 2
R 1
D 2
L 3
R 2
D 3
R 3
U 1
D 2
R 4
D 3
L 3
D 1
R 2
D 1
L 1
D 1
L 2
D 1
U 1
R 1
D 3
L 1
U 1
L 2
R 2
U 1
D 4
U 1
L 2
R 3
D 1
U 4
R 4
L 2
D 2
U 4
R 2
U 3
R 4
U 1
D 1
U 2
D 2
U 3
L 3
U 4
R 3
D 2
U 2
L 1
R 2
U 4
D 4
U 1
D 2
L 4
U 4
R 4
U 3
R 3
L 1
R 1
U 3
L 2
U 3
L 2
R 3
D 2
R 1
L 1
U 3
D 1
U 3
L 1
D 1
R 1
L 4
R 1
L 1
R 5
D 5
R 5
D 2
R 3
U 2
D 5
L 4
U 2
L 2
U 1
D 3
L 1
D 2
L 2
D 3
L 1
U 2
L 1
R 4
U 5
L 2
U 5
R 4
L 3
R 4
U 5
L 3
D 2
R 2
L 4
D 2
L 1
U 2
L 1
R 5
D 5
U 3
R 4
U 4
L 3
U 3
D 1
R 2
L 5
D 2
R 4
L 2
R 1
D 4
R 2
L 2
D 3
L 4
D 1
L 1
D 1
U 5
R 5
U 5
L 1
U 5
D 5
R 5
L 2
U 5
L 2
D 1
U 4
R 1
U 2
R 2
U 2
R 2
L 4
R 3
U 3
R 5
L 5
R 4
L 4
U 4
D 2
U 3
L 1
R 2
U 3
D 1
L 3
R 4
D 5
L 3
R 2
U 2
L 1
R 1
D 2
L 5
U 2
L 3
U 2
D 5
R 2
L 1
R 3
U 1
R 5
L 5
U 6
L 6
R 5
L 2
R 2
D 4
U 3
L 5
R 4
D 6
R 2
L 2
U 1
D 6
R 2
D 4
U 3
D 4
L 5
U 1
D 6
R 6
U 2
R 3
L 6
D 3
U 2
L 3
D 6
R 5
D 4
L 1
U 3
L 5
U 3
R 2
D 2
R 3
U 3
D 5
L 3
R 6
U 2
R 6
L 3
U 1
D 3
U 1
R 3
L 1
D 3
R 2
L 3
U 5
R 6
U 3
D 2
U 5
R 6
D 5
L 6
R 5
D 2
R 3
U 6
L 6
D 3
R 6
L 4
R 1
U 4
D 5
L 4
U 2
D 1
R 4
U 3
L 1
D 6
L 6
R 1
D 2
R 2
L 2
U 4
D 2
L 2
U 5
R 2
L 4
U 2
L 5
R 1
L 2
D 4
R 1
L 2
U 6
R 1
U 6
R 4
D 3
R 1
L 1
R 1
L 4
U 2
L 6
R 6
U 4
L 2
U 2
L 7
R 4
L 6
R 4
L 7
U 4
D 3
L 3
R 7
L 2
R 6
L 2
R 6
D 6
U 3
D 7
R 3
U 6
R 3
D 6
U 4
D 3
L 5
U 5
L 3
R 3
D 7
L 3
U 3
R 7
L 1
U 2
L 6
R 7
D 3
L 1
D 7
R 4
U 5
L 3
R 3
U 6
R 6
U 1
D 5
U 4
L 1
U 6
L 2
D 1
R 2
D 3
R 1
U 3
D 6
U 7
D 6
U 3
R 7
U 2
R 2
U 1
R 4
L 6
U 4
D 5
R 1
L 1
D 7
L 4
R 2
U 6
R 4
D 1
R 5
L 5
R 3
D 1
L 4
R 5
L 4
U 1
D 6
L 1
D 3
L 4
R 2
U 5
L 6
R 4
U 6
L 4
U 5
D 1
L 4
D 3
R 1
U 3
R 2
L 2
D 6
R 1
L 5
D 6
U 3
D 2
L 6
D 1
U 7
D 3
U 2
D 8
R 2
D 3
L 4
R 3
U 7
R 8
D 3
R 2
D 8
L 4
R 4
L 7
R 1
U 8
R 7
U 2
D 5
L 6
D 8
U 3
L 5
R 4
D 1
R 4
L 6
U 1
R 5
L 8
D 6
L 3
D 6
L 4
U 2
R 3
D 4
R 6
D 6
U 6
D 5
R 5
L 5
U 6
D 5
L 6
R 8
D 6
R 6
U 5
D 6
U 3
D 7
L 8
R 2
D 8
L 5
R 7
L 5
U 5
L 2
D 3
R 3
D 2
R 2
D 8
R 8
D 2
L 2
U 1
D 4
U 6
R 2
L 2
R 5
L 3
R 2
D 6
R 2
U 5
L 7
D 7
L 7
R 1
D 4
L 3
D 1
R 3
L 8
D 8
R 4
U 7
R 2
L 6
D 5
U 4
D 3
L 2
U 5
D 5
U 1
D 7
L 7
U 7
D 4
U 3
L 7
U 6
L 6
R 5
U 5
D 2
L 3
U 6
L 6
U 2
R 3
U 5
D 1
U 8
D 1
L 6
D 5
L 8
R 6
L 9
U 3
D 2
R 9
U 1
L 3
R 8
U 6
R 8
L 6
R 2
L 3
R 7
D 6
R 3
U 2
R 1
U 3
L 4
R 6
L 9
D 4
L 7
U 2
D 4
L 8
U 7
R 7
U 5
D 4
U 1
L 8
R 4
D 5
R 3
D 5
U 9
D 6
R 3
D 6
L 7
U 7
L 1
R 3
U 2
L 4
R 7
L 9
D 3
L 6
D 4
R 6
D 2
L 6
D 3
L 3
R 2
L 7
D 4
R 9
D 7
U 1
L 3
R 9
L 7
R 2
L 2
D 5
L 9
D 6
R 5
D 8
U 7
D 5
R 1
D 4
R 5
D 3
R 4
D 5
U 7
D 5
U 3
D 3
L 8
U 1
R 5
U 3
L 7
D 5
R 4
D 6
L 5
D 7
R 5
D 7
U 7
R 8
U 6
D 1
U 4
R 5
U 4
D 5
U 2
D 5
L 4
D 1
R 7
L 10
D 4
R 10
U 8
D 5
L 4
R 3
D 10
U 1
R 8
L 4
R 7
L 5
U 8
L 3
U 6
R 5
D 4
L 7
D 9
R 10
U 9
L 7
R 1
D 9
R 3
U 1
L 7
D 3
U 7
L 4
D 2
R 5
L 1
U 3
L 1
D 1
R 9
U 8
L 4
D 7
R 5
U 9
L 10
U 9
R 4
D 6
U 9
R 8
D 3
R 8
D 4
U 2
R 7
L 9
D 3
U 8
L 8
R 1
L 9
D 9
U 9
L 9
R 8
L 5
D 6
L 5
U 10
L 1
U 2
R 7
D 7
U 5
L 1
D 5
U 8
D 10
U 10
R 2
L 4
R 5
U 9
L 4
U 10
D 5
U 7
D 1
U 1
R 7
U 2
L 3
U 3
L 9
U 5
L 10
U 7
L 1
D 1
L 2
U 7
L 4
U 6
D 9
U 6
R 4
L 6
R 4
L 8
R 4
D 4
L 11
U 7
R 6
D 2
L 1
D 3
L 2
R 8
D 2
L 7
U 9
L 9
D 1
U 9
D 2
R 1
L 6
R 10
L 2
U 2
L 9
R 8
L 3
R 6
L 9
U 1
R 3
D 7
U 10
D 1
L 1
D 5
L 1
U 2
L 6
D 11
R 11
L 3
D 6
R 6
D 8
L 6
R 5
U 9
D 7
R 2
U 6
R 4
U 11
D 2
R 4
L 11
D 1
R 6
D 8
L 11
U 7
D 6
R 6
L 4
R 5
L 4
R 3
D 2
R 1
U 6
D 9
R 11
U 4
L 10
R 5
L 2
U 7
R 4
L 2
R 9
D 1
R 7
L 11
R 3
L 4
D 7
L 10
D 3
U 11
L 6
U 9
R 8
U 3
D 10
U 8
D 5
R 1
L 5
R 11
L 11
R 9
U 7
R 1
L 9
D 3
L 9
U 1
L 5
D 10
L 8
R 4
D 10
R 4
D 5
L 12
D 2
U 1
D 1
U 9
R 2
D 2
L 2
D 6
R 10
U 3
R 8
L 3
R 11
L 5
U 11
D 5
U 4
L 3
D 11
R 5
D 4
L 11
D 2
U 9
R 10
L 10
R 7
D 11
R 1
U 9
L 4
U 3
L 7
U 1
L 10
R 7
L 11
U 11
R 9
L 4
R 2
L 10
R 1
U 3
R 11
D 5
U 1
R 12
D 9
R 4
L 9
D 8
R 7
L 2
D 7
R 8
D 5
U 7
L 12
R 6
U 3
R 11
L 9
R 8
U 8
R 4
L 10
U 4
R 7
D 10
U 4
L 11
D 1
U 6
L 4
D 2
L 2
D 3
L 7
U 6
R 3
U 9
L 8
R 8
U 10
R 8
U 9
R 4
U 6
R 12
U 8
L 10
D 8
R 10
L 1
U 5
D 7
L 10
R 9
U 2
D 2
R 12
L 1
D 12
U 8
R 4
L 12
R 7
D 4
U 10
L 11
U 8
L 4
R 1
L 10
U 12
R 5
U 2
R 5
L 11
U 6
D 8
L 11
D 1
U 13
D 8
R 4
D 6
R 11
L 10
R 6
L 4
D 7
U 10
R 1
L 8
U 7
D 1
R 5
D 5
L 1
U 7
L 2
U 4
D 13
R 10
U 11
D 13
L 8
R 2
L 2
D 2
L 9
R 10
U 10
D 3
U 13
R 4
D 9
L 9
U 10
D 7
U 4
R 13
U 6
D 1
U 8
L 1
R 13
U 3
L 11
U 12
R 9
L 8
R 9
L 2
R 5
L 7
D 11
L 8
R 13
D 10
U 4
R 5
D 13
L 2
R 9
L 5
D 8
L 13
U 13
D 1
L 4
U 5
D 13
R 2
L 6
D 12
L 4
U 5
R 10
U 2
L 1
R 8
U 10
L 9
U 3
D 3
U 12
D 6
U 13
L 5
D 6
U 14
L 13
U 7
D 2
U 8
L 13
R 10
U 12
R 5
D 10
U 5
R 8
D 2
L 4
R 9
U 11
R 10
U 7
L 14
R 14
U 14
D 13
U 9
D 13
R 5
D 14
U 11
R 7
U 11
D 8
R 11
U 2
R 4
U 6
L 11
U 6
R 8
L 6
D 1
L 1
D 5
R 2
U 11
D 9
R 14
L 13
U 4
R 8
L 4
R 8
D 14
U 3
L 3
D 4
L 4
D 11
U 4
L 5
U 3
L 1
R 2
D 8
L 11
R 7
U 13
L 9
R 14
D 1
R 8
U 10
R 9
L 5
U 7
L 2
R 14
D 6
U 12
D 10
U 1
L 14
D 2
R 14
U 6
L 7
U 3
L 6
U 7
L 6
D 7
R 9
U 11
D 9
U 7
R 14
U 8
L 4
U 8
R 1
L 12
U 1
R 2
D 11
R 8
D 9
U 9
D 3
U 5
D 8
U 8
D 1
L 12
U 15
L 2
D 13
R 4
U 2
R 7
U 1
R 4
L 4
D 1
R 5
D 9
L 12
D 11
R 4
D 14
L 1
U 13
R 6
D 13
U 6
D 10
L 3
U 8
L 12
D 1
L 6
D 5
R 12
D 11
L 10
U 14
R 2
D 4
U 15
D 10
R 12
D 7
U 14
L 14
R 7
U 15
D 5
U 13
R 5
U 6
D 10
L 12
D 1
U 2
R 14
D 9
U 7
L 5
D 9
R 14
L 12
R 2
L 6
R 13
U 11
L 10
U 2
R 2
L 5
R 8
L 7
U 7
L 15
R 5
U 11
L 10
R 11
L 5
U 6
R 4
D 14
U 13
R 8
L 10
R 5
U 8
L 14
D 6
L 5
D 7
R 11
L 15
R 6
L 15
D 7
R 9
D 12
L 3
R 11
L 10
D 6
L 6
U 9
L 10
R 11
D 7
L 13
D 11
R 1
U 4
L 2
R 4
D 6
U 8
R 15
L 10
U 16
R 10
D 1
U 14
L 6
D 1
R 8
U 6
D 11
L 9
R 13
U 12
R 8
D 7
R 15
U 1
D 3
U 13
L 11
D 6
U 2
D 1
R 14
U 10
R 8
L 9
D 12
U 16
L 8
U 10
D 8
U 2
L 4
D 4
L 3
R 3
D 5
U 2
L 11
D 16
L 15
R 15
L 12
D 8
R 1
D 11
U 5
R 10
U 12
D 9
L 16
U 10
L 13
R 10
L 4
R 6
D 15
U 12
L 2
R 15
U 5
R 1
L 8
D 6
L 7
D 13
L 7
D 10
R 11
L 14
U 12
D 7
R 1
L 1
R 3
U 1
L 1
R 14
U 8
D 7
L 4
D 15
R 8
U 1
L 4
R 12
U 6
L 13
R 14
U 3
L 4
D 9
R 6
U 5
R 14
L 6
D 5
U 6
R 14
L 10
U 13
L 16
R 15
L 12
U 10
L 2
D 8
L 7
D 16
L 17
U 16
R 13
D 6
U 8
L 13
D 9
U 13
D 3
R 1
U 12
D 6
L 8
U 15
R 6
L 14
U 15
D 4
U 9
R 6
D 8
L 10
D 3
R 12
L 10
R 2
D 2
U 17
D 16
R 5
L 13
R 12
L 4
U 4
L 12
U 3
L 15
U 10
L 5
U 2
D 13
R 10
U 7
R 9
D 7
R 12
U 14
L 7
R 10
D 11
U 8
L 5
R 4
U 17
L 12
R 6
L 17
D 11
R 17
U 1
D 2
U 14
D 17
L 2
D 14
R 14
L 1
U 5
D 8
U 1
R 12
U 8
D 3
L 1
D 8
R 5
U 9
D 13
U 15
D 3
R 15
L 2
R 2
U 8
R 1
L 17
U 16
L 10
D 2
R 6
L 11
D 14
R 3
L 15
U 16
R 12
U 1
D 2
U 2
R 7
L 4
D 12
U 12
L 12
U 15
R 4
U 8
R 12
L 5
R 3
L 5
R 18
L 4
D 14
R 4
L 17
R 18
U 9
D 18
R 4
U 13
R 3
D 6
R 15
L 1
U 5
L 3
U 2
R 15
U 8
L 7
D 9
L 15
U 13
L 17
D 17
R 15
D 15
L 14
D 15
R 16
U 3
D 6
L 1
R 14
U 16
L 7
U 2
L 14
D 5
L 6
D 12
R 18
U 16
D 7
L 18
D 10
R 1
L 8
R 16
L 5
U 4
R 18
D 18
R 5
D 10
L 18
R 8
L 2
R 13
L 16
D 16
U 17
L 2
R 5
L 12
D 10
L 10
D 6
R 12
L 17
R 8
L 14
R 15
L 18
D 5
R 1
L 8
U 13
R 9
U 14
L 18
U 4
D 16
L 7
U 16
L 14
D 18
R 1
D 2
R 11
D 3
U 12
D 11
U 2
L 12
U 7
D 2
R 8
D 1
U 9
L 5
D 14
L 9
D 5
L 18
D 16
U 3
R 13
D 16
R 18
D 10
L 4
R 2
L 5
D 13
L 15
U 16
R 9
D 15
R 5
L 3
U 5
R 10
U 3
L 11
R 14
L 4
D 7
R 11
D 12
L 14
R 13
U 15
R 14
D 2
R 16
U 3
D 3
L 17
D 11
U 10
L 2
U 12
D 3
L 5
D 9
U 7
D 14
R 5
L 14
D 3
R 11
L 17
U 7
L 8
U 2
R 3
D 18
L 16
U 12
D 16
R 6
U 16
D 10
R 1
L 7
U 18
R 2
U 13
R 15
D 1
R 15
D 9
U 9
R 15
D 6
L 14
D 15
R 15
D 10
U 15
L 17
D 12
R 6
L 1
R 15
L 7
D 11
R 12
D 9
L 17
D 18
L 19
D 2
L 9
D 1
L 7
R 5
L 5
R 1
D 2
L 1
U 19
L 5
U 2
D 15
L 10
R 7
L 18
D 13
R 2
D 18
U 8
D 15
U 14
R 5
L 9
U 9
D 18
L 9
U 6
R 7
L 10`
