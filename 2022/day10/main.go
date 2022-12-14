package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

type command int64

const (
	noop command = iota
	addx
)

func newCommand(in string) command {
	switch in {
	case "noop":
		return noop
	case "addx":
		return addx
	default:
		panic(fmt.Sprintf("unsupported direction: %s", in))
	}
}

type instruction struct {
	command   command
	value     int
	execution int
	next      *instruction
}

func (i instruction) completed() bool {
	if i.command == noop {
		return i.execution == 1
	}

	return i.execution == 2
}

func (i *instruction) execute() {
	i.execution++
}

type bus struct {
	head, tail *instruction
	length     int
}

func newBus(in string) *bus {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = &bus{nil, nil, 0}
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		exp := regexp.MustCompile(`^(noop|addx)\s?(-?\d+)?`).FindStringSubmatch(line)
		if len(exp) != 3 {
			panic(fmt.Sprintf("malformed instruction: %s", line))
		}

		i := instruction{
			command: newCommand(exp[1]),
		}

		if v, err := strconv.Atoi(exp[2]); err == nil {
			i.value = v
		}

		out.queue(i)
	}

	return out
}

func (i *bus) queue(in instruction) {
	i.length++
	if i.head == nil && i.tail == nil {
		i.head = &in
		i.tail = &in
		return
	}

	i.tail.next = &in
	i.tail = i.tail.next
}

func (i *bus) next() int {
	i.head.execute()

	if !i.head.completed() {
		return 0
	}

	i.length--
	v := i.head
	i.head = i.head.next
	return v.value
}

type traces map[int]int

func newTraces(cycles ...int) traces {
	out := make(traces, len(cycles))
	for _, v := range cycles {
		out[v] = 0
	}
	return out
}

func (i traces) record(cycle, value int) {
	if _, ok := i[cycle]; ok {
		i[cycle] = cycle * value
	}
}

func (i traces) sum() int {
	out := 0
	for _, v := range i {
		out += v
	}
	return out
}

type crt struct {
	width  int
	height int
	pixels []string
	sprite int
	cursor int
}

func newCrt(width, height int) crt {
	pc := (height * width)
	px := []string{}
	for i := 0; i < pc; i++ {
		px = append(px, ".")
	}

	return crt{
		height: height,
		width:  width,
		pixels: px,
		sprite: 1,
		cursor: 0,
	}
}

func (i *crt) draw(cycle, position int) {
	if cycle%i.width == 0 {
		i.cursor += i.width
	}

	pixel := cycle - 1
	if pixel >= i.sprite-1 && pixel <= i.sprite+1 {
		i.pixels[pixel] = "#"
	}
	i.sprite = i.cursor + position
}

func (i *crt) render() string {
	out := ""
	for x := 0; x < len(i.pixels); x += i.width {
		if x != 0 {
			out += "\n"
		}
		out += strings.Join(i.pixels[x:x+i.width], "")
	}
	return out
}

func exercise1(stream string) int {
	var (
		bus      = newBus(stream)
		traces   = newTraces(20, 60, 100, 140, 180, 220)
		strength = 1
	)

	for cycle := 1; bus.length > 0; cycle++ {
		traces.record(cycle, strength)
		strength += bus.next()
	}

	return traces.sum()
}

func exercise2(stream string) string {
	var (
		bus      = newBus(stream)
		crt      = newCrt(40, 6)
		position = 1
	)

	for cycle := 1; bus.length > 0; cycle++ {
		position += bus.next()
		crt.draw(cycle, position)
	}

	return crt.render()
}
