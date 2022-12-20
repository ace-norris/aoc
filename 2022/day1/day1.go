package day1

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func Exercise1(stream string) int {
	scanner := bufio.NewScanner(strings.NewReader(stream))

	var (
		elf  = 0
		elfs = map[int][]int{}
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		calorie, err := strconv.Atoi(line)
		if err != nil {
			elf += 1
			continue
		}

		calories := elfs[elf]
		calories = append(calories, calorie)
		elfs[elf] = calories
	}

	max := 0
	for _, calories := range elfs {
		total := 0
		for _, calarie := range calories {
			total += calarie
		}
		if total > max {
			max = total
		}
	}

	return max
}

func Exercise2(stream string) int {
	scanner := bufio.NewScanner(strings.NewReader(stream))

	var (
		elf  = 0
		elfs = map[int]int{}
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		calorie, err := strconv.Atoi(line)
		if err != nil {
			elf += 1
			continue
		}

		calories := elfs[elf]
		calories += calorie
		elfs[elf] = calories

	}

	totals := []int{}
	for _, calories := range elfs {
		totals = append(totals, calories)
	}

	sort.Ints(totals)
	top := totals[len(totals)-3:]
	max := 0
	for _, v := range top {
		max += v
	}

	return max
}
