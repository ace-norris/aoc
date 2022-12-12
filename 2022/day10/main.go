package main

import (
	"bufio"
	"strings"
)

func main() {}

func exercise1(in string) int {
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
