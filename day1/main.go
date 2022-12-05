package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
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

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		calories += cal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("Part one: %d\n", elves[0])
	fmt.Printf("Part two: %d\n", elves[0]+elves[1]+elves[2])
}
