package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(2)
	q2part1(data)
	q2part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s\n", elapsed)
}

func q2part1(data []string) {
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
	fmt.Printf("Question 2 Part 1 Solution: %d\n", start_x*start_y)
}

func q2part2(data []string) {
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
	fmt.Printf("Question 2 Part 2 Solution: %d\n", start_x*start_y)
}
