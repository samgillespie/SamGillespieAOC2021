package answers

import "fmt"

func Day21() []int {
	p1Start := 7
	p2Start := 9
	return []int{q21part1(p1Start, p2Start), q21part2(p1Start, p2Start)}
}

func q21part1(p1Pos int, p2Pos int) int {
	var p1Score, p2Score, rounds int

	currdice := 1
	isPlayer1 := true

	for p1Score < 1000 && p2Score < 1000 {
		rounds++
		roll := currdice*3 + 3
		currdice = currdice + 3
		if currdice > 100 {
			currdice = 1 + (currdice - 101)
		}

		if isPlayer1 == true {
			p1Pos = (p1Pos + roll) % 10
			if p1Pos == 0 {
				p1Pos = 10
			}
			p1Score += p1Pos
			isPlayer1 = false
			// fmt.Printf("Player 1 rolls %d and move to space %d for a total score of %d\n", roll, p1Pos, p1Score)
		} else {
			p2Pos = (p2Pos + roll) % 10
			if p2Pos == 0 {
				p2Pos = 10
			}
			isPlayer1 = true
			p2Score += p2Pos
			// fmt.Printf("Player 2 rolls %d and move to space %d for a total score of %d\n", roll, p2Pos, p2Score)
		}
	}
	minScore := Min(p1Score, p2Score)
	return minScore * rounds * 3
}

type State struct {
	p1Pos   int
	p2Pos   int
	p1Score int
	p2Score int
}

func ClipPos(x int) int {
	if x > 10 {
		return x - 10
	}
	return x
}

func ClipScore(x int) int {
	if x > 21 {
		return 21
	}
	return x
}

func (s State) StepForward(isPlayer1 bool) map[State]int {
	// Simulate
	states := map[State]int{}
	diceDistributions := map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	if isPlayer1 == true {
		for spaces, universes := range diceDistributions {
			newPos := ClipPos(s.p1Pos + spaces)
			newState := State{
				p1Pos:   newPos,
				p2Pos:   s.p2Pos,
				p1Score: ClipScore(s.p1Score + newPos),
				p2Score: s.p2Score,
			}
			states[newState] += universes
		}
	} else {
		for spaces, universes := range diceDistributions {
			newPos := ClipPos(s.p2Pos + spaces)
			newState := State{
				p1Pos:   s.p1Pos,
				p2Pos:   newPos,
				p1Score: s.p1Score,
				p2Score: ClipScore(s.p2Score + newPos),
			}
			states[newState] += universes
		}
	}
	return states
}

func q21part2(p1Pos int, p2Pos int) int {
	states := map[State]int{}
	initialState := State{p1Pos: p1Pos, p2Pos: p2Pos, p1Score: 0, p2Score: 0}
	states[initialState] = 1
	finished := false
	rounds := 0
	isPlayer1 := true
	for finished == false {
		rounds++
		finished = true
		newStates := map[State]int{}
		for state, universes := range states {
			if state.p1Score == 21 || state.p2Score == 21 {
				newStates[state] += universes
				continue
			}
			finished = false
			steppedStates := state.StepForward(isPlayer1)
			for state, newUniverses := range steppedStates {
				newStates[state] += universes * newUniverses
			}
		}
		isPlayer1 = !isPlayer1
		states = newStates
		universes := 0
		for _, j := range states {
			universes += j
		}
	}

	var p1Wins, p2Wins int
	for state, universes := range states {
		if state.p1Score == 21 {
			p1Wins += universes
		} else if state.p2Score == 21 {
			p2Wins += universes
		} else {
			panic("Something went wrong")
		}
	}
	fmt.Println(p1Wins, p2Wins)
	return Max(p1Wins, p2Wins)

}
