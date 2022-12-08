package main

import (
	"bufio"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {}

func exercise1(in string) int {
	low, high := buildPriorities()
	scanner := bufio.NewScanner(strings.NewReader(in))
	total := 0
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		middle := len(line) / 2
		left := line[:middle]
		right := line[middle:]
		dups := map[string]struct{}{}
		for _, l := range left {
			for _, r := range right {
				if l == r {
					dups[l] = struct{}{}
				}
			}
		}
		for k := range dups {
			if p, ok := low[k]; ok {
				total += p
				continue
			}
			if p, ok := high[k]; ok {
				total += p
				continue
			}
		}
	}

	return total
}

func exercise2(in string) int {
	low, high := buildPriorities()
	scanner := bufio.NewScanner(strings.NewReader(in))
	groups := map[int][]string{}
	index := 0
	group := 0
	for scanner.Scan() {
		if index%3 == 0 {
			group++
		}
		lines := groups[group]
		lines = append(lines, strings.TrimSpace(scanner.Text()))
		groups[group] = lines
		index++
	}

	badges := []string{}
	for _, group := range groups {
		for _, item := range strings.Split(group[0], "") {

			x := slices.IndexFunc(strings.Split(group[1], ""), func(i string) bool { return i == item })
			y := slices.IndexFunc(strings.Split(group[2], ""), func(i string) bool { return i == item })

			if x > -1 && y > -1 {
				badges = append(badges, item)
				break
			}
		}
	}

	total := 0
	for _, badge := range badges {
		if p, ok := low[badge]; ok {
			total += p
			continue
		}
		if p, ok := high[badge]; ok {
			total += p
			continue
		}
	}

	return total
}

func buildPriorities() (map[string]int, map[string]int) {
	var (
		characters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		low        = map[string]int{}
		high       = map[string]int{}
	)

	for i, v := range characters {
		low[strings.ToLower(v)] = i + 1
	}

	for i, v := range characters {
		high[strings.ToUpper(v)] = (i + 27)
	}

	return low, high
}
