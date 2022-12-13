package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

type Node struct {
	i, j int
}

const (
	Start  = 'S'
	End    = 'E'
	Bottom = 'a'
	Top    = 'z'
)

func parseInput(lines []string, startIndicator rune) (grid [][]byte, start []Node, end Node) {
	grid = make([][]byte, len(lines))

	for i, line := range lines {
		for j, h := range line {
			switch h {
			case startIndicator, Start:
				start = append(start, Node{
					i: i,
					j: j,
				})
			case End:
				end.i = i
				end.j = j

				h = Top
			}

			if h == Start {
				h = Bottom
			}

			grid[i] = append(grid[i], byte(h))
		}
	}

	return
}

func neighbors(grid [][]byte, node Node) (n []Node) {
	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		cn := Node{
			i: node.i + d[0],
			j: node.j + d[1],
		}

		// Skip out of bounds candidates.
		if cn.i < 0 || cn.i >= len(grid) || cn.j < 0 || cn.j >= len(grid[0]) {
			continue
		}

		if grid[cn.i][cn.j] <= grid[node.i][node.j]+1 {
			n = append(n, cn)
		}
	}

	return
}

func bfs(grid [][]byte, start, end Node) (int, bool) {
	var (
		queue   = []Node{start}
		cost    = map[Node]int{start: 0}
		current Node
	)

	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]

		if current == end {
			return cost[current], true
		}

		for _, next := range neighbors(grid, current) {
			if _, ok := cost[next]; !ok {
				cost[next] = cost[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return 0, false
}

func part1and2(lines []string, startIndicator rune) int {
	grid, start, end := parseInput(lines, startIndicator)
	min := -1

	for _, s := range start {
		if steps, ok := bfs(grid, s, end); ok && (min == -1 || steps < min) {
			min = steps
		}
	}

	return min
}

func main() {
	var (
		lines = common.Lines(input)
	)

	fmt.Printf("Part one: %d\n", part1and2(lines, Start))  // Part one: 449
	fmt.Printf("Part two: %d\n", part1and2(lines, Bottom)) // Part two: 443
}
