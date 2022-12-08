package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

type elf struct {
	SectionStart int
	SectionEnd   int
}

type pair struct {
	First  elf
	Second elf
}

func (i *pair) fullOverlap() bool {
	if i.First.SectionStart >= i.Second.SectionStart && i.First.SectionEnd <= i.Second.SectionEnd {
		return true
	}

	if i.Second.SectionStart >= i.First.SectionStart && i.Second.SectionEnd <= i.First.SectionEnd {
		return true
	}

	return false
}
func (i *pair) anyOverlap() bool {
	if i.First.SectionStart >= i.Second.SectionStart && i.First.SectionStart <= i.Second.SectionEnd {
		return true
	}

	if i.Second.SectionStart >= i.First.SectionStart && i.Second.SectionStart <= i.First.SectionEnd {
		return true
	}

	return false
}

func newPair(in string) *pair {
	exp := regexp.MustCompile(`^(\d+)\-(\d+),(\d+)\-(\d+)$`).FindStringSubmatch(in)
	if len(exp) != 5 {
		return nil
	}

	xs, err := strconv.Atoi(exp[1])
	if err != nil {
		return nil
	}
	xe, err := strconv.Atoi(exp[2])
	if err != nil {
		return nil
	}
	ys, err := strconv.Atoi(exp[3])
	if err != nil {
		return nil
	}
	ye, err := strconv.Atoi(exp[4])
	if err != nil {
		return nil
	}

	return &pair{
		First: elf{
			SectionStart: xs,
			SectionEnd:   xe,
		},
		Second: elf{
			SectionStart: ys,
			SectionEnd:   ye,
		},
	}
}

func process(in string, allOverlapping bool) int {
	scanner := bufio.NewScanner(strings.NewReader((in)))
	total := 0
	for scanner.Scan() {
		p := newPair(strings.TrimSpace(scanner.Text()))
		if allOverlapping && p.fullOverlap() {
			total++
		}

		if !allOverlapping && p.anyOverlap() {
			total++
		}

	}
	return total
}
