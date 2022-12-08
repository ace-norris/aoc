package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

func process(in string) []Directory {
	var (
		scanner     = bufio.NewScanner(strings.NewReader(in))
		listing     = false
		currentPath = []string{}
		out         = []Directory{}
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

		out = appendFile(out, currentPath, File{
			Size: size,
			Name: name,
		})
	}

	return out
}

func flatten(in []Directory, out *[]Directory) {
	for _, d := range in {
		*out = append(*out, d)
		flatten(d.Directories, out)
	}
}

func appendFile(directories []Directory, currentPath []string, file File) []Directory {
	var i int
	for i = 0; i < len(directories); i++ {
		if directories[i].Name == currentPath[0] {
			break
		}
	}

	if i == len(directories) {
		directories = append(directories, Directory{Name: currentPath[0], Path: strings.Join(currentPath, "/")})
	}

	if len(currentPath[1:]) > 0 {
		directories[i].Directories = appendFile(directories[i].Directories, currentPath[1:], file)
	} else {
		directories[i].Files = append(directories[i].Files, file)
	}

	return directories
}

type File struct {
	Name string
	Size int
}

type Directory struct {
	Path        string
	Name        string
	Files       []File
	Directories []Directory
}

func (i Directory) CalculateSize() int {
	out := 0
	for _, file := range i.Files {
		out += file.Size
	}
	for _, d := range i.Directories {
		out += d.CalculateSize()
	}
	return out
}
