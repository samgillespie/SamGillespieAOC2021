package answers

import (
	"strconv"
	"strings"
)

func Day4() []int {
	data := ReadInputAsStr(4)
	return []int{
		q4part1(data),
		q4part2(data),
	}
}

type Board struct {
	id        int
	positions []*BoardPosition
	// Store References to rows and columns for fast retrieval
	rows       [][]*BoardPosition
	columns    [][]*BoardPosition
	lastCalled int
}

type BoardPosition struct {
	id       int
	x        int
	y        int
	value    int
	selected bool
}

func (b *Board) CallNumber(number int) {
	b.lastCalled = number
	for index, elem := range b.positions {
		if elem.value == number {
			elem.selected = true
			b.positions[index] = elem
			return
		}
	}
}

func (b Board) Score() int {
	score := 0
	for _, elem := range b.positions {
		if elem.selected == false {
			score += elem.value
		}
	}
	return score * b.lastCalled
}

func (b Board) isWinning() bool {
	for _, row := range b.rows {
		winning := true
		for _, elem := range row {
			if elem.selected == false {
				winning = false
				break
			}
		}
		if winning == true {
			return true
		}
	}

	for _, col := range b.columns {
		winning := true
		for _, elem := range col {
			if elem.selected == false {
				winning = false
				break
			}
		}
		if winning == true {
			return true
		}
	}
	return false
}

func ConvertRowsToBoard(boardData [][]int) Board {
	var positions []*BoardPosition
	rows := make([][]*BoardPosition, 5)
	columns := make([][]*BoardPosition, 5)
	for y, row := range boardData {
		for x, value := range row {
			position := BoardPosition{
				x:        x,
				y:        y,
				value:    value,
				id:       x*5 + y,
				selected: false,
			}
			positions = append(positions, &position)
			rows[y] = append(rows[y], &position)
			columns[x] = append(columns[x], &position)
		}
	}
	return Board{positions: positions, rows: rows, columns: columns}

}

func parseInput(input []string) ([]int, []*Board) {
	numbersCalled := []int{}
	row := strings.Split(input[0], ",")
	for _, num := range row {
		numInt, _ := strconv.Atoi(num)
		numbersCalled = append(numbersCalled, numInt)
	}

	boardChunk := [][]int{}
	boards := []*Board{}
	for index, row := range input {
		if index == 0 {
			continue
		}
		if row == "" {
			continue
		}

		rowSplit := strings.Split(row, " ")
		rowParsed := []int{}
		for _, rowElem := range rowSplit {
			if rowElem == "" {
				continue
			}
			elemParsed, _ := strconv.Atoi(rowElem)
			rowParsed = append(rowParsed, elemParsed)
		}
		boardChunk = append(boardChunk, rowParsed)

		if len(boardChunk) == 5 {
			newBoard := ConvertRowsToBoard(boardChunk)
			newBoard.id = len(boards)
			boards = append(boards, &newBoard)
			boardChunk = [][]int{}
		}
	}

	return numbersCalled, boards
}

func q4part1(data []string) int {
	numbersCalled, boards := parseInput(data)
	for roundNumber, numberCalled := range numbersCalled {
		for _, board := range boards {
			board.CallNumber(numberCalled)
			if roundNumber >= 5 && board.isWinning() {
				return board.Score()
			}
		}
	}
	return -1
}

func q4part2(data []string) int {
	numbersCalled, boards := parseInput(data)
	for _, numberCalled := range numbersCalled {
		stillPlaying := []*Board{}
		for _, board := range boards {
			board.CallNumber(numberCalled)
		}
		// Remove the losers
		for _, board := range boards {
			if board.isWinning() == false {
				stillPlaying = append(stillPlaying, board)

			} else if len(boards) == 1 {
				return boards[0].Score()
			}
		}
		boards = stillPlaying
	}
	return -1
}
