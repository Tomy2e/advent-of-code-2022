package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

func main() {
	var (
		lines    = strings.Split(input, "\n")
		elves    = make([]int, 0)
		calories = 0
	)

	for _, line := range lines {
		if line == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}

		calories += common.MustAtoi(line)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("Part one: %d\n", elves[0])                   // Part one: 75501
	fmt.Printf("Part two: %d\n", elves[0]+elves[1]+elves[2]) // Part two: 215594
}
