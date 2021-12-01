package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(1)
	q1part1(data)
	q1part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q1part1(data []int) {
	prev_value := data[0]
	counter := 0
	for i := 1; i < len(data); i++ {
		if data[i] > prev_value {
			counter++
		}
		prev_value = data[i]
	}
	fmt.Printf("Question 1 Part 1 Solution: %d\n", counter)
}

func q1part2(data []int) {
	// Part 2
	prev_value := data[0] + data[1] + data[2]
	counter := 0
	for i := 1; i < len(data)-2; i++ {
		next_value := data[i] + data[i+1] + data[i+2]
		if next_value > prev_value {
			counter++
		}
		prev_value = next_value
	}
	fmt.Printf("Question 1 Part 1 Solution: %d\n", counter)
}
