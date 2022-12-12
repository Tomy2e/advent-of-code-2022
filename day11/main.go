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

// MonkeyOperationFunc is an operation that a monkey can do.
type MonkeyOperationFunc func(int) int

// mult returns a function that multiplies a number by n.
func mult(n int) MonkeyOperationFunc {
	return func(a int) int {
		return a * n
	}
}

// add returns a function that adds n to a number.
func add(n int) MonkeyOperationFunc {
	return func(a int) int {
		return a + n
	}
}

// square is a function that squares a number.
var square MonkeyOperationFunc = func(a int) int {
	return a * a
}

// ReducerFunc is a function that reduces the worry level.
type ReducerFunc func(int) int

// part1Reducer reduces the worry level by dividing it by 3. This function
// is used in part 1.
var part1Reducer ReducerFunc = func(a int) int {
	return a / 3
}

// part2Reducer returns a function that uses the least common multiple to reduce
// the worry level. This function is used in part 2.
func part2Reducer(lcm int) ReducerFunc {
	return func(a int) int {
		return a % lcm
	}
}

// monkeysLCM finds the least common multiple of the monkeys test number.
// We assume they are always prime numbers so we just return their product.
func monkeysLCM(monkeys []*Monkey) int {
	lcm := 1

	for _, m := range monkeys {
		lcm *= m.test
	}

	return lcm
}

type Monkey struct {
	items       []int
	operation   MonkeyOperationFunc
	test        int
	monkeyTrue  int
	monkeyFalse int
}

// NewMonkey parses the lines that describe the monkey and returns a new
// instance of it.
func NewMonkey(monkeyLines []string) *Monkey {
	m := &Monkey{}

	// Starting items (line 1).
	s := strings.Split(monkeyLines[1], ": ")
	s = strings.Split(s[1], ", ")

	for _, item := range s {
		m.items = append(m.items, common.MustAtoi(item))
	}

	// Operation (line 2).
	s = strings.Split(monkeyLines[2], " ")

	switch {
	case s[6] == "*" && s[7] == "old":
		m.operation = square
	case s[6] == "*":
		m.operation = mult(common.MustAtoi(s[7]))
	case s[6] == "+":
		m.operation = add(common.MustAtoi(s[7]))
	}

	// Test (line 3).
	s = strings.Split(monkeyLines[3], "by ")
	m.test = common.MustAtoi(s[1])

	// If true (line 4).
	s = strings.Split(monkeyLines[4], "monkey ")
	m.monkeyTrue = common.MustAtoi(s[1])

	// If false (line 5).
	s = strings.Split(monkeyLines[5], "monkey ")
	m.monkeyFalse = common.MustAtoi(s[1])

	return m
}

// CanThrow returns true if the monkey has at least one item.
func (m *Monkey) CanThrow() bool {
	return len(m.items) != 0
}

// Throw throws an item to another monkey. It takes a reducer function so that
// this method works for part 1 and 2. It returns the new level of the item and
// the index of the monkey that receives the item.
func (m *Monkey) Throw(reduce ReducerFunc) (level, monkey int) {
	level = reduce(m.operation(m.items[0]))

	if level%m.test == 0 {
		monkey = m.monkeyTrue
	} else {
		monkey = m.monkeyFalse
	}

	m.items = m.items[1:len(m.items)]

	return
}

// Receive receives an item from another monkey.
func (m *Monkey) Receive(item int) {
	m.items = append(m.items, item)
}

// parseInput parses the input file and returns a list of monkeys.
func parseInput(lines []string) (m []*Monkey) {
	for i := 0; i < len(lines); i += 7 {
		m = append(m, NewMonkey(lines[i:i+6]))
	}

	return
}

func part1and2(lines []string, rounds int, reducer func([]*Monkey) ReducerFunc) int {
	var (
		monkeys  = parseInput(lines)
		inspects = make([]int, len(monkeys))
	)

	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			for monkey.CanThrow() {
				inspects[j]++
				l, m := monkey.Throw(reducer(monkeys))
				monkeys[m].Receive(l)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))

	return inspects[0] * inspects[1]
}

func main() {
	var (
		lines = common.Lines(input)
	)

	fmt.Printf("Part one: %d\n", part1and2(lines, 20, func(m []*Monkey) ReducerFunc {
		return part1Reducer
	})) // Part one: 112815

	fmt.Printf("Part two: %d\n", part1and2(lines, 10000, func(m []*Monkey) ReducerFunc {
		return part2Reducer(monkeysLCM(m))
	})) // Part two: 25738411485
}
