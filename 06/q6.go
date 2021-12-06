package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadCSVAsInt(6)
	q6part1(data)
	q6part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func SimulateStep(fish map[int]int) map[int]int {
	newMap := make(map[int]int)
	// Drop everything by 1
	for i := 1; i <= 8; i++ {
		newMap[i-1] = fish[i]
	}

	// Move 0 into 6 and 8
	newMap[6] += fish[0]
	newMap[8] += fish[0]
	return newMap
}

func q6part1(fishes []int) {
	fishmap := make(map[int]int)
	for _, fish := range fishes {
		fishmap[fish]++
	}

	for i := 0; i < 80; i++ {
		fishmap = SimulateStep(fishmap)
	}

	totalFish := 0
	for _, fish := range fishmap {
		totalFish += fish
	}
	fmt.Printf("Question 6 Part 1 Answer %d\n", totalFish)
}

func q6part2(fishes []int) {
	fishmap := make(map[int]int)
	for _, fish := range fishes {
		fishmap[fish]++
	}

	for i := 0; i < 256; i++ {
		fishmap = SimulateStep(fishmap)
	}

	totalFish := 0
	for _, fish := range fishmap {
		totalFish += fish
	}
	fmt.Printf("Question 6 Part 2 Answer %d\n", totalFish)
}
