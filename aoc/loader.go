package aoc

import (
	"bufio"
	"fmt"
	"strings"
)

type Loader struct {

}

func LoadAsLines(input string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("\nLoaded %d lines.\n\n", len(lines))

	return lines
}
