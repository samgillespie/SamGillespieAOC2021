package answers

import (
	"strconv"
	"strings"
)

func Day2() []int {
	data := ReadInputAsStr(2)
	return []int{
		q2part1(data),
		q2part2(data),
	}
}

func q2part1(data []string) int {
	start_x := 0
	start_y := 0

	for _, row := range data {
		direction := strings.Split(row, " ")[0]
		distance, _ := strconv.Atoi(strings.Split(row, " ")[1])
		if direction == "forward" {
			start_x += distance
		}
		if direction == "backward" {
			start_x += -distance
		}
		if direction == "down" {
			start_y += distance
		}
		if direction == "up" {
			start_y += -distance
		}
	}
	return start_x * start_y
}

func q2part2(data []string) int {
	start_x := 0
	start_y := 0
	aim := 0

	for _, row := range data {
		direction := strings.Split(row, " ")[0]
		distance, _ := strconv.Atoi(strings.Split(row, " ")[1])
		if direction == "forward" {
			start_x += distance
			start_y += aim * distance
		}
		if direction == "backward" {
			start_x += -distance
			start_y += aim * -distance
		}
		if direction == "down" {
			aim += distance
		}
		if direction == "up" {
			aim += -distance
		}
	}
	return start_x * start_y
}
