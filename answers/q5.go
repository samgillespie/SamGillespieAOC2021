package answers

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Line struct {
	xstart  int
	xfinish int
	ystart  int
	yfinish int
}

func Print(vents map[Coord]int) {
	// For printing vents
	for y := 0; y < 11; y++ {
		row := []string{}
		for x := 0; x < 11; x++ {
			value := fmt.Sprintf("%d", vents[Coord{x: x, y: y}])
			row = append(row, value)
		}
		fmt.Println(strings.Join(row[:], ""))
	}
}

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func ApplyToVent(vent map[Coord]int, line Line, includeDiagonal bool) map[Coord]int {
	if line.ystart == line.yfinish {
		for x := Min(line.xstart, line.xfinish); x <= Max(line.xstart, line.xfinish); x++ {
			coord := Coord{x: x, y: line.ystart}
			// If coord not in vent, vent[coord] = 0
			vent[coord] = vent[coord] + 1
		}
	} else if line.xstart == line.xfinish {
		for y := Min(line.ystart, line.yfinish); y <= Max(line.ystart, line.yfinish); y++ {
			coord := Coord{x: line.xstart, y: y}
			vent[coord] = vent[coord] + 1
		}
	} else if includeDiagonal == true {
		numberOfSteps := Max(line.ystart, line.yfinish) - Min(line.ystart, line.yfinish)
		y := line.ystart
		x := line.xstart

		y_ispos := line.yfinish > line.ystart
		x_ispos := line.xfinish > line.xstart
		for i := 0; i <= numberOfSteps; i++ {
			coord := Coord{x: x, y: y}
			vent[coord] = vent[coord] + 1
			if y_ispos == true {
				y++
			} else {
				y--
			}
			if x_ispos {
				x++
			} else {
				x--
			}

		}
	}
	return vent
}

func ParseInput(input []string) []Line {
	lines := []Line{}
	for _, row := range input {
		rowStr := strings.Replace(row, " -> ", ",", 1)
		rowSplit := strings.Split(rowStr, ",")
		rowInts := []int{}
		for _, elem := range rowSplit {
			elemInt, _ := strconv.Atoi(elem)
			rowInts = append(rowInts, elemInt)
		}
		lines = append(lines, Line{
			xstart:  rowInts[0],
			ystart:  rowInts[1],
			xfinish: rowInts[2],
			yfinish: rowInts[3],
		})
	}
	return lines
}

func Day5() []int {
	data := ReadInputAsStr(5)
	return []int{
		q5part1(data),
		q5part2(data),
	}

}

func q5part1(data []string) int {
	lines := ParseInput(data)
	vents := make(map[Coord]int)
	for _, line := range lines {
		vents = ApplyToVent(vents, line, false)
	}
	counter := 0
	for _, overlap := range vents {

		if overlap >= 2 {
			counter++
		}
	}
	return counter
}

func q5part2(data []string) int {
	lines := ParseInput(data)
	vents := make(map[Coord]int)
	for _, line := range lines {
		vents = ApplyToVent(vents, line, true)
	}
	counter := 0
	for _, overlap := range vents {

		if overlap >= 2 {
			counter++
		}
	}
	return counter
}
