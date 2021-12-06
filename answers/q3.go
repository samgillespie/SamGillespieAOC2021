package answers

import (
	"math"
)

var rowLength = 12

func Day3() []int {
	data := ReadInputAsStr(3)
	return []int{
		q3part1(data),
		q3part2(data),
	}
}

func q3part1(data []string) int {
	gammaRate := ""
	epsilonRate := ""
	sum_gamma := 0.0
	sum_epsilon := 0.0
	for index := 0; index < rowLength; index++ {
		ones := 0
		zeros := 0
		for _, row := range data {
			if row[index] == '1' {
				ones++
			} else {
				zeros++
			}
		}
		if ones > zeros {
			sum_gamma += math.Pow(2.0, float64(rowLength-1-index))
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			sum_epsilon += math.Pow(2.0, float64(rowLength-1-index))
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	return int(sum_gamma * sum_epsilon)
}

func calculate_oxygen(data []string) string {

	for index := 0; index < rowLength; index++ {
		ones := 0
		zeros := 0
		newList := []string{}
		for _, row := range data {
			if row[index] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		for _, row := range data {
			if ones >= zeros && row[index] == '1' {
				newList = append(newList, row)
			}
			if zeros > ones && row[index] == '0' {
				newList = append(newList, row)
			}
		}
		data = newList
		if len(data) == 1 {
			return data[0]
		}
	}
	return data[0]
}

func calculate_co2(data []string) string {
	for index := 0; index < rowLength; index++ {
		ones := 0
		zeros := 0
		newList := []string{}
		for _, row := range data {
			if row[index] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		for _, row := range data {
			if ones >= zeros && row[index] == '0' {
				newList = append(newList, row)
			}
			if zeros > ones && row[index] == '1' {
				newList = append(newList, row)
			}
		}
		data = newList
		if len(data) == 1 {
			return data[0]
		}
	}
	return data[0]
}

func binary_to_decimal(value string) int {
	sum := 0.0
	for index := 0; index < rowLength; index++ {
		if value[index] == '1' {
			sum += math.Pow(2.0, float64(rowLength-1-index))
		}
	}
	return int(sum)
}

func q3part2(data []string) int {
	oxy := calculate_oxygen(data)
	co2 := calculate_co2(data)
	solution := binary_to_decimal(oxy) * binary_to_decimal(co2)
	return solution
}
