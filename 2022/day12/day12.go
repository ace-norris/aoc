package day12

import "aoc/2022/day12/grid"

func Exercise1(stream string) int {
	var (
		r = grid.NewRows(stream)
		g = grid.NewGrid(r)
	)

	g.SetStart(*r.Find('S'))
	g.SetMove(func(a grid.Coordinate, av rune, b grid.Coordinate, bv rune) bool {
		if av == 'S' && bv == 'a' || av == 'z' && bv == 'E' {
			return true
		}

		return bv-av <= 1
	})
	g.SetEnd(func(a grid.Coordinate, av rune, b grid.Coordinate, bv rune) bool {
		return av == 'z' && bv == 'E'
	})

	return g.BFS()
}

func Exercise2(stream string) int {
	var (
		r = grid.NewRows(stream)
		g = grid.NewGrid(r)
	)

	g.SetStart(*r.Find('E'))
	g.SetMove(func(a grid.Coordinate, av rune, b grid.Coordinate, bv rune) bool {
		if av == 'E' {
			return bv == 'z'
		}

		if bv == 'S' {
			return av == 'a' || av == 'b'
		}

		return av-bv <= 1
	})
	g.SetEnd(func(a grid.Coordinate, av rune, b grid.Coordinate, bv rune) bool {
		return av == 'b' && bv == 'a' || av == 'b' && bv == 'S'
	})

	return g.BFS()
}
