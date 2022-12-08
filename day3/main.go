package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

func toPriority(r rune) int {
	ri := int(r)

	// Is uppercase?
	if ri < 97 {
		return ri - 38
	}

	return ri - 96
}

func part1(lines []string) (sum int) {
	for _, line := range lines {
		left := line[:len(line)/2]
		right := line[len(line)/2:]

		for _, r := range right {
			if strings.Contains(left, string(r)) {
				sum += toPriority(r)
				break
			}
		}
	}

	return sum
}

func part2(lines []string) (sum int) {
	for i := 0; i < len(lines); i += 3 {
		for _, r := range lines[i+2] {
			if strings.Contains(lines[i], string(r)) &&
				strings.Contains(lines[i+1], string(r)) {
				sum += toPriority(r)
				break
			}
		}
	}

	return sum
}

func main() {
	var (
		lines = common.Lines(input)
	)

	fmt.Printf("Part one: %d\n", part1(lines)) // Part one: 7967
	fmt.Printf("Part two: %d\n", part2(lines)) // Part two: 2716
}
