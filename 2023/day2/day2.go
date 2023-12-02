package day2

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Exercise1(stream string, possibilities map[string]int) int {
	games := newGames(stream)

	return games.SumPossible(possibilities)
}

func Exercise2(stream string) int {
	games := newGames(stream)

	return games.SumPower()
}

type games struct {
	Games []game
}

func newGames(in string) games {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = games{}
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		out.Games = append(out.Games, newGame(line))
	}

	return out
}

func (i games) SumPossible(in map[string]int) int {
	out := 0

	for _, g := range i.Games {

		if g.Possible(in) {
			out += g.ID
		}
	}
	return out
}

func (i games) SumPower() int {
	out := 0

	for _, g := range i.Games {
		s := map[string][]int{}
		for _, h := range g.Hands {
			for c, n := range h.colours {
				s[c] = append(s[c], n)
			}
		}

		t := []int{}
		for _, ns := range s {
			sort.Ints(ns)
			t = append(t, ns[len(ns)-1])
		}

		c := 0
		for _, n := range t {
			if c == 0 {
				c = n
				continue
			}
			c = c * n
		}

		out += c
	}
	return out
}

type game struct {
	ID    int
	Hands []hand
}

func (i game) Possible(in map[string]int) bool {
	for _, h := range i.Hands {
		if !h.Possible(in) {
			return false
		}
	}
	return true
}

func newGame(in string) game {
	parts := strings.Split(in, ":")

	exp := regexp.MustCompile(`^Game\s(.*)$`).FindStringSubmatch(parts[0])
	out := game{}
	id, _ := strconv.Atoi(exp[1])
	out.ID = id
	for _, h := range strings.Split(parts[1], ";") {
		out.Hands = append(out.Hands, newHand(h))
	}

	return out
}

type hand struct {
	colours map[string]int
}

func (i hand) Possible(in map[string]int) bool {
	for c, n := range i.colours {

		pn, ok := in[c]
		if !ok {
			return false
		}

		if n > pn {
			return false
		}
	}

	return true
}

func newHand(in string) hand {
	out := hand{
		colours: map[string]int{},
	}

	for _, g := range strings.Split(in, ",") {
		exp := regexp.MustCompile(`^\s(.*)\s(.*)$`).FindStringSubmatch(g)

		c := exp[2]
		n, _ := strconv.Atoi(exp[1])

		out.colours[c] = n

	}

	return out
}
