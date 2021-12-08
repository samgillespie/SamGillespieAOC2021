package answers

import (
	"fmt"
	"strconv"
	"strings"
)

type Catalog struct {
	// Known At Read Time
	One        []byte
	Four       []byte
	Seven      []byte
	Eight      []byte
	FiveDigits [][]byte // 2, 3, 5
	SixDigits  [][]byte // 0, 6, 9

	// Calculated
	Zero  []byte
	Two   []byte
	Three []byte
	Five  []byte
	Six   []byte
	Nine  []byte
}

type Digit struct {
	Top         byte
	LeftTop     byte
	RightTop    byte
	Middle      byte
	LeftBottom  byte
	RightBottom byte
	Bottom      byte
}

func (d Digit) CalculateDisplay(display []byte) int {
	if len(display) == 2 {
		return 1
	} else if len(display) == 3 {
		return 7
	} else if len(display) == 4 {
		return 4
	} else if len(display) == 7 {
		return 8
	}

	if len(display) == 5 {
		if BytesContains(display, d.LeftTop) == true {
			return 5
		} else if BytesContains(display, d.LeftBottom) == true {
			return 2
		} else {
			return 3
		}
	}

	if len(display) == 6 {
		if BytesContains(display, d.Middle) == false {
			return 0
		} else if BytesContains(display, d.RightTop) == false {
			return 6
		} else {
			return 9
		}
	}
	panic("something bad has happened")
}

func BytesContains(input []byte, b byte) bool {
	for _, char := range input {
		if char == b {
			return true
		}
	}
	return false
}

func Day8() []int {
	data := ReadInputAsStr(8)

	catalog := [][]string{}
	display := [][]string{}
	for _, row := range data {
		values := strings.Split(row, " | ")
		catalog = append(catalog, strings.Split(values[0], " "))
		display = append(display, strings.Split(values[1], " "))
	}
	return []int{q8part1(catalog, display), q8part2(catalog, display)}

}

func q8part1(catalog [][]string, display [][]string) int {
	counter := 0
	for _, row := range display {
		for _, elem := range row {
			if len(elem) != 5 && len(elem) != 6 {
				counter++
			}
		}
	}
	return counter
}

func stringDiff(a []byte, b []byte) []byte {
	// returns all the strings that are in a, but not b
	// Direction matters.  Put the bigger one in a
	diff := []byte{}
	// Are there any characters in a that aren't in b
	for _, runeA := range a {
		isDiff := true
		for _, runeB := range b {
			if runeA == runeB {
				isDiff = false
				break
			}
		}
		if isDiff == true {
			diff = append(diff, runeA)
		}
	}
	return diff
}

func q8part2(catalog [][]string, display [][]string) int {
	total := 0
	for index, entry := range catalog {
		// Convert to a much easier to work with format
		cat := Catalog{}
		for _, row := range entry {
			if len(row) == 2 {
				cat.One = []byte(row)
			} else if len(row) == 3 {
				cat.Seven = []byte(row)
			} else if len(row) == 4 {
				cat.Four = []byte(row)
			} else if len(row) == 5 {
				cat.FiveDigits = append(cat.FiveDigits, []byte(row))
			} else if len(row) == 6 {
				cat.SixDigits = append(cat.SixDigits, []byte(row))
			} else {
				cat.Eight = []byte(row)
			}
		}

		digit := Digit{}
		diff := stringDiff(cat.Seven, cat.One)
		if len(diff) != 1 {
			panic("something went wrong")
		}
		// Easy one
		digit.Top = diff[0]

		// Returns TopLeft and Middle
		diff = stringDiff(cat.Four, cat.One)

		// The 6's are missing middle, Bottom Left and Top Right
		// 6ers Diff (1 Diff 4) will return middle.  The other value of 1 Diff 4 will be Top Left
		for _, sixer := range cat.SixDigits {
			sixer_compare := stringDiff(diff, sixer)
			if len(sixer_compare) == 1 {
				if diff[0] == sixer_compare[0] {
					digit.Middle = diff[0]
					digit.LeftTop = diff[1]
				} else {
					digit.Middle = diff[1]
					digit.LeftTop = diff[0]
				}
				cat.Zero = sixer
				break
			}
		}
		known := []byte{digit.Middle, digit.LeftTop, digit.Top}

		// If we compare our knowns to the fivers, 2 and 3 have 2 of our shapes, and 5 has all three
		for _, fiver := range cat.FiveDigits {
			fiver_compare := stringDiff(known, fiver)
			if len(fiver_compare) == 0 {
				cat.Five = fiver
			}
		}
		diff1 := stringDiff(cat.Five, known) // RB, B
		diff2 := stringDiff(cat.One, diff1)  // RT
		digit.RightTop = diff2[0]

		diff3 := stringDiff(diff1, cat.One) // B
		digit.Bottom = diff3[0]

		diff4 := stringDiff(diff1, []byte{digit.Bottom})
		digit.RightBottom = diff4[0]

		diff = stringDiff(cat.Eight, []byte{digit.Middle, digit.LeftTop, digit.Top, digit.RightTop, digit.RightBottom, digit.Bottom})

		digit.LeftBottom = diff[0]
		result := ""
		for _, char := range display[index] {
			result += fmt.Sprintf("%d", digit.CalculateDisplay([]byte(char)))
		}
		resultInt, _ := strconv.Atoi(result)
		total += resultInt
	}
	return total
}
