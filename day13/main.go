package main

import (
	"fmt"
	"github.com/tlj/adventofcode2018/aoc"
	"log"
	"sort"
)

const LEFT = 0
const UP = 1
const RIGHT = 2
const DOWN = 3

const FIRST = -1
const SECOND = 0
const THIRD = 1

type train struct {
	direction int
	x int
	y int
	nextDirection int
	crashed bool
}

func (t *train) step(view *[]string) {
	switch t.direction {
	case RIGHT:
		t.x += 1
	case DOWN:
		t.y += 1
	case LEFT:
		t.x -= 1
	case UP:
		t.y -= 1
	}
	road := (*view)[t.y][t.x]
	switch string(road) {
	case "+":
		t.direction += t.nextDirection
		if t.direction < LEFT { t.direction = DOWN }
		if t.direction > DOWN { t.direction = LEFT }
		t.nextDirection++
		if t.nextDirection > THIRD { t.nextDirection = FIRST }
	case "\\":
		if t.direction == LEFT  { t.direction = UP    } else
		if t.direction == DOWN  { t.direction = RIGHT } else
		if t.direction == UP    { t.direction = LEFT  } else
		if t.direction == RIGHT { t.direction = DOWN  }
	case "/":
		if t.direction == LEFT  { t.direction = DOWN  } else
		if t.direction == DOWN  { t.direction = LEFT  } else
		if t.direction == UP 	{ t.direction = RIGHT } else
		if t.direction == RIGHT { t.direction = UP    }
	case " ":
		log.Fatalf("Standing in an empty spot: %v", t)

	}
}

type byPos []*train

func (s byPos) Len() int {
	return len(s)
}
func (s byPos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPos) Less(i, j int) bool {
	if s[i].x == s[j].x {
		return s[i].y < s[j].y
	}
	return s[i].x < s[j].x
}

func main() {
	lines := aoc.ReadFileAsLines("./day13/input")

	trains := make([]*train, 0)

	for y := range lines {
		for x := range lines[y] {
			switch string(lines[y][x]) {
			case "v":
				trains = append(trains, &train{direction: DOWN, x: x, y: y, nextDirection:FIRST})
			case ">":
				trains = append(trains, &train{direction: RIGHT, x: x, y: y, nextDirection:FIRST})
			case "<":
				trains = append(trains, &train{direction: LEFT, x: x, y: y, nextDirection:FIRST})
			case "^":
				trains = append(trains, &train{direction: UP, x: x, y: y, nextDirection:FIRST})
			}
		}
	}

	crashes := 0
	stepCount := 0
	aliveCount := len(trains)
	for aliveCount > 1 {
		sort.Sort(byPos(trains))
		for idx, train := range trains {
			if train.crashed { continue }
			train.step(&lines)
			for idx2, train2 := range trains {
				if idx2 == idx { continue }
				if train2.crashed { continue }

				if train.x == train2.x && train.y == train2.y {
					crashes++
					if crashes == 1 {
						fmt.Printf("Part1: %d,%d (%d)\n", train.x, train.y, stepCount)
					}
					trains[idx].crashed = true
					trains[idx2].crashed = true
				}
			}
		}
		aliveCount = 0
		for _, train := range trains {
			if !train.crashed {
				aliveCount++
			}
		}
		stepCount++

	}
	for _, train := range trains {
		if !train.crashed {
			fmt.Printf("Part2: %d,%d (%d)\n", train.x, train.y, stepCount)
		}
	}

}
