package answers

import (
	"fmt"
	"sort"
)

func Day10() []int {
	data := ReadInputAsStr(10)
	// get incomplete from part 1 and feed to part 2
	incomplete, p1answer := q10part1(data)
	return []int{p1answer, q10part2(incomplete)}

}

func inSlice(str rune, slice []rune) bool {
	for _, elem := range slice {
		if elem == str {
			return true
		}
	}
	return false
}

func q10part1(data []string) ([]string, int) {
	openChars := []rune{'{', '[', '<', '('}
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	score := 0
	incomplete := []string{}
	for _, elem := range data {
		stack := []rune{}
		corrupted := false
		for _, char := range elem {
			if inSlice(char, openChars) {
				stack = append(stack, char)
				continue
			}
			// Handle the correct cases
			if char == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if char == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if char == '>' && stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1]
			} else if char == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				// Error
				// fmt.Printf("Expected %s got %s\n", string(stack[len(stack)-1]), string(char))
				score += points[char]
				corrupted = true
				break
			}
		}
		if corrupted == false {
			incomplete = append(incomplete, elem)
		}
	}
	return incomplete, score
}

func q10part2(incomplete []string) int {
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	openChars := []rune{'{', '[', '<', '('}
	scores := []int{}
	for _, elem := range incomplete {
		stack := []rune{}
		for _, char := range elem {
			if inSlice(char, openChars) {
				stack = append(stack, char)
				continue
			}
			// Handle the correct cases
			if char == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if char == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if char == '>' && stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1]
			} else if char == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				fmt.Println("Looks like we have a corrupted element in the incomplete list?!")
			}
		}

		// Resolve the stack in reverse order
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score * 5
			score += points[stack[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}
