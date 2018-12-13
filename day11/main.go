package main

import "fmt"

var serialNumber = 6042
var xSize = 300
var ySize = 300

type cell struct {
	x int
	y int
	powerLevel int
	totalPowerLevel int
	maxTotalPowerLevelSize int
	maxTotalPowerLevel int
	neighbours [8]*cell
}

func (c *cell) setPowerLevel() int {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackId := c.x + 10

	// Begin with a power level of the rack ID times the Y coordinate.
	c.powerLevel = rackId * c.y

	// Increase the power level by the value of the grid serial number (your puzzle input).
	c.powerLevel += serialNumber

	// Set the power level to itself multiplied by the rack ID.
	c.powerLevel *= rackId

	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	if c.powerLevel < 100 {
		c.powerLevel = 0
	} else {
		c.powerLevel /= 100
		c.powerLevel = c.powerLevel % 10
	}

	// Subtract 5 from the power level.
	c.powerLevel -= 5

	return c.powerLevel
}

func (c *cell) getTotalPowerLevel(maxSquareSize int, fuelCells [][]*cell) int {
	totalPowerLevel := c.powerLevel

	for i := 0; i < maxSquareSize; i++ {
		for j := 0; j < maxSquareSize; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if c.x - 1 + j < xSize && c.y - 1 + i < ySize {
				totalPowerLevel += fuelCells[c.y - 1 + i][c.x - 1 + j].powerLevel
			}
		}
	}

	return totalPowerLevel
}

func (c *cell) setMaxTotalPowerLevel(fuelCells [][]*cell) {
	prev := c.getTotalPowerLevel(0, fuelCells)
	next := prev
	bestCount := 0
	lowerCount := 0
	for count := 1; count < 300; count++ {
		next = c.getTotalPowerLevel(count, fuelCells)
		if next < prev {
			lowerCount++
		} else if next > prev {
			prev = next
			bestCount = count
			lowerCount = 0
		}
		if lowerCount > 10 {
			break
		}
	}
	c.maxTotalPowerLevel = prev
	c.maxTotalPowerLevelSize = bestCount
}


func (c *cell) setTotalPowerLevel() {
	c.totalPowerLevel = c.powerLevel
	for _, n := range c.neighbours {
		if n != nil {
			c.totalPowerLevel += n.powerLevel
		}
	}
}

func main() {
	fuelCells := make([][]*cell, ySize)
	for i := range fuelCells {
		fuelCells[i] = make([]*cell, xSize)
		for j := range fuelCells[i] {
			fuelCells[i][j] = &cell{
			}
		}
	}

	topPowerLevelCell := cell{}
	topPowerLevelCell2 := cell{}
	for y := len(fuelCells)-1; y >= 0; y-- {
		for x := len(fuelCells[y])-1; x >= 0; x-- {
			cell := fuelCells[y][x]

			cell.x = x + 1
			cell.y = y + 1
			cell.setPowerLevel()

			// Part1
			count := 0
			maxSquareSize := 3
			for i := 0; i < maxSquareSize; i++ {
				for j := 0; j < maxSquareSize; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if x+j < xSize && y+i < ySize {
						cell.neighbours[count] = fuelCells[y + i][x + j]
						count++
					}
				}
			}
			cell.setTotalPowerLevel()

			if cell.totalPowerLevel > topPowerLevelCell.totalPowerLevel {
				topPowerLevelCell = *cell
			}

			// Part2
			cell.setMaxTotalPowerLevel(fuelCells)
			if cell.maxTotalPowerLevel > topPowerLevelCell2.maxTotalPowerLevel {
				topPowerLevelCell2 = *cell
			}

		}

	}

	fmt.Printf("Part1: %d,%d (%d)\n", topPowerLevelCell.x, topPowerLevelCell.y, topPowerLevelCell.totalPowerLevel)
	fmt.Printf("Part2: %d,%d,%d (%d)\n", topPowerLevelCell2.x, topPowerLevelCell2.y, topPowerLevelCell2.maxTotalPowerLevelSize, topPowerLevelCell.maxTotalPowerLevel)

}

