package answers

func Day7() []int {
	data := ReadCSVAsInt(7)
	return []int{q7part1(data), q7part2(data)}

}

func q7part1(data []int) int {

	_, totalCrabDistance := maxSlice(data)

	allfuel := make([]int, totalCrabDistance)
	for i := 0; i < totalCrabDistance; i++ {
		fuel := 0
		for _, elem := range data {
			fuel += abs(elem - i)
		}
		allfuel[i] = fuel
	}
	_, minFuel := minSlice(allfuel)
	return minFuel
}

func triangleSlice(steps int) []int {
	// Returns an array of triangle numbers such that 1:1, 2:3, 3:6, etc
	tris := make([]int, steps+1)
	for i := 1; i <= steps; i++ {
		tris[i] = tris[i-1] + i
	}
	return tris
}
func q7part2(data []int) int {

	_, totalCrabDistance := maxSlice(data)

	triangleNumbers := triangleSlice(totalCrabDistance)

	allfuel := make([]int, totalCrabDistance)
	for i := 0; i < totalCrabDistance; i++ {
		fuel := 0
		for _, elem := range data {
			fuel += triangleNumbers[abs(elem-i)]
		}
		allfuel[i] = fuel
	}
	_, minFuel := minSlice(allfuel)
	return minFuel
}
