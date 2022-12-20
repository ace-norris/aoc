package day5

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	Count int
	From  int
	To    int
}

func newInstruction(in string) *instruction {
	exp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`).FindStringSubmatch(in)
	if len(exp) != 4 {
		return nil
	}

	count, err := strconv.Atoi(exp[1])
	if err != nil {
		return nil
	}

	from, err := strconv.Atoi(exp[2])
	if err != nil {
		return nil
	}

	to, err := strconv.Atoi(exp[3])
	if err != nil {
		return nil
	}

	return &instruction{
		Count: count,
		From:  from,
		To:    to,
	}
}

func loadInstructions(in string) []instruction {
	scanner := bufio.NewScanner(strings.NewReader(in))
	out := []instruction{}
	for scanner.Scan() {
		if i := newInstruction(scanner.Text()); i != nil {
			out = append(out, *i)
		}
	}
	return out
}

type stack []string

func process(stacks []stack, instructions []instruction, preserveOrder bool) string {
	s := append([]stack(nil), stacks...)
	for _, instruction := range instructions {
		if preserveOrder {
			f := s[instruction.From-1]
			s[instruction.To-1] = append(s[instruction.To-1], f[len(f)-(instruction.Count):]...)
			s[instruction.From-1] = f[:len(f)-(instruction.Count)]
		} else {
			for i := 0; i < instruction.Count; i++ {
				f := s[instruction.From-1]
				s[instruction.To-1] = append(s[instruction.To-1], f[len(f)-1])
				s[instruction.From-1] = f[:len(f)-1]
			}
		}
	}
	out := ""
	for _, stack := range s {
		out += stack[len(stack)-1]
	}
	return out
}

func Exercise1(stream string, stacks []stack) string {
	return process(stacks, loadInstructions(stream), false)
}

func Exercise2(stream string, stacks []stack) string {
	return process(stacks, loadInstructions(stream), true)
}
