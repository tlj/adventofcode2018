package aoc

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	Sum int64
}

func (c *Calculator) Add(num int64) {
	c.Sum += num
}

func (c *Calculator) Sub(num int64) {
	c.Sum -= num
}

func (c *Calculator) ParseLine(line string) error {
	instruction := line[:1]
	number, err := strconv.ParseInt(line[1:], 10, 32)
	if err != nil {
		return fmt.Errorf("Unable to parse line: %s.\n", line)
	}

	switch instruction {
	case "+":
		c.Add(number)
	case "-":
		c.Sub(number)
	default:
		return fmt.Errorf("Unknown instruction %s.\n", instruction)
	}

	return nil
}