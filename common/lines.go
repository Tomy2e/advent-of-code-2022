package common

import (
	"bufio"
	"strings"
)

// Lines splits an input string into a slice of lines.
func Lines(input string) (lines []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}
