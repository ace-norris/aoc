package main

import (
	"strings"
)

func find(stream string, count int) int {
	var (
		parts = strings.Split(stream, "")
		index = 0
	)
	for index < len(parts) {
		end := index + count
		if c := uniq(parts[index:end]); len(c) == count {
			return end
		}
		index++

	}
	return -1
}

func uniq(in []string) []string {
	uniq := map[string]struct{}{}
	for _, v := range in {
		uniq[v] = struct{}{}
	}

	out := []string{}
	for k := range uniq {
		out = append(out, k)
	}
	return out
}
