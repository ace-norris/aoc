package dayx

import (
	"bufio"
	"strings"
)

func Exercise1(stream string) int {
	var (
		scanner = bufio.NewScanner(strings.NewReader(stream))
		out     = 0
	)
	for scanner.Scan() {
		// line := strings.TrimSpace(scanner.Text())
		out++
	}

	return out
}

func Exercise2(stream string) int {
	var (
		scanner = bufio.NewScanner(strings.NewReader(stream))
		out     = 0
	)
	for scanner.Scan() {
		// line := strings.TrimSpace(scanner.Text())
		out++
	}

	return out
}
