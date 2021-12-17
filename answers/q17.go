package answers

type TargetArea struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func (ta TargetArea) IsWithin(x int, y int) bool {
	if x >= ta.xmin && x <= ta.xmax {
		if y >= ta.ymin && y <= ta.ymax {
			return true
		}
	}
	return false
}

func (ta TargetArea) Exceeded(x int, y int) bool {
	if x > ta.xmax {
		return true
	}
	if y < ta.ymin {
		return true
	}
	return false
}

func (ta TargetArea) ShootProbe(vx int, vy int) int {
	posx, posy, maxy := 0, 0, 0
	for {
		posx += vx
		posy += vy
		if posy > maxy {
			maxy = posy
		}

		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		vy--
		if ta.IsWithin(posx, posy) {
			return maxy
		}
		if ta.Exceeded(posx, posy) {
			return -1
		}
	}
}

func SampleInput() TargetArea {
	return TargetArea{
		xmin: 20,
		xmax: 30,
		ymin: -10,
		ymax: -5,
	}
}
func MyInput() TargetArea {
	return TargetArea{
		xmin: 81,
		xmax: 129,
		ymin: -150,
		ymax: -108,
	}
}

func Day17() []int {
	targetArea := MyInput()
	return []int{
		q17part1(targetArea),
		q17part2(targetArea),
	}
}

func q17part1(targetArea TargetArea) int {
	maxy := 0
	for x := 1; x < targetArea.xmax; x++ {
		for y := 1; y < -2*targetArea.ymax; y++ {
			ycurr := targetArea.ShootProbe(x, y)
			if ycurr > maxy {
				maxy = ycurr
			}
		}
	}
	return maxy
}

func q17part2(targetArea TargetArea) int {
	sol := [][]int{}
	for x := 0; x < targetArea.xmax+1; x++ {
		for y := targetArea.ymin; y < -2*targetArea.ymax; y++ {
			ycurr := targetArea.ShootProbe(x, y)
			if ycurr >= 0 {
				sol = append(sol, []int{x, y})
			}
		}
	}
	return len(sol)
}
