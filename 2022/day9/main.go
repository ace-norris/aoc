package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

type instruction struct {
	Direction string
	Count     int
}

type instructions []instruction

func newInstructions(in string) instructions {
	scanner := bufio.NewScanner(strings.NewReader(in))

	var (
		out = instructions{}
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		exp := regexp.MustCompile(`^([UDLR])\s(\d+)$`).FindStringSubmatch(line)
		if len(exp) != 3 {
			continue
		}

		c, err := strconv.Atoi(exp[2])
		if err != nil {
			continue
		}

		out = append(out, instruction{
			Direction: exp[1],
			Count:     c,
		})
	}

	return out
}

type position struct {
	X, Y int
}

func (i position) key() string {
	return fmt.Sprintf("%v:%v", i.X, i.Y)
}

func newPosition(x, y int) position {
	return position{x, y}
}

type grid struct {
	HeadPositions []position
	TailPositions []position
}

func (i *grid) applyInstruction(instruction instruction) {
	for j := 0; j < instruction.Count; j++ {
		chp := i.HeadPositions[len(i.HeadPositions)-1]
		ctp := i.TailPositions[len(i.TailPositions)-1]

		nhp := i.nextHeadPosition(chp, instruction)
		ntp := i.nextTailPosition(chp, ctp, nhp)

		i.HeadPositions = append(i.HeadPositions, nhp)
		i.TailPositions = append(i.TailPositions, ntp)
	}
}

func (i grid) nextHeadPosition(current position, instruction instruction) position {
	switch instruction.Direction {
	case "U":
		return newPosition(current.X-1, current.Y)
	case "D":
		return newPosition(current.X+1, current.Y)
	case "L":
		return newPosition(current.X, current.Y-1)
	case "R":
		return newPosition(current.X, current.Y+1)
	default:
		return current
	}
}

func (i grid) nextTailPosition(chead, ctail, nhead position) position {
	if ctail.X == nhead.X && ctail.Y == nhead.Y {
		// . . .
		// . S .
		// . . .
		return ctail
	}

	if ctail.Y == nhead.Y && (ctail.X+1 == nhead.X || ctail.X-1 == nhead.X) {
		// . T .
		// . H .
		// . . .

		// . . .
		// . H .
		// . T .
		return ctail
	}

	if ctail.Y-1 == nhead.Y && (ctail.X+1 == nhead.X || ctail.X-1 == nhead.X) {
		// . . T
		// . H .
		// . . .

		// . . .
		// . H .
		// . . T
		return ctail
	}

	if ctail.Y+1 == nhead.Y && (ctail.X+1 == nhead.X || ctail.X-1 == nhead.X) {
		// T . .
		// . H .
		// . . .

		// . . .
		// . H .
		// T . .
		return ctail
	}

	if ctail.X == nhead.X && (ctail.Y+1 == nhead.Y || ctail.Y-1 == nhead.Y) {
		// . . .
		// . H T
		// . . .

		// . . .
		// T H .
		// . . .
		return ctail
	}

	return chead
}

func (i grid) uniquePositions(positions []position) []position {
	var (
		up  = map[string]position{}
		out = []position{}
	)

	for _, v := range positions {
		up[v.key()] = v
	}

	for _, v := range up {
		out = append(out, v)
	}

	return out
}

func (i grid) string(positions []position) string {
	var (
		minx = 0
		maxx = 0
		miny = 0
		maxy = 0
		up   = map[string]position{}
	)

	for _, p := range positions {
		up[p.key()] = p

		if p.X > maxx {
			maxx = p.X
		}
		if p.X < minx {
			minx = p.X
		}
		if p.Y > maxy {
			maxy = p.Y
		}
		if p.Y < miny {
			miny = p.Y
		}
	}

	out := ""
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			if _, ok := up[newPosition(x, y).key()]; ok {
				if x == 0 && y == 0 {
					out += "S "
				} else {
					out += "# "
				}
			} else {
				out += ". "
			}
		}
		out += "\n"
	}

	return out
}

func newGrid() grid {
	return grid{
		HeadPositions: []position{newPosition(0, 0)},
		TailPositions: []position{newPosition(0, 0)},
	}
}

func exercise1(in string) (int, string) {
	var (
		instructions = newInstructions(in)
		grid         = newGrid()
	)
	for _, instruction := range instructions {
		grid.applyInstruction(instruction)
	}

	return len(grid.uniquePositions(grid.TailPositions)), grid.string(grid.TailPositions)
}

func exercise2(in string) int {
	scanner := bufio.NewScanner(strings.NewReader(in))

	var (
		total = 0
	)
	for scanner.Scan() {
		// line := strings.TrimSpace(scanner.Text())
		total++
	}

	return total
}
