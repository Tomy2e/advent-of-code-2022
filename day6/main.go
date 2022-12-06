package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Buffer is a buffer that discards older data when new data is added.
type Buffer struct {
	b    []byte
	size int
}

// NewBuffer returns a new buffer with a maximum size.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		b:    make([]byte, 0),
		size: size,
	}
}

// Push pushes a new byte to the buffer.
func (b *Buffer) Push(p byte) {
	b.b = append(b.b, p)

	if l := len(b.b); l > b.size {
		b.b = b.b[l-b.size:]
	}
}

// IsMarker returns true if a marker is detected (i.e. there are b.size distinct
// characters in the buffer).
func (b Buffer) IsMarker() bool {
	if len(b.b) != b.size {
		return false
	}

	for _, c := range b.b {
		if bytes.Count(b.b, []byte{c}) > 1 {
			return false
		}
	}

	return true
}

func part1and2(line string, distinctCharacters int) int {
	b := NewBuffer(distinctCharacters)
	for i, c := range line {
		b.Push(byte(c))
		if b.IsMarker() {
			return i + 1
		}
	}

	panic("marker not found")
}

func main() {
	var (
		lines = strings.Split(input, "\n")
	)

	fmt.Printf("Part one: %d\n", part1and2(lines[0], 4))  // Part one: 1262
	fmt.Printf("Part two: %d\n", part1and2(lines[0], 14)) // Part two: 3444

}
