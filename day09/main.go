package main

import (
	"container/list"
	"fmt"
	"time"
)

func play(playerCount int, lastMarbleWorth int) (int, int) {
	l := list.New()

	var currentElement *list.Element
	currentElement = l.PushFront(0)

	players := make([]int, playerCount)

	marbleWorth := 1

	for marbleWorth <= lastMarbleWorth {
		for i := 0; i < playerCount; i++ {
			if marbleWorth % 23 == 0 {
				players[i] += marbleWorth
				var removeElement *list.Element
				for j := 0; j < 7; j++ {
					if removeElement != nil {
						removeElement = removeElement.Prev()
					} else {
						removeElement = currentElement.Prev()
					}
					if removeElement == nil {
						removeElement = l.Back()
					}
				}
				currentElement = removeElement.Next()
				if currentElement == nil {
					currentElement = l.Front()
				}
				players[i] += removeElement.Value.(int)
				l.Remove(removeElement)
			} else {
				nextElement := currentElement.Next()
				if nextElement == nil {
					nextElement = l.Front()
				}
				currentElement = l.InsertAfter(marbleWorth, nextElement)
			}

			marbleWorth++

			if marbleWorth > lastMarbleWorth {
				break
			}
		}
	}

	var topPlayer int
	var topPlayerScore int
	for p := 0; p < playerCount; p++ {
		if players[p] > topPlayerScore {
			topPlayerScore = players[p]
			topPlayer = p
		}
	}

	return topPlayer, topPlayerScore
}

func main() {
	start := time.Now()

	_, part1 := play(423, 71944)
	fmt.Printf("Part1: %d\n", part1)

	fmt.Printf("\nPart 1 Finished in %s\n\n", time.Since(start))

	start = time.Now()

	_, part2 := play(423, 71944 * 100)
	fmt.Printf("Part2: %d\n", part2)

	fmt.Printf("\nPart 1 Finished in %s\n\n", time.Since(start))
}