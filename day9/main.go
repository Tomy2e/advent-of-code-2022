package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

// Direction describes the movement of a point.
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
	None
)

// NewDirection parses a single letter direction into a Direction.
func NewDirection(d byte) Direction {
	switch d {
	case 'L':
		return Left
	case 'R':
		return Right
	case 'U':
		return Up
	case 'D':
		return Down
	}

	panic("unknown direction")
}

// Point is a 2D point.
type Point struct {
	x, y int
}

// Move moves the point according to the specified direction.
func (p *Point) Move(direction Direction) {
	switch direction {
	case Left:
		p.x--
	case Right:
		p.x++
	case Up:
		p.y++
	case Down:
		p.y--
	case None:
	default:
		panic("cannot move to unknown direction")
	}
}

// Diff returns as a new Point the difference between the coordinates of p
// (the current point) and p2.
func (p *Point) Diff(p2 *Point) *Point {
	return &Point{
		x: p.x - p2.x,
		y: p.y - p2.y,
	}
}

// Follow moves the point so that it remains close to p2.
func (p *Point) Follow(p2 *Point) {
	diff := p2.Diff(p)

	// Do nothing if the points are close to each other.
	if common.Abs(diff.x) <= 1 && common.Abs(diff.y) <= 1 {
		return
	}

	// Horizontal and vertical directions.
	var (
		hd = None
		vd = None
	)

	switch {
	case diff.x < 0:
		hd = Left
	case diff.x > 0:
		hd = Right
	}

	switch {
	case diff.y < 0:
		vd = Down
	case diff.y > 0:
		vd = Up
	}

	p.Move(hd)
	p.Move(vd)
}

// parseInstruction parses the direction and the distance contained in an
// instruction.
func parseInstruction(line string) (direction Direction, distance int) {
	instruction := strings.Split(line, " ")

	direction = NewDirection(instruction[0][0])
	distance = common.MustAtoi(instruction[1])

	return
}

func part1(lines []string) int {
	var (
		head      = &Point{}
		tail      = &Point{}
		positions = make(map[Point]bool)
	)

	for _, line := range lines {
		dir, dist := parseInstruction(line)

		for i := 0; i < dist; i++ {
			head.Move(dir)
			tail.Follow(head)
			positions[*tail] = true
		}
	}

	return len(positions)
}

func part2(lines []string) int {
	var (
		knots     = make([]*Point, 10)
		positions = make(map[Point]bool)
	)

	for i := range knots {
		knots[i] = new(Point)
	}

	for _, line := range lines {
		dir, dist := parseInstruction(line)

		for i := 0; i < dist; i++ {
			knots[0].Move(dir)

			for j := 1; j < len(knots); j++ {
				knots[j].Follow(knots[j-1])
			}

			positions[*knots[len(knots)-1]] = true
		}
	}

	return len(positions)
}

func main() {
	var (
		lines = common.Lines(input)
	)

	fmt.Printf("Part one: %d\n", part1(lines)) // Part one: 6486
	fmt.Printf("Part two: %d\n", part2(lines)) // Part two: 2678
}
