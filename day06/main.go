package main

import (
	"fmt"
	"github.com/tlj/adventofcode2018/aoc"
	"math"
	"os"
	"strconv"
	"strings"
)

type criticalCoordinate struct {
	x int64
	y int64
	disabled bool
	ownsCount int64
	ownsEdge int64
}

type coordinateOwner struct {
	distance int64
	ownerCount int64
	owner int
	totalDistance int64
}

func getCriticalCoordinates(input []string) ([]criticalCoordinate, error) {
	criticalCoordinates := make([]criticalCoordinate, len(input))
	var err error

	for idx, line := range input {
		splitString := strings.Split(line, ", ")
		c := criticalCoordinate{}
		c.x, err = strconv.ParseInt(splitString[0], 10, 64)
		if err != nil {
			return nil, err
		}
		c.y, err = strconv.ParseInt(splitString[1], 10, 64)
		if err != nil {
			return nil, err
		}

		criticalCoordinates[idx] = c
	}

	return criticalCoordinates, nil
}

func viewSize(criticalCoordinates []criticalCoordinate) (int64, int64) {
	var x int64
	var y int64
	for _, criticalCoordinate := range criticalCoordinates {
		if criticalCoordinate.x > x {
			x = criticalCoordinate.x
		}
		if criticalCoordinate.y > y {
			y = criticalCoordinate.y
		}
	}

	return x, y
}

func distance(x1, y1 int, x2, y2 int64) int64 {
	return int64(math.Abs(float64(int64(x1)-x2))) + int64(math.Abs(float64(int64(y1)-y2)))
}

func isEdge(x, y int, maxX, maxY int64) bool {
	return x == 0 || int64(x) == maxX - 1 || int64(y) == maxY - 1 || y == 0
}

func main() {
	lines := aoc.LoadAsLines(input)
	criticalCoordinates, err := getCriticalCoordinates(lines)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		os.Exit(1)
	}

	maxX, maxY := viewSize(criticalCoordinates)
	maxX += 2
	maxY += 2

	view := make([][]*coordinateOwner, maxY)
	for i := range view {
		view[i] = make([]*coordinateOwner, maxX)
		for j := range view[i] {
			view[i][j] = &coordinateOwner{
				distance: -1,
				owner: -1,
				totalDistance: 0,
			}
		}
	}
	for y, row := range view {
		for x, col := range row {
			for idx, criticalCoordinate := range criticalCoordinates {
				dist := distance(x, y, criticalCoordinate.x, criticalCoordinate.y)

				col.totalDistance += dist

				if dist < col.distance || col.distance == -1 {
					if col.distance > -1 && col.owner > -1 {
						if isEdge(x, y, maxX, maxY) {
							criticalCoordinates[col.owner].ownsEdge--
						}
						criticalCoordinates[col.owner].ownsCount--
					}
					col.distance = dist
					col.owner = idx
					col.ownerCount = 1
					criticalCoordinates[idx].ownsCount++
					if isEdge(x, y, maxX, maxY) {
						criticalCoordinates[col.owner].ownsEdge++
					}
				} else if dist == col.distance && col.owner > -1 {
					criticalCoordinates[col.owner].ownsCount--
					if isEdge(x, y, maxX, maxY) {
						criticalCoordinates[col.owner].ownsEdge--
					}
					col.ownerCount++
					col.owner = -1
				}
			}
		}
	}

	topCriticalCoordinate := criticalCoordinate{}

	for _, criticalCoordinate := range criticalCoordinates {
		if criticalCoordinate.ownsEdge == 0 && criticalCoordinate.ownsCount > topCriticalCoordinate.ownsCount {
			topCriticalCoordinate = criticalCoordinate
		}
	}

	safeRegionSize := 0
	for _, row := range view {
		for _, col := range row {
			if col.totalDistance < 10000 {
				safeRegionSize++
			}
		}
	}

	fmt.Printf("Part1: %d\n", topCriticalCoordinate.ownsCount)
	fmt.Printf("Part2: %d\n", safeRegionSize)
}

var testInput1 = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

var input = `80, 357
252, 184
187, 139
101, 247
332, 328
302, 60
196, 113
271, 201
334, 89
85, 139
327, 161
316, 352
343, 208
303, 325
316, 149
270, 319
318, 153
257, 332
306, 348
299, 358
172, 289
303, 349
271, 205
347, 296
220, 276
235, 231
133, 201
262, 355
72, 71
73, 145
310, 298
138, 244
322, 334
278, 148
126, 135
340, 133
311, 118
193, 173
319, 99
50, 309
160, 356
155, 195
61, 319
80, 259
106, 318
49, 169
134, 61
74, 204
337, 174
108, 287`