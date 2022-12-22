package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type values []interface{}

func (i values) extract(index int) interface{} {
	if index < len(i) {
		return i[index]
	}

	return nil
}

func (i values) indexOf(value string) int {
	for i, v := range i {
		if fmt.Sprintf("%v", v) == value {
			return i
		}
	}
	return -1
}

func (i values) Compare(r values) int {
	ml := len(i)
	if len(r) > ml {
		ml = len(r)
	}

	for j := 0; j < ml; j++ {
		var (
			lv = i.extract(j)
			rv = r.extract(j)
		)
		if lv == nil {
			return -1
		}
		if rv == nil {
			return 1
		}

		if !isValues(lv) && !isValues(rv) {
			lpv, _ := lv.(float64)
			rpv, _ := rv.(float64)

			if lpv < rpv {
				return -1
			} else if lpv > rpv {
				return 1
			}
		} else {
			if r := valueToValues(lv).Compare(valueToValues(rv)); r != 0 {
				return r
			}
		}
	}

	return 0
}

func (i values) Len() int {
	return len(i)
}

func (i values) Less(a, b int) bool {
	l, _ := i[a].(values)
	r, _ := i[b].(values)

	return l.Compare(r) == -1
}

func (i values) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func newValues(in string) values {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = values{}
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		v := values{}
		json.Unmarshal([]byte(line), &v)
		out = append(out, v)
	}

	return out
}

func valueToValues(in interface{}) values {
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
		if r := v.left.Compare(v.right); r == -1 {
			out += v.index
		}
	}

	return out
}

func Exercise2(stream string) int {
	values := newValues(stream)
	sort.Sort(values)

	return (values.indexOf("[[2]]") + 1) * (values.indexOf("[[6]]") + 1)
}
