package answers

import (
	"sort"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

type Basin struct {
	vectors []Vector
}

func (b Basin) isAdjacent(x int, y int) bool {
	for _, v := range b.vectors {
		if abs(v.x-x) == 1 && abs(v.y-y) == 0 {
			return true
		}
		if abs(v.x-x) == 0 && abs(v.y-y) == 1 {
			return true
		}
	}
	return false
}

func Day9() []int {
	input := ReadInputAsStr(9)
	// Converty from string to list of list of ints
	tubes := [][]int{}
	for _, row := range input {
		split := strings.Split(row, "")
		tubeInts := []int{}
		for _, elem := range split {
			parsedInt, _ := strconv.Atoi(elem)
			tubeInts = append(tubeInts, parsedInt)
		}
		tubes = append(tubes, tubeInts)
	}
	return []int{q9part1(tubes), q9part2(tubes)}
}

func getLowPoints(tubes [][]int) []Vector {
	// Returns x,y of all the low points
	xmax := len(tubes[0])
	ymax := len(tubes)
	solution := []Vector{}
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			cell := tubes[y][x]
			if x > 0 {
				if tubes[y][x-1] <= cell {
					continue
				}
			}
			if x < xmax-1 {
				if tubes[y][x+1] <= cell {
					continue
				}
			}
			if y > 0 {
				if tubes[y-1][x] <= cell {
					continue
				}
			}
			if y < ymax-1 {
				if tubes[y+1][x] <= cell {
					continue
				}
			}
			solution = append(solution, Vector{x, y})
		}
	}
	return solution
}

func q9part1(tubes [][]int) int {
	lowPoints := getLowPoints(tubes)
	counter := 0
	for _, vec := range lowPoints {
		counter += tubes[vec.y][vec.x] + 1
	}
	return counter
}

func q9part2(tubes [][]int) int {
	lowPoints := getLowPoints(tubes)
	xmax := len(tubes[0])
	ymax := len(tubes)
	basins := []Basin{}
	for _, point := range lowPoints {
		startingVector := []Vector{point}
		basins = append(basins, Basin{
			vectors: startingVector,
		})
	}

	vectors := []Vector{}
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			// Don't include low points
			isLowPoint := false
			for _, vector := range lowPoints {
				if x == vector.x && y == vector.y {
					isLowPoint = true
					break
				}
			}
			if tubes[y][x] == 9 || isLowPoint {
				continue
			}
			vectors = append(vectors, Vector{x: x, y: y})
		}
	}

	// Iterate until all points are put into a basin
	for len(vectors) > 0 {
		unprocessedVectors := []Vector{}
		for _, elem := range vectors {
			found := false
			for idx := range basins {
				if basins[idx].isAdjacent(elem.x, elem.y) {
					basins[idx].vectors = append(basins[idx].vectors, elem)
					found = true
				}
			}
			if found == false {
				unprocessedVectors = append(unprocessedVectors, elem)
			}
		}
		vectors = unprocessedVectors
	}

	basinSize := []int{}
	for _, basin := range basins {
		basinSize = append(basinSize, len(basin.vectors))
	}
	sort.Ints(basinSize)
	solution := basinSize[len(basinSize)-1] * basinSize[len(basinSize)-2] * basinSize[len(basinSize)-3]
	return solution
}
