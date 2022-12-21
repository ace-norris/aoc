package day13

import (
	"bufio"
	"encoding/json"
	"strings"
)

type result int64

const (
	equal result = iota
	left
	right
)

func newResult(lv, rv interface{}) result {
	lpv, _ := lv.(float64)
	rpv, _ := rv.(float64)

	if lpv == rpv {
		return equal
	} else if lpv < rpv {
		return left
	} else {
		return right
	}
}

func compareValues(l, r values) result {
	ml := len(l)
	if len(r) > ml {
		ml = len(r)
	}

	for j := 0; j < ml; j++ {
		var (
			lv = l.extract(j)
			rv = r.extract(j)
		)
		if lv == nil {
			return left
		}
		if rv == nil {
			return right
		}

		if !isValues(lv) && !isValues(rv) {
			if r := newResult(lv, rv); r != equal {
				return r
			}
		} else {
			if r := compareValues(newValues(lv), newValues(rv)); r != 0 {
				return r
			}
		}
	}

	return equal
}

type values []interface{}

func (i values) extract(index int) interface{} {
	if index < len(i) {
		return i[index]
	}

	return nil
}

func newValues(in interface{}) values {
	if v, ok := in.([]interface{}); ok {
		return v
	}

	if v, ok := in.(float64); ok {
		return []interface{}{v}
	}

	return nil
}

func isValues(in interface{}) bool {
	if _, ok := in.([]interface{}); ok {
		return true
	}

	return false
}

type pair struct {
	index int
	left  values
	right values
}

func (i pair) compare() result {
	return compareValues(i.left, i.right)
}

type pairs []pair

func newPairs(in string) pairs {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = pairs{}
	)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	index := 1
	for i := 0; i < len(lines); i += 3 {
		left := values{}
		json.Unmarshal([]byte(lines[i]), &left)

		right := values{}
		json.Unmarshal([]byte(lines[i+1]), &right)

		out = append(out, pair{
			index: index,
			left:  left,
			right: right,
		})
		index++
	}

	return out
}

func Exercise1(stream string) int {
	var (
		pairs = newPairs(stream)
		out   = 0
	)

	for _, v := range pairs {
		if r := v.compare(); r == left {
			out += v.index
		}
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
