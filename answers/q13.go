package answers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Dot struct {
	x int
	y int
}

type Fold struct {
	x int
	y int
}

type Dots []*Dot
type Folds []Fold

func (d Dots) Print() {
	xmax := 0
	ymax := 0
	for _, dot := range d {
		if dot.x > xmax {
			xmax = dot.x
		}
		if dot.y > ymax {
			ymax = dot.y
		}
	}
	diagram := [][]rune{}
	for y := 0; y <= ymax; y++ {
		str := []rune{}
		for x := 0; x <= xmax; x++ {
			str = append(str, '.')
		}
		diagram = append(diagram, str)
	}

	for _, dot := range d {
		diagram[dot.y][dot.x] = '#'
	}
	for _, row := range diagram {
		fmt.Println(string(row))
	}

}

func (d Dots) Fold(x int, y int) Dots {
	for i := range d {
		d[i] = d[i].Fold(x, y)
	}
	return d.RemoveDuplicates()
}

func (d Dots) RemoveDuplicates() Dots {
	dots := map[Dot]bool{}
	toRemove := []int{}
	for index, dot := range d {
		_, exists := dots[*dot]
		if exists == true {
			toRemove = append(toRemove, index)
		} else {
			dots[*dot] = true
		}
	}
	sort.Ints(toRemove)
	for i := len(toRemove) - 1; i >= 0; i-- {
		elem := toRemove[i]
		d = append(d[0:elem], d[elem+1:]...)
	}
	return d
}

func (d *Dot) Fold(x int, y int) *Dot {
	if x == 0 {
		// Fold right to left
		dist := d.y - y
		if d.y > y {
			d.y = d.y - 2*dist
		}
	} else {
		dist := d.x - x
		if d.x > x {
			d.x = d.x - 2*dist
		}
	}
	return d
}

func Day13() []int {
	data := ReadInputAsStr(13)
	dots := Dots{}
	folds := Folds{}
	atDots := true
	for _, row := range data {
		if row == "" {
			atDots = false
			continue
		}
		if atDots == true {
			splitStr := strings.Split(row, ",")
			x, _ := strconv.Atoi(splitStr[0])
			y, _ := strconv.Atoi(splitStr[1])
			dots = append(dots, &Dot{x, y})
		} else {
			splitStr := strings.Split(row, " ")
			splitStr2 := strings.Split(splitStr[2], "=")
			if splitStr2[0] == "x" {
				foldx, _ := strconv.Atoi(splitStr2[1])
				folds = append(folds, Fold{x: foldx})
			} else {
				foldy, _ := strconv.Atoi(splitStr2[1])
				folds = append(folds, Fold{y: foldy})
			}
		}

	}

	return []int{q13part1(dots, folds), q13part2(dots, folds)}
}

func q13part1(dots Dots, folds Folds) int {
	dots = dots.Fold(folds[0].x, folds[0].y)
	return len(dots)
}

func q13part2(dots Dots, folds Folds) int {
	for _, fold := range folds {
		dots = dots.Fold(fold.x, fold.y)
	}
	dots.Print()
	return -1
}
