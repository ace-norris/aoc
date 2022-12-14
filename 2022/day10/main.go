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
	command command
	value   int
	cycle   int
	next    *instruction
}

func (i instruction) completed() bool {
	if i.command == noop {
		return i.cycle == 1
	}

	return i.cycle == 2
}

type instructions struct {
	head, tail *instruction
	length     int
}

func newInstructions(in string) *instructions {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = &instructions{nil, nil, 0}
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

		out.append(i)
	}

	return out
}

func (i *instructions) append(in instruction) {
	i.length++
	if i.head == nil && i.tail == nil {
		i.head = &in
		i.tail = &in
		return
	}

	i.tail.next = &in
	i.tail = i.tail.next
}

func (i *instructions) next() int {
	i.head.cycle++

	if !i.head.completed() {
		return 0
	}

	i.length--
	v := i.head
	i.head = i.head.next
	return v.value
}

func exercise1(stream string) int {
	var (
		instructions = newInstructions(stream)
		cycle        = 1
		points       = map[int]int{
			20:  0,
			60:  0,
			100: 0,
			140: 0,
			180: 0,
			220: 0,
		}
		strength = 1
		out      = 0
	)

	for instructions.length > 0 {
		v := instructions.next()
		if _, ok := points[cycle]; ok {
			points[cycle] = cycle * strength
		}

		strength += v
		cycle++
	}

	for _, v := range points {
		out += v
	}

	return out
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
