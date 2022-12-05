package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Outcome is the outcome of the round.
type Outcome int

const (
	Win  Outcome = 6
	Draw Outcome = 3
	Loss Outcome = 0
)

// NewOutcome returns the outcome associated to the letter.
func NewOutcome(outcome byte) Outcome {
	switch outcome {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}

	panic("unknown outcome")
}

// Shape is a Rock Paper Scissors hand shape.
type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

// NewShape returns the shape associated to the player or opponent letter.
func NewShape(play byte) Shape {
	switch play {
	case 'X', 'A':
		return Rock
	case 'Y', 'B':
		return Paper
	case 'Z', 'C':
		return Scissors
	}

	panic("unknown shape")
}

// outcome returns the outcome for the player.
func outcome(opponent, player Shape) Outcome {
	if opponent == player {
		return Draw
	}

	if shouldPlayToWin(opponent) == player {
		return Win
	}

	return Loss
}

// shouldPlayToWin returns the shape that needs to be played in order to win.
func shouldPlayToWin(shape Shape) Shape {
	return shape%3 + 1
}

// shouldPlayToLose returns the shape that needs to be played in order to lose.
func shouldPlayToLose(shape Shape) Shape {
	return (shape+1)%3 + 1
}

// outcomePlay returns what the player needs to play to meet the desired outcome.
func outcomePlay(opponent Shape, outcome Outcome) Shape {
	switch outcome {
	case Draw:
		return opponent
	case Win:
		return shouldPlayToWin(opponent)
	case Loss:
		return shouldPlayToLose(opponent)
	}

	panic("unknown outcome")
}

func main() {
	var (
		lines      = strings.Split(input, "\n")
		totalScore = 0
	)

	// Part one.
	for _, line := range lines {
		cols := strings.Split(line, " ")
		opponent := NewShape(cols[0][0])
		player := NewShape(cols[1][0])
		totalScore += int(player) + int(outcome(opponent, player))
	}

	fmt.Printf("Part one: %d\n", totalScore)

	// Part two.
	totalScore = 0

	for _, line := range lines {
		cols := strings.Split(line, " ")
		opponent := NewShape(cols[0][0])
		desiredOutcome := NewOutcome(cols[1][0])
		player := outcomePlay(opponent, desiredOutcome)
		totalScore += int(player) + int(outcome(opponent, player))
	}

	fmt.Printf("Part two: %d\n", totalScore)
}
