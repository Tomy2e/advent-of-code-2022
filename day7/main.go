package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

type Stack []string

// Push adds an element to the Stack.
func (s Stack) Push(v string) Stack {
	return append(s, v)
}

// Pop removes the top element of the Stack.
func (s Stack) Pop() Stack {
	return s[:len(s)-1]
}

// Clear returns an empty Stack.
func (s Stack) Clear() Stack {
	return make(Stack, 0)
}

// String returns the Stack as a filesystem path.
func (s Stack) String() string {
	return fmt.Sprintf("/%s", strings.Join(s, "/"))
}

// Hierarchy returns the path of all the parent directories, including
// the current path and the root path ("/").
func (s Stack) Hierarchy() (h []string) {
	s2 := make(Stack, len(s))
	copy(s2, s)

	for len(s2) > 0 {
		h = append(h, s2.String())
		s2 = s2.Pop()
	}

	// Root path.
	h = append(h, s2.String())

	return
}

// parseCommandCD returns the parameter of the cd command.
func parseCommandCD(line string) string {
	return strings.Replace(line, "$ cd ", "", 1)
}

// parseFileLine parses a line returned by the cd command. It returns 0 if
// the current line is a directory (or if the file is empty).
func parseFileLine(line string) int {
	cols := strings.Split(line, " ")
	if cols[0] == "dir" {
		return 0
	}

	return common.MustAtoi(cols[0])
}

// parseCommand parses a command (that is prefixed by "$"). The name of the
// command is the first return value. If the line is not a command, the second
// return value is false.
func parseCommand(line string) (string, bool) {
	if !strings.HasPrefix(line, "$") {
		return "", false
	}

	return strings.Split(line, " ")[1], true
}

// parseInput parses the input lines and returns a map with all the directories
// and their size.
func parseInput(lines []string) map[string]int {
	wd := make(Stack, 0)
	dirs := make(map[string]int)

	for _, line := range lines {
		if cmd, ok := parseCommand(line); ok {
			if cmd == "cd" {
				path := parseCommandCD(line)
				switch path {
				case "/":
					wd = wd.Clear()
				case "..":
					wd = wd.Pop()
				default:
					wd = wd.Push(path)
				}
			}

			continue
		}

		for _, h := range wd.Hierarchy() {
			dirs[h] += parseFileLine(line)
		}
	}

	return dirs
}

func part1(dirs map[string]int) (total int) {
	for _, s := range dirs {
		if s < 100000 {
			total += s
		}
	}

	return
}

func part2(dirs map[string]int) (min int) {
	var (
		diskSpace = 70000000
		needed    = 30000000
		free      = diskSpace - dirs["/"]
		toFree    = needed - free
	)

	for _, s := range dirs {
		if s > toFree {
			if min == 0 || s < min {
				min = s
			}
		}
	}

	return
}

func main() {
	var (
		lines = strings.Split(input, "\n")
	)

	dirs := parseInput(lines)

	fmt.Printf("Part one: %d\n", part1(dirs)) // Part one: 1297159
	fmt.Printf("Part two: %d\n", part2(dirs)) // Part two: 3866390
}
