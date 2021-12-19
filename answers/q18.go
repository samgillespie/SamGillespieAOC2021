package answers

import (
	"fmt"
	"regexp"
	"strconv"
)

type SnailfishPair struct {
	leftElem   *SnailfishPair
	rightElem  *SnailfishPair
	parent     *SnailfishPair
	leftValue  int
	rightValue int
	depth      int
}

func (sp *SnailfishPair) AsString() string {
	str := "["
	if sp.leftValue != -1 {
		str += fmt.Sprintf("%d", sp.leftValue)
	} else {
		str += sp.leftElem.AsString()
	}
	str += ","

	if sp.rightValue != -1 {
		str += fmt.Sprintf("%d", sp.rightValue)
	} else {
		str += sp.rightElem.AsString()
	}
	str += "]"
	return str
}

func (sp *SnailfishPair) Calculate4Deep(depth int) *SnailfishPair {
	sp.depth = depth
	if sp.leftValue == -1 {
		res := sp.leftElem.Calculate4Deep(depth + 1)
		if res.depth == 4 {
			return res
		}
	}
	if sp.rightValue == -1 {
		res := sp.rightElem.Calculate4Deep(depth + 1)
		if res.depth == 4 {
			return res
		}
	}
	return sp
}

func (sp *SnailfishPair) RightMost() *SnailfishPair {
	if sp.rightValue == -1 {
		return sp.rightElem.RightMost()
	}
	return sp
}

func (sp *SnailfishPair) LeftMost() *SnailfishPair {
	if sp.leftValue == -1 {
		return sp.leftElem.LeftMost()
	}
	return sp
}

func (sp *SnailfishPair) Magnitude() int {
	value := 0
	if sp.leftValue == -1 {
		value += sp.leftElem.Magnitude() * 3
	} else {
		value += sp.leftValue * 3
	}

	if sp.rightValue == -1 {
		value += sp.rightElem.Magnitude() * 2
	} else {
		value += sp.rightValue * 2
	}
	return value
}

func (sp *SnailfishPair) Explode() {
	if sp.leftValue == -1 {
		sp.leftElem.Explode()
		return
	} else if sp.rightValue == -1 {
		sp.rightElem.Explode()
		return
	}
	// Do left hand side
	parent := sp.parent
	child := sp
	for {
		if parent.leftValue != -1 {
			parent.leftValue += sp.leftValue
			break
		}
		if parent.leftElem != child {
			parent.leftElem.RightMost().rightValue += sp.leftValue
			break
		}
		child = parent
		parent = parent.parent
		if parent == nil {
			break
		}
	}

	parent = sp.parent
	child = sp
	for {
		if parent.rightValue != -1 {
			parent.rightValue += sp.rightValue
			break
		}
		if parent.rightElem != child {
			if parent.rightElem.LeftMost().leftValue == -1 {
				panic("Whoops")
			}
			parent.rightElem.LeftMost().leftValue += sp.rightValue
			break
		}
		child = parent
		parent = parent.parent
		if parent == nil {
			break
		}
	}

	if sp.parent.leftElem == sp {
		sp.parent.leftValue = 0
		sp.parent.leftElem = nil
	} else if sp.parent.rightElem == sp {
		sp.parent.rightValue = 0
		sp.parent.rightElem = nil
	} else {
		panic("Shouldn't happen")
	}
}

func (sp *SnailfishPair) DoSplit() bool {
	if sp.leftValue > 9 {
		newPair := SnailfishPair{
			leftValue:  sp.leftValue / 2,
			rightValue: sp.leftValue - sp.leftValue/2,
			parent:     sp,
		}
		sp.leftElem = &newPair
		sp.leftValue = -1
		return true
	}

	if sp.rightValue > 9 {
		newPair := SnailfishPair{
			leftValue:  sp.rightValue / 2,
			rightValue: sp.rightValue - sp.rightValue/2,
			parent:     sp,
		}
		sp.rightElem = &newPair
		sp.rightValue = -1
		return true
	}
	return false
}

func (sp *SnailfishPair) FindAndDoSplit() bool {
	if sp.leftValue == -1 {
		if sp.leftElem.FindAndDoSplit() == true {
			return true
		}
	}
	if sp.DoSplit() == true {
		return true
	}
	if sp.rightValue == -1 {
		if sp.rightElem.FindAndDoSplit() == true {
			return true
		}
	}
	return false
}

func (sp *SnailfishPair) Add(newPair *SnailfishPair) *SnailfishPair {
	res := &SnailfishPair{
		leftElem:   sp,
		rightElem:  newPair,
		leftValue:  -1,
		rightValue: -1,
	}
	sp.parent = res
	newPair.parent = res
	return res
}

func (sp *SnailfishPair) DoAddition(newFish *SnailfishPair) *SnailfishPair {
	newRoot := sp.Add(newFish)

	for {
		explosion := false
		for {
			nextToExplode := newRoot.Calculate4Deep(0)
			if nextToExplode.depth == 0 {
				break
			}
			nextToExplode.Explode()
			explosion = true
		}
		split := newRoot.FindAndDoSplit()

		if explosion == false && split == false {
			break
		}
	}
	return newRoot
}

func Parse(pattern []byte) *SnailfishPair {
	chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	fishmap := map[byte]*SnailfishPair{}
	// Find a closed pair
	re := regexp.MustCompile(`\[[0-9,a-z],[0-9,a-z]\]`)
	charCursor := 0
	for len(pattern) > 1 {
		fishLetter := chars[charCursor]
		res := re.FindIndex(pattern)
		fish := SnailfishPair{leftValue: -1, rightValue: -1}

		// If we find a number, it's a value.  If we find a letter, it's a snailfishpair
		l, err := strconv.Atoi(string(pattern[res[0]+1]))
		if err != nil {
			fish.leftElem = fishmap[pattern[res[0]+1]]
			fishmap[pattern[res[0]+1]].parent = &fish
		} else {
			fish.leftValue = l
		}

		r, err := strconv.Atoi(string(pattern[res[0]+3]))
		if err != nil {
			fish.rightElem = fishmap[pattern[res[0]+3]]
			fishmap[pattern[res[0]+3]].parent = &fish
		} else {
			fish.rightValue = r
		}

		pattern[res[0]] = fishLetter
		pattern = append(pattern[0:res[0]+1], pattern[res[1]:]...)
		fishmap[fishLetter] = &fish
		charCursor++
	}
	return fishmap[chars[charCursor-1]]
}

func Day18() []int {
	// RunUnitTests()
	return []int{q18part1(), q18part2()}
}

func q18part1() int {
	data := ReadInputAsStr(0)
	newFish := Parse([]byte(data[0]))
	for _, row := range data[1:] {
		newFish = newFish.DoAddition(Parse([]byte(row)))
	}
	return newFish.Magnitude()
}

func q18part2() int {
	data := ReadInputAsStr(18)

	best := 0
	for x, xdata := range data {
		for y, ydata := range data {
			if x == y {
				continue
			}
			xfish := Parse([]byte(xdata))
			yfish := Parse([]byte(ydata))

			mag := xfish.DoAddition(yfish).Magnitude()
			if mag > best {
				best = mag

			}
		}
	}
	return best
}

// --------- UNIT TESTS -----------
func RunUnitTests() {
	test4Deep()
	testExplode()
	testMagnitude()
	testAddition()
	testMultipleAddition()
}

func test4Deep() {
	fish := Parse([]byte("[[[[9,8],1],2],3]"))
	res := fish.Calculate4Deep(0)
	if res.depth != 0 {
		panic("Oof")
	}
	fish = Parse([]byte("[[[[[9,8],1],2],3],4]"))
	res = fish.Calculate4Deep(0)
	if res.depth != 4 {
		panic("Bad at math")
	}
	fmt.Println("test4Deep Success")
}

func testExplode() {
	cases := [][]string{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for idx, testCase := range cases {
		input := testCase[0]
		solution := testCase[1]
		fish := Parse([]byte(input))
		target := fish.Calculate4Deep(0)
		target.Explode()
		result := fish.AsString()
		if solution != result {
			fmt.Printf("Error in case %d\n", idx)
			fmt.Printf("Input:     %s\n", input)
			fmt.Printf("Expected:  %s\n", solution)
			fmt.Printf("Produced:  %s\n", result)
			panic("Failed Test")
		}
	}

	fmt.Println("testExplode success")
}

func testAddition() {
	input1 := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	fish1 := Parse([]byte(input1))

	input2 := "[1,1]"
	fish2 := Parse([]byte(input2))
	result := fish1.DoAddition(fish2)
	solution := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	if solution != result.AsString() {
		fmt.Printf("Error in testAddition\n")
		fmt.Printf("Expected:  %s\n", solution)
		fmt.Printf("Produced:  %s\n", result.AsString())
		panic("Failed Test")
	}

	fish1 = Parse([]byte(input1))
	fish2 = Parse([]byte(input2))
	result = fish2.DoAddition(fish1)
	solution = "[[1,5],[[[0,7],4],[[7,8],[6,0]]]]"
	if solution != result.AsString() {
		fmt.Printf("Error in testAddition\n")
		fmt.Printf("Expected:  %s\n", solution)
		fmt.Printf("Produced:  %s\n", result.AsString())
		panic("Failed Test")
	}
	fmt.Println("testAddition success")
}

func testMultipleAddition() {
	elems := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}
	fish := Parse([]byte(elems[0]))
	for _, row := range elems[1:] {
		newfish := Parse([]byte(row))
		fish = fish.DoAddition(newfish)
	}
	solution := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	if fish.AsString() != solution {
		fmt.Printf("Error in testAddition\n")
		fmt.Printf("Expected:  %s\n", solution)
		fmt.Printf("Produced:  %s\n", fish.AsString())
		panic("Failed Test")
	}
	fmt.Println("testMultipleAddition success")
}

func testMagnitude() {
	type magnitude struct {
		inputString string
		solution    int
	}
	cases := []magnitude{
		{inputString: "[[1,2],[[3,4],5]]", solution: 143},
		{inputString: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", solution: 1384},
		{inputString: "[[[[1,1],[2,2]],[3,3]],[4,4]]", solution: 445},
		{inputString: "[[[[3,0],[5,3]],[4,4]],[5,5]]", solution: 791},
		{inputString: "[[[[5,0],[7,4]],[5,5]],[6,6]]", solution: 1137},
		{inputString: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", solution: 3488},
	}

	for idx, testCase := range cases {
		input := testCase.inputString
		solution := testCase.solution
		fish := Parse([]byte(input))
		if solution != fish.Magnitude() {
			fmt.Printf("Error in case %d\n", idx)
			fmt.Printf("Input:     %s\n", input)
			fmt.Printf("Expected:  %d\n", solution)
			fmt.Printf("Produced:  %d\n", fish.Magnitude())
			panic("Failed Test")
		}
	}
	fmt.Println("testMagnitude Success")
}
