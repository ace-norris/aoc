package day8

import (
	"bufio"
	"strconv"
	"strings"
)

type grid [][]int

func newGrid(in string) grid {
	scanner := bufio.NewScanner(strings.NewReader(in))
	out := grid{}
	for scanner.Scan() {
		c := []int{}
		for _, v := range strings.Split(strings.TrimSpace(scanner.Text()), "") {
			pv, _ := strconv.Atoi(v)
			c = append(c, pv)
		}
		out = append(out, c)
	}

	return out
}

func isMaximum(val, ox, oy int, grid grid) bool {
	var (
		vr = []int{}
		vc = []int{}
	)
	for x, r := range grid {
		for y, c := range r {
			if x == ox {
				vr = append(vr, c)
			}
			if y == oy {
				vc = append(vc, c)
			}
		}
	}
	if compare(val, vr[:oy]) {
		return true
	}
	if compare(val, vr[oy+1:]) {
		return true
	}
	if compare(val, vc[:ox]) {
		return true
	}
	if compare(val, vc[ox+1:]) {
		return true
	}

	return false
}

func compare(val int, items []int) bool {
	for _, v := range items {
		if v >= val {
			return false
		}
	}

	return true
}

func calculateScore(val, ox, oy int, grid grid) int {
	var (
		vr = []int{}
		vc = []int{}
	)
	for x, r := range grid {
		for y, c := range r {
			if x == ox {
				vr = append(vr, c)
			}
			if y == oy {
				vc = append(vc, c)
			}
		}
	}

	var (
		l = countVisible(val, reverse(vr[:oy]))
		r = countVisible(val, vr[oy+1:])
		t = countVisible(val, reverse(vc[:ox]))
		d = countVisible(val, vc[ox+1:])
	)

	return t * l * r * d
}

func countVisible(val int, items []int) int {
	out := 0
	for _, v := range items {
		out++
		if v >= val {
			break
		}
	}

	return out
}

func reverse(items []int) []int {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	return items
}

func Exercise1(in string) int {
	var (
		grid = newGrid(in)
		fr   = 0
		lr   = len(grid) - 1
		lc   = 0
		rc   = len(grid[0]) - 1
		out  = ((len(grid[0]) * 2) + (len(grid) * 2) - 4)
	)

	for x, r := range grid {
		for y, c := range r {
			dc := x == fr || x == lr || y == rc || y == lc
			if dc {
				continue
			}

			if !isMaximum(c, x, y, grid) {
				continue
			}

			out++
		}
	}

	return out
}

func Exercise2(in string) int {
	var (
		grid = newGrid(in)
		out  = 0
	)

	for x, r := range grid {
		for y, c := range r {
			if s := calculateScore(c, x, y, grid); s > out {
				out = s
			}
		}
	}

	return out
}
