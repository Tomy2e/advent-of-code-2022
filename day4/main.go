package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseAssignment(assignment string) (int, int) {
	bounds := strings.Split(assignment, "-")
	low, err := strconv.Atoi(bounds[0])
	if err != nil {
		panic(err)
	}

	high, err := strconv.Atoi(bounds[1])
	if err != nil {
		panic(err)
	}

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

		if secondLow >= firstLow && secondLow <= firstHigh ||
			secondHigh >= firstLow && secondHigh <= firstHigh ||
			firstLow >= secondLow && firstLow <= secondHigh ||
			firstHigh >= secondLow && firstHigh <= secondHigh {
			total++
		}
	}

	return
}

func main() {
	var (
		lines = strings.Split(input, "\n")
	)

	fmt.Printf("Part one: %d\n", part1(lines))
	fmt.Printf("Part two: %d\n", part2(lines))
}
