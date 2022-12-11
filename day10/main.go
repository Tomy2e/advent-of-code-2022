package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

// Screen is a 40 pixels wide and 6 pixels high CRT screen.
type Screen [6][40]bool

// String returns the content of the screen as a string.
func (s Screen) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] {
				sb.WriteString(" # ")
			} else {
				sb.WriteString(" . ")
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func part1(lines []string) (signal int) {
	var (
		cycles    = 0
		xRegister = 1
		next      = 20
	)

	for _, line := range lines {
		instruction := strings.Split(line, " ")

		switch instruction[0] {
		case "noop":
			cycles++
		case "addx":
			cycles += 2
		}

		if cycles >= next {
			signal += next * xRegister
			next += 40
		}

		if instruction[0] == "addx" {
			xRegister += common.MustAtoi(instruction[1])
		}
	}

	return
}

func part2(lines []string) (screen Screen) {
	var (
		cycles    = 0
		xRegister = 1
	)

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		pending := 0

		switch instruction[0] {
		case "noop":
			pending = 1
		case "addx":
			pending = 2
		}

		for i := 0; i < pending; i++ {
			hpos := cycles % 40
			vpos := cycles / 40

			if hpos >= xRegister-1 && hpos <= xRegister+1 {
				screen[vpos][hpos] = true
			}

			cycles++
		}

		if instruction[0] == "addx" {
			xRegister += common.MustAtoi(instruction[1])
		}
	}

	return
}

func main() {
	var (
		lines = common.Lines(input)
	)

	fmt.Printf("Part one: %d\n", part1(lines)) // Part one: 11720
	fmt.Printf("Part two: \n%s", part2(lines)) // Part two: ERCREPCJ
}
