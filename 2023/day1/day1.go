package day1

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Exercise1(stream string) int {
	values := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}

	doc := newCalibrationDocument(stream)
	sum, err := doc.Sum(values)
	if err != nil {
		panic(err)
	}

	return sum
}

func Exercise2(stream string) int {
	values := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	doc := newCalibrationDocument(stream)
	sum, err := doc.Sum(values)
	if err != nil {
		panic(err)
	}

	return sum
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type calibrationLine struct {
	Value string
}

func newCalibrationLine(in string) calibrationLine {
	return calibrationLine{in}
}

func (i calibrationLine) matches(values map[string]string) map[int]string {
	out := make(map[int]string, 0)

	for k, v := range values {
		ss := regexp.MustCompile(k).FindAllStringSubmatchIndex(i.Value, -1)

		for _, s := range ss {
			out[s[0]] = v
		}
	}

	return out
}

func (i calibrationLine) Sum(values map[string]string) (int, error) {
	m := i.matches(values)
	k := keys(m)
	sort.Ints(k)

	n := fmt.Sprintf("%s%s", m[k[0]], m[k[len(k)-1]])

	return strconv.Atoi(n)
}

type calibrationDocument []calibrationLine

func newCalibrationDocument(in string) calibrationDocument {
	out := calibrationDocument{}

	scanner := bufio.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		out = append(out, newCalibrationLine(scanner.Text()))
	}

	return out
}

func (i calibrationDocument) Sum(values map[string]string) (int, error) {
	out := 0

	for _, v := range i {
		n, err := v.Sum(values)
		if err != nil {
			return -1, err
		}
		out += n
	}

	return out, nil
}
