package day2

import (
	"bufio"
	"regexp"
	"strings"
)

func Exercise1(stream string) int {
	scanner := bufio.NewScanner(strings.NewReader(stream))

	var (
		shapes = map[string]int{
			"A": 1, // rock
			"B": 2, // paper
			"C": 3, // scissors
			"X": 1, // rock
			"Y": 2, // paper
			"Z": 3, // scissors
		}
		out = 0
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		exp := regexp.MustCompile(`^([ABC])\s([XYZ])$`).FindStringSubmatch(line)
		if len(exp) != 3 {
			continue
		}
		var (
			o = shapes[exp[1]]
			u = shapes[exp[2]]
		)
		s := u

		// Rock defeats Scissors,
		// Scissors defeats Paper,
		// Paper defeats Rock.
		if (u == 1 && o == 3) || (u == 3 && o == 2) || (u == 2 && o == 1) { // won
			s += 6
		} else if u == o { // draw
			s += 3
		}

		out += s
	}

	return out
}

func Exercise2(stream string) int {
	scanner := bufio.NewScanner(strings.NewReader(stream))

	var (
		shapes = map[string]int{
			"A": 1, // rock
			"B": 2, // paper
			"C": 3, // scissors
		}
		winningMoves = map[string]string{
			"A": "B", // rock > paper
			"B": "C", // paper > scissors
			"C": "A", // scissors > rock
		}
		loosingMoves = map[string]string{
			"B": "A", // paper > rock
			"C": "B", // scissors > paper
			"A": "C", // rock > scissors
		}
		outcomes = map[string]int{
			"X": 0, // loose
			"Y": 3, // draw
			"Z": 6, // win
		}
		out = 0
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		exp := regexp.MustCompile(`^([ABC])\s([XYZ])$`).FindStringSubmatch(line)
		if len(exp) != 3 {
			continue
		}
		var (
			u = outcomes[exp[2]]
			s = u
		)

		if u == 6 { //win
			s += shapes[winningMoves[exp[1]]]
		} else if u == 3 { // draw
			s += shapes[exp[1]]
		} else { // loose
			s += shapes[loosingMoves[exp[1]]]
		}

		out += s
	}

	return out
}
