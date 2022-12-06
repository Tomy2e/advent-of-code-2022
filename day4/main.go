package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

func parseAssignment(assignment string) (int, int) {
	bounds := strings.Split(assignment, "-")
	low := common.MustAtoi(bounds[0])
	high := common.MustAtoi(bounds[1])

	return low, high
}

func part1(lines []string) (total int) {
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		firstLow, firstHigh := parseAssignment(assignments[0])
		secondLow, secondHigh := parseAssignment(assignments[1])

		if firstLow >= secondLow && firstHigh <= secondHigh ||
			secondLow >= firstLow && secondHigh <= firstHigh {
			total++
		}
	}

	return
}

func part2(lines []string) (total int) {
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		firstLow, firstHigh := parseAssignment(assignments[0])
		secondLow, secondHigh := parseAssignment(assignments[1])

		if firstHigh >= secondLow && firstLow <= secondHigh {
			total++
		}
	}

	return
}

func main() {
	var (
		lines = strings.Split(input, "\n")
	)

	fmt.Printf("Part one: %d\n", part1(lines)) // Part one: 524
	fmt.Printf("Part two: %d\n", part2(lines)) // Part two: 798
}
