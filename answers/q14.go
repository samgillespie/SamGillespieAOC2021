package answers

import (
	"strings"
)

type Polymer map[string]int

type Instruction struct {
	condition string
	result    string
}

func Day14() []int {
	data := ReadInputAsStr(14)

	initialLetter := data[0][0:1]
	polymer := Polymer{}
	for i := 0; i < len(data[0])-1; i++ {
		polymer[data[0][i:i+2]] += 1
	}

	instructions := []Instruction{}
	for _, row := range data[2:] {
		rowSplit := strings.Split(row, " -> ")
		instructions = append(instructions, Instruction{
			condition: row[0:2],
			result:    rowSplit[1],
		})
	}
	return []int{q14part1(polymer, instructions, initialLetter), q14part2(polymer, instructions, initialLetter)}
}

func ApplyInstructions(polymer Polymer, instructions []Instruction) Polymer {
	additions := Polymer{}
	for pair, polymerCount := range polymer {
		for _, instruction := range instructions {
			if instruction.condition == pair {
				additions[pair[0:1]+instruction.result] += polymerCount
				additions[instruction.result+pair[1:]] += polymerCount
				additions[pair] -= polymerCount
			}
		}
	}
	for elem, value := range additions {
		polymer[elem] += value
	}
	return polymer
}

func ConvertToLetterCount(input map[string]int) map[string]int {
	// Only count the second letter, because the first will always be the same
	result := map[string]int{}
	for str, count := range input {
		result[str[1:]] += count
	}
	return result
}

func SumElements(input map[string]int) int {
	sum := 0
	for _, value := range input {
		sum += value
	}
	return sum
}

func MinMaxInMap(mapItem map[string]int) (int, int) {
	maxValue := -99999999999999
	minValue := 999999999999999
	for _, value := range mapItem {
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}
	return minValue, maxValue
}

func q14part1(polymer Polymer, instructions []Instruction, initialLetter string) int {
	for i := 0; i < 10; i++ {
		polymer = ApplyInstructions(polymer, instructions)
	}
	letterCount := ConvertToLetterCount(polymer)
	letterCount[initialLetter] += 1
	minValue, maxValue := MinMaxInMap(letterCount)

	return maxValue - minValue
}

func q14part2(polymer Polymer, instructions []Instruction, initialLetter string) int {
	for i := 0; i < 30; i++ {
		polymer = ApplyInstructions(polymer, instructions)
	}
	letterCount := ConvertToLetterCount(polymer)
	letterCount[initialLetter] += 1
	minValue, maxValue := MinMaxInMap(letterCount)

	return maxValue - minValue
}
