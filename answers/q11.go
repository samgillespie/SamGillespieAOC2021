package answers

import (
	"fmt"
	"strconv"
)

type Cell struct {
	x         int
	y         int
	value     int
	flashing  bool
	processed bool
}

func Day11() []int {
	data := ReadInputAsStr(11)
	cells := make([]Cell, 100)
	cursor := 0
	for y := range data {
		for x := range data[y] {
			value, _ := strconv.Atoi(string(data[y][x]))
			cell := Cell{
				x:         x,
				y:         y,
				value:     value,
				flashing:  false,
				processed: false,
			}
			cells[cursor] = cell
			cursor++
		}
	}

	return []int{q11part1(cells), q11part2(cells)}
}

func lookup(x int, y int) int {
	return y*10 + x
}

func PrintCells(cells []Cell) {
	intslices := make([]int, 100)
	for _, cell := range cells {
		intslices[lookup(cell.x, cell.y)] = cell.value
	}

	for i := 0; i < 10; i++ {
		var str string
		row := intslices[(10 * i):(10 * (i + 1))]
		for _, elem := range row {
			str += fmt.Sprintf("%d", elem)
		}
		fmt.Println(str)
	}

}

func IncrementNeighbours(cell Cell, cellSlice []Cell) ([]Cell, []Cell) {
	// Returns the all the cells, plus the newly flashing
	y := cell.y
	x := cell.x
	xmax := 10
	ymax := 10

	newFlash := []Cell{}

	if y > 0 {
		// Top Left
		if x > 0 {
			if cellSlice[lookup(x-1, y-1)].value == 9 {
				newFlash = append(newFlash, cellSlice[lookup(x-1, y-1)])
				cellSlice[lookup(x-1, y-1)].flashing = true
			}
			cellSlice[lookup(x-1, y-1)].value++
			//fmt.Println("Top Left")
		}

		// Top
		if cellSlice[lookup(x, y-1)].value == 9 {
			newFlash = append(newFlash, cellSlice[lookup(x, y-1)])
			cellSlice[lookup(x, y-1)].flashing = true
		}
		cellSlice[lookup(x, y-1)].value++
		//fmt.Println("Top")

		// Top Right
		if x < xmax-1 {
			if cellSlice[lookup(x+1, y-1)].value == 9 {
				newFlash = append(newFlash, cellSlice[lookup(x+1, y-1)])
				cellSlice[lookup(x+1, y-1)].flashing = true
			}
			cellSlice[lookup(x+1, y-1)].value++
			//fmt.Println("Top Right")
		}
	}
	// Left
	if x > 0 {
		if cellSlice[lookup(x-1, y)].value == 9 {
			newFlash = append(newFlash, cellSlice[lookup(x-1, y)])
			cellSlice[lookup(x-1, y)].flashing = true
		}
		cellSlice[lookup(x-1, y)].value++
		//fmt.Println("Left")
	}

	// right
	if x < xmax-1 {
		if cellSlice[lookup(x+1, y)].value == 9 {
			newFlash = append(newFlash, cellSlice[lookup(x+1, y)])
			cellSlice[lookup(x+1, y)].flashing = true
		}
		cellSlice[lookup(x+1, y)].value++
		//fmt.Println("Right")
	}

	if y < ymax-1 {
		// Bottom Left
		if x > 0 {
			if cellSlice[lookup(x-1, y+1)].value == 9 {
				newFlash = append(newFlash, cellSlice[lookup(x-1, y+1)])
				cellSlice[lookup(x-1, y+1)].flashing = true
			}
			cellSlice[lookup(x-1, y+1)].value++
			//fmt.Println("Bottom Left")
		}
		// Bottom
		if cellSlice[lookup(x, y+1)].value == 9 {
			newFlash = append(newFlash, cellSlice[lookup(x, y+1)])
			cellSlice[lookup(x, y+1)].flashing = true
		}
		//fmt.Println("Bottom")
		cellSlice[lookup(x, y+1)].value++

		//Bottom Right
		if x < xmax-1 {
			if cellSlice[lookup(x+1, y+1)].value == 9 {
				newFlash = append(newFlash, cellSlice[lookup(x+1, y+1)])
				cellSlice[lookup(x+1, y+1)].flashing = true
			}
			cellSlice[lookup(x+1, y+1)].value++
			//fmt.Println("Bottom Right")
		}
	}
	return cellSlice, newFlash
}

func SimulateOctopus(cells []Cell) ([]Cell, int) {
	// Increase by one
	isFlashing := []Cell{}
	for i := range cells {
		cells[i].value += 1
		cells[i].flashing = true
		if cells[i].value >= 10 {
			isFlashing = append(isFlashing, cells[i])
		}
	}
	// Add one to all
	cursor := 0
	for {
		// fmt.Println(isFlashing)
		if cursor >= len(isFlashing) {
			break
		}
		cell := isFlashing[cursor]
		var newFlashing []Cell
		cells, newFlashing = IncrementNeighbours(cell, cells)
		isFlashing = append(isFlashing, newFlashing...)
		cursor++
	}
	for i := range cells {
		if cells[i].value >= 10 {
			cells[i].value = 0
		}
	}
	return cells, len(isFlashing)
}

func q11part1(cells []Cell) int {
	//PrintCells(cells)
	flashes := 0
	for i := 0; i < 100; i++ {
		var flash int
		// PrintCells(cells)
		cells, flash = SimulateOctopus(cells)
		flashes += flash
		// fmt.Println("!!!!!!!!")
		// PrintCells(cells)
	}
	return flashes
}

func q11part2(cells []Cell) int {
	for i := 0; i < 2000; i++ {
		var flash int
		//
		cells, flash = SimulateOctopus(cells)
		if flash == 100 {
			// Naughty hack.  Since we start simulating from step 100
			// and +1 because the step that actually fires is the next one
			return i + 101
		}
	}
	return -1
}
