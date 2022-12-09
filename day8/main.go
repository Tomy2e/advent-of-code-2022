package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

// parseInput parses the lines into a list of rows.
func parseInput(lines []string) [][]int {
	trees := make([][]int, len(lines))

	for i, line := range lines {
		for _, h := range line {
			trees[i] = append(trees[i], common.MustAtoi(string(h)))
		}
	}

	return trees
}

// checkHorizontal returns true if the tree is visible from the left or the right.
func checkHorizontal(trees [][]int, i, j int) bool {
	visible := true

	for k := 0; k < len(trees); k++ {
		if k == j {
			if visible {
				break
			}

			visible = true
			continue
		}

		if trees[i][k] >= trees[i][j] {
			visible = false
		}
	}

	return visible
}

// checkVertical returns true if the tree is visible from the top or the bottom.
func checkVertical(trees [][]int, i, j int) bool {
	visible := true

	for k := 0; k < len(trees); k++ {
		if k == i {
			if visible {
				break
			}

			visible = true
			continue
		}

		if trees[k][j] >= trees[i][j] {
			visible = false
		}
	}

	return visible
}

// checkHorizontalView returns the number of trees that can be seen on the
// left and right.
func checkHorizontalView(trees [][]int, i, j int) (left, right int) {
	left = j
	right = len(trees[0]) - 1 - j

	current := &left

	for k := 0; k < len(trees); k++ {
		if k == j {
			current = &right
			continue
		}

		if trees[i][k] >= trees[i][j] {
			*current = common.Abs(j - k)

			if k > j {
				break
			}
		}
	}

	return
}

// checkVerticalView returns the number of trees that can be seen on the
// top and bottom.
func checkVerticalView(trees [][]int, i, j int) (top, bottom int) {
	top = i
	bottom = len(trees[0]) - 1 - i

	current := &top

	for k := 0; k < len(trees); k++ {
		if k == i {
			current = &bottom
			continue
		}

		if trees[k][j] >= trees[i][j] {
			*current = common.Abs(i - k)

			if k > i {
				break
			}
		}
	}

	return
}

func part1(trees [][]int) (total int) {
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees); j++ {
			if checkHorizontal(trees, i, j) || checkVertical(trees, i, j) {
				total++
				continue
			}
		}
	}

	return
}

func part2(trees [][]int) (max int) {
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees); j++ {
			left, right := checkHorizontalView(trees, i, j)
			top, bottom := checkVerticalView(trees, i, j)

			current := left * right * top * bottom

			if current > max {
				max = current
			}
		}
	}

	return
}

func main() {
	var (
		lines = common.Lines(input)
	)

	trees := parseInput(lines)

	fmt.Printf("Part one: %d\n", part1(trees)) // Part one: 1820
	fmt.Printf("Part two: %d\n", part2(trees)) // Part two: 385112
}
