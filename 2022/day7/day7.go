package day7

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	Name string
	Size int
}

type directory struct {
	Path        string
	Name        string
	Files       []file
	Directories []directory
}

func (i directory) CalculateSize() int {
	out := 0
	for _, file := range i.Files {
		out += file.Size
	}
	for _, d := range i.Directories {
		out += d.CalculateSize()
	}
	return out
}

func newDirectories(in string) []directory {
	var (
		scanner     = bufio.NewScanner(strings.NewReader(in))
		listing     = false
		currentPath = []string{}
		out         = []directory{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		if ok := regexp.MustCompile(`^\$ cd /$`).MatchString(line); ok {
			listing = false
			currentPath = []string{"root"}
			continue
		} else if ok := regexp.MustCompile(`^\$ ls$`).MatchString(line); ok {
			listing = true
			continue
		} else if ok := regexp.MustCompile(`^\$ cd \.\.$`).MatchString(line); ok {
			listing = false
			currentPath = currentPath[:len(currentPath)-1]
			continue
		} else if v := regexp.MustCompile(`^\$ cd (.*)$`).FindStringSubmatch(line); len(v) == 2 {
			listing = false
			currentPath = append(currentPath, v[1])
			continue
		}

		if !listing {
			continue
		}

		lineValues := regexp.MustCompile(`^(\d+) (.*)$`).FindStringSubmatch(line)
		if len(lineValues) != 3 {
			continue
		}

		var (
			name    = lineValues[2]
			size, _ = strconv.Atoi(lineValues[1])
		)

		out = appendFile(out, currentPath, file{
			Size: size,
			Name: name,
		})
	}

	return out
}

func flatten(in []directory, out *[]directory) {
	for _, d := range in {
		*out = append(*out, d)
		flatten(d.Directories, out)
	}
}

func appendFile(directories []directory, currentPath []string, file file) []directory {
	var i int
	for i = 0; i < len(directories); i++ {
		if directories[i].Name == currentPath[0] {
			break
		}
	}

	if i == len(directories) {
		directories = append(directories, directory{Name: currentPath[0], Path: strings.Join(currentPath, "/")})
	}

	if len(currentPath[1:]) > 0 {
		directories[i].Directories = appendFile(directories[i].Directories, currentPath[1:], file)
	} else {
		directories[i].Files = append(directories[i].Files, file)
	}

	return directories
}

func Exercise1(stream string) int {
	directories := newDirectories(stream)

	flattened := &[]directory{}
	flatten(directories, flattened)
	out := 0
	for _, v := range *flattened {
		s := v.CalculateSize()
		if s > 100000 {
			continue
		}
		out += s
	}

	return out
}

func Exercise2(stream string) int {
	directories := newDirectories(stream)
	flattened := &[]directory{}
	flatten(directories, flattened)

	var (
		used     = directories[0].CalculateSize()
		unnused  = 70000000 - used
		required = 30000000 - unnused
		filtered = []directory{}
	)

	for _, v := range *flattened {
		s := v.CalculateSize()
		if s < required {
			continue
		}

		filtered = append(filtered, v)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].CalculateSize() < filtered[j].CalculateSize()
	})

	return filtered[0].CalculateSize()
}
