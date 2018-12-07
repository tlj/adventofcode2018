package main

import (
	"fmt"
	"github.com/tlj/adventofcode2018/aoc"
	"sort"
)

type letter struct {
	letter string
	requires []*letter
	done bool
	steps int
	stepsModifier int
}

func (l *letter) hasRequirement(search *letter) bool {
	for _, ll := range l.requires {
		if ll == search {
			return true
		}
	}
	return false
}

func (l *letter) addRequirement(add *letter) {
	if !l.hasRequirement(add) {
		l.requires = append(l.requires, add)
	}
}

func (l *letter) stepsLeft() int {
	return l.value() - l.steps
}

func (l *letter) waited() bool {
	if l.value() > l.steps {
		return false
	}
	return true
}

func (l *letter) ready() bool {
	for _, ll := range l.requires {
		if !ll.done {
			return false
		}
	}
	return true
}

func (l *letter) value() int {
	return int(l.letter[0]) - 64 + l.stepsModifier
}

func solve(lines []string, totalWorkers, stepsModifier int) (string, int) {
	letters := make(map[string]*letter)

	for _, line := range lines {
		first := string(line[5])
		second := string(line[36])

		_, ok1 := letters[first]
		if !ok1 {
			letters[first] = &letter{letter: first, stepsModifier:stepsModifier}
		}

		_, ok2 := letters[second]
		if !ok2 {
			letters[second] = &letter{letter: second, stepsModifier:stepsModifier}
		}

		letters[second].addRequirement(letters[first])
	}

	finished := false
	totalSteps := 0
	correctOrder := ""

	for !finished {
		var readyLetters []*letter
		for _, ll := range letters {
			if ll.ready() && !ll.done {
				readyLetters = append(readyLetters, ll)
			}
		}
		if len(readyLetters) == 0 {
			finished = true
			continue
		}

		sort.Slice(readyLetters, func(i, j int) bool {
			if readyLetters[i].steps > readyLetters[j].steps {
				return true
			}
			if readyLetters[i].steps < readyLetters[j].steps {
				return false
			}
			return readyLetters[i].letter < readyLetters[j].letter
		})

		if len(readyLetters) > totalWorkers {
			readyLetters = readyLetters[:totalWorkers]
		}
		totalSteps++

		for _, rl := range readyLetters {

			k := rl.letter
			letters[k].steps++
			if letters[k].ready() && !letters[k].done && letters[k].waited() {
				correctOrder += letters[k].letter
				letters[k].done = true
			}
		}
	}

	return correctOrder, totalSteps
}

func main() {
	lines := aoc.LoadAsLines(input)

	part1Order, _ := solve(lines, 2, 0)
	fmt.Printf("Part1: %s\n", part1Order)

	part2Order, part2Steps := solve(lines, 5, 60)
	fmt.Printf("Part2: %d (%s)\n", part2Steps, part2Order)
}

var testInput1 = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

var input = `Step Y must be finished before step A can begin.
Step O must be finished before step C can begin.
Step P must be finished before step A can begin.
Step D must be finished before step N can begin.
Step T must be finished before step G can begin.
Step L must be finished before step M can begin.
Step X must be finished before step V can begin.
Step C must be finished before step R can begin.
Step G must be finished before step E can begin.
Step H must be finished before step N can begin.
Step Q must be finished before step B can begin.
Step S must be finished before step R can begin.
Step M must be finished before step F can begin.
Step N must be finished before step Z can begin.
Step E must be finished before step I can begin.
Step A must be finished before step R can begin.
Step Z must be finished before step F can begin.
Step K must be finished before step V can begin.
Step I must be finished before step J can begin.
Step R must be finished before step W can begin.
Step B must be finished before step J can begin.
Step W must be finished before step V can begin.
Step V must be finished before step F can begin.
Step U must be finished before step F can begin.
Step F must be finished before step J can begin.
Step X must be finished before step C can begin.
Step T must be finished before step Q can begin.
Step B must be finished before step F can begin.
Step Y must be finished before step L can begin.
Step P must be finished before step E can begin.
Step A must be finished before step J can begin.
Step S must be finished before step I can begin.
Step S must be finished before step A can begin.
Step K must be finished before step R can begin.
Step D must be finished before step C can begin.
Step R must be finished before step U can begin.
Step K must be finished before step U can begin.
Step D must be finished before step K can begin.
Step S must be finished before step M can begin.
Step D must be finished before step E can begin.
Step A must be finished before step K can begin.
Step G must be finished before step I can begin.
Step O must be finished before step M can begin.
Step U must be finished before step J can begin.
Step T must be finished before step S can begin.
Step C must be finished before step M can begin.
Step S must be finished before step J can begin.
Step N must be finished before step V can begin.
Step P must be finished before step N can begin.
Step D must be finished before step M can begin.
Step A must be finished before step B can begin.
Step Z must be finished before step R can begin.
Step T must be finished before step N can begin.
Step K must be finished before step J can begin.
Step N must be finished before step A can begin.
Step M must be finished before step R can begin.
Step E must be finished before step A can begin.
Step Y must be finished before step O can begin.
Step O must be finished before step B can begin.
Step O must be finished before step A can begin.
Step I must be finished before step V can begin.
Step M must be finished before step Z can begin.
Step D must be finished before step U can begin.
Step O must be finished before step S can begin.
Step Z must be finished before step W can begin.
Step M must be finished before step A can begin.
Step N must be finished before step E can begin.
Step M must be finished before step U can begin.
Step R must be finished before step J can begin.
Step W must be finished before step F can begin.
Step I must be finished before step U can begin.
Step E must be finished before step U can begin.
Step Y must be finished before step R can begin.
Step Z must be finished before step K can begin.
Step C must be finished before step F can begin.
Step B must be finished before step V can begin.
Step G must be finished before step B can begin.
Step O must be finished before step G can begin.
Step E must be finished before step Z can begin.
Step A must be finished before step V can begin.
Step Y must be finished before step Q can begin.
Step P must be finished before step D can begin.
Step X must be finished before step G can begin.
Step I must be finished before step W can begin.
Step M must be finished before step V can begin.
Step T must be finished before step M can begin.
Step G must be finished before step J can begin.
Step T must be finished before step I can begin.
Step H must be finished before step B can begin.
Step C must be finished before step E can begin.
Step Q must be finished before step V can begin.
Step H must be finished before step U can begin.
Step X must be finished before step K can begin.
Step D must be finished before step T can begin.
Step X must be finished before step W can begin.
Step P must be finished before step Z can begin.
Step C must be finished before step U can begin.
Step Y must be finished before step Z can begin.
Step L must be finished before step F can begin.
Step C must be finished before step J can begin.
Step T must be finished before step W can begin.`