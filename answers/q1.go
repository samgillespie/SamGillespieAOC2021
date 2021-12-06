package answers

func Day1() []int {
	data := ReadInputAsInt(1)
	return []int{q1part1(data), q1part2(data)}

}

func q1part1(data []int) int {
	prev_value := data[0]
	counter := 0
	for i := 1; i < len(data); i++ {
		if data[i] > prev_value {
			counter++
		}
		prev_value = data[i]
	}
	return counter
}

func q1part2(data []int) int {
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
	return counter
}
