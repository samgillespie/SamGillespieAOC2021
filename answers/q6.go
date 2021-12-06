package answers

func Day6() []int {
	data := ReadCSVAsInt(6)
	return []int{
		q6part1(data),
		q6part2(data),
	}
}

func SimulateStep(fish []int) []int {
	newFish := fish[0]
	// Drop everything by 1
	for i := 0; i < 8; i++ {
		fish[i] = fish[i+1]
	}
	fish[8] = 0

	// Move 0 into 6 and 8
	fish[6] += newFish
	fish[8] += newFish
	return fish
}

func q6part1(fishes []int) int {
	fishslice := make([]int, 9)
	for _, fish := range fishes {
		fishslice[fish]++
	}

	for i := 0; i < 80; i++ {
		fishslice = SimulateStep(fishslice)
	}
	return sumIntSlice(fishslice)
}

func q6part2(fishes []int) int {
	fishslice := make([]int, 9)
	for _, fish := range fishes {
		fishslice[fish]++
	}
	for i := 0; i < 256; i++ {
		fishslice = SimulateStep(fishslice)
	}
	return sumIntSlice(fishslice)
}

func sumIntSlice(vals []int) int {
	sum := 0
	for _, num := range vals {
		sum += num
	}
	return sum
}
