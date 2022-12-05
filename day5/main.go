package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomy2e/advent-of-code-2022/common"
)

//go:embed input.txt
var input string

type MoveInstruction struct {
	Amount int
	Src    int
	Dst    int
}

type Stack []byte

// Push adds an element to the Stack.
func (s Stack) Push(v byte) Stack {
	return append(s, v)
}

// Pop removes the top element of the Stack and returns it.
func (s Stack) Pop() (Stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// Peek returns the top element of the stack without removing it. Panics if
// the Stack is empty.
func (s Stack) Peek() byte {
	return s[len(s)-1]
}

// Reverse implements this: https://github.com/golang/go/wiki/SliceTricks#reversing
func (s Stack) Reverse() {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// expandStacks makes sure that the number of stacks is at least equal to the
// size parameter.
func expandStacks(stacks []Stack, size int) []Stack {
	l := len(stacks)
	if l < size {
		for i := 0; i < size-l; i++ {
			stacks = append(stacks, make(Stack, 0))
		}
	}
	return stacks
}

// parseInput parse the lines of the input file and returns a list of Stacks
// and instructions to process.
func parseInput(lines []string) (stacks []Stack, instructions []MoveInstruction) {
	parseInstructions := false

	for _, line := range lines {
		if line == "" {
			parseInstructions = true
			continue
		}

		if parseInstructions {
			instructions = append(instructions, parseInstructionLine(line))
			continue
		}

		// We parse the stack line.
		crates, isCrateLine := parseCrateLine(line)
		if !isCrateLine {
			continue
		}

		stacks = expandStacks(stacks, len(crates))

		for i, crate := range crates {
			if crate != ' ' {
				stacks[i] = append(stacks[i], crate)
			}
		}
	}

	// Reverse all stacks as the lines are read in reverse order.
	for _, s := range stacks {
		s.Reverse()
	}

	return
}

func parseInstructionLine(line string) MoveInstruction {
	line = strings.Replace(line, "move ", "", 1)
	line = strings.Replace(line, "from ", "", 1)
	line = strings.Replace(line, "to ", "", 1)
	params := strings.Split(line, " ")

	return MoveInstruction{
		Amount: common.MustAtoi(params[0]),
		Src:    common.MustAtoi(params[1]),
		Dst:    common.MustAtoi(params[2]),
	}
}

func parseCrateLine(line string) (crates []byte, isCrateLine bool) {
	// Exit if this line does not contain any crate
	if strings.TrimSpace(line)[0] != '[' {
		return
	}

	isCrateLine = true

	for i := 0; i < len(line); i += 4 {
		crates = append(crates, line[i+1])
	}

	return
}

func processMoveInstruction(stacks []Stack, instruction MoveInstruction) {
	var b byte
	for i := 0; i < instruction.Amount; i++ {
		stacks[instruction.Src-1], b = stacks[instruction.Src-1].Pop()
		stacks[instruction.Dst-1] = stacks[instruction.Dst-1].Push(b)
	}
}

func processMoveInstructionRetainOrder(stacks []Stack, instruction MoveInstruction) {
	var b byte
	tmpStack := make(Stack, 0)

	for i := 0; i < instruction.Amount; i++ {
		stacks[instruction.Src-1], b = stacks[instruction.Src-1].Pop()
		tmpStack = tmpStack.Push(b)
	}

	for i := 0; i < instruction.Amount; i++ {
		tmpStack, b = tmpStack.Pop()
		stacks[instruction.Dst-1] = stacks[instruction.Dst-1].Push(b)
	}
}

// stacksToString concatenates the top element of each Stack into a string.
func stacksToString(stacks []Stack) string {
	sb := strings.Builder{}

	for _, stack := range stacks {
		sb.WriteByte(stack.Peek())
	}

	return sb.String()

}

func part1(lines []string) string {
	stacks, instructions := parseInput(lines)
	for _, instruction := range instructions {
		processMoveInstruction(stacks, instruction)
	}

	return stacksToString(stacks)
}

func part2(lines []string) string {
	stacks, instructions := parseInput(lines)
	for _, instruction := range instructions {
		processMoveInstructionRetainOrder(stacks, instruction)
	}

	return stacksToString(stacks)
}

func main() {
	var (
		lines = strings.Split(input, "\n")
	)

	fmt.Printf("Part one: %s\n", part1(lines)) // Part one: FRDSQRRCD
	fmt.Printf("Part two: %s\n", part2(lines)) // Part two: HRFTQVWNN
}
