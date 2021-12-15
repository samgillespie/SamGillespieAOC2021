package answers

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	x          int
	y          int
	n          int
	value      int
	neighbours []*Node
}

type Nodes struct {
	xmax  int
	ymax  int
	cells int
	nodes []Node
}

func (n Nodes) GetXY(x int, y int) Node {
	return n.nodes[n.ymax*y+x]
}

func (n Nodes) GetN(idx int) Node {
	return n.nodes[idx]
}

func (n Nodes) Print() {
	str := ""
	for idx, node := range n.nodes {
		if idx%n.xmax == 0 {
			fmt.Println(str)
			str = ""
		}
		str += fmt.Sprintf("%d", node.value)
	}
	fmt.Println(str)
}

type NodeQueue []Node

func (n NodeQueue) Pop(idx int) (NodeQueue, Node) {
	node := n[idx]
	n = append(n[0:idx], n[idx+1:]...)
	return n, node
}

func (n NodeQueue) LowestFScore(fscores map[int]int) (NodeQueue, Node) {
	minscore := 9999999999999
	pos := -1
	for idx, node := range n {
		if fscores[node.n] < minscore {

			minscore = fscores[node.n]
			pos = idx
		}
	}
	return n.Pop(pos)
}

func (n NodeQueue) Add(node Node) NodeQueue {
	n = append(n, node)
	return n
}

func ReadInputP1() Nodes {
	data := ReadInputAsStr(15)
	nodes := Nodes{ymax: len(data), xmax: len(data[0]), cells: (len(data) + 1) * len(data[0])}
	n := 0
	for y, row := range data {
		for x, cell := range row {
			value, _ := strconv.Atoi(string(cell))
			nodes.nodes = append(nodes.nodes, Node{
				x:     x,
				y:     y,
				n:     n,
				value: value,
			})
			n++
		}
	}
	for index, currnode := range nodes.nodes {
		// left
		if currnode.x > 0 {
			leftNode := nodes.GetXY(currnode.x-1, currnode.y)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &leftNode)
		}
		// right
		if currnode.x < nodes.xmax-1 {
			rightNode := nodes.GetXY(currnode.x+1, currnode.y)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &rightNode)
		}
		// up
		if currnode.y > 0 {
			upNode := nodes.GetXY(currnode.x, currnode.y-1)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &upNode)
		}
		// down
		if currnode.y < nodes.ymax-1 {
			downNode := nodes.GetXY(currnode.x, currnode.y+1)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &downNode)
		}
	}
	return nodes
}

func ReadInputP2() Nodes {
	data := ReadInputAsStr(15)
	nodes := Nodes{ymax: len(data) * 5, xmax: len(data[0]) * 5, cells: (len(data) + 1) * len(data[0]) * 25}

	n := 0
	for y := 0; y < (5 * len(data)); y++ {
		for x := 0; x < 5*len(data[0]); x++ {
			// Get the cell
			tempx := x % len(data[0])
			tempy := y % len(data)
			cellx := int(x / len(data[0]))
			celly := int(y / len(data))
			value, _ := strconv.Atoi(string(data[tempy][tempx]))

			value = value + cellx + celly
			if value >= 10 {
				value = value - 9
			}
			nodes.nodes = append(nodes.nodes, Node{
				x:     x,
				y:     y,
				n:     n,
				value: value,
			})
			n++
		}
	}

	for index, currnode := range nodes.nodes {
		// left
		if currnode.x > 0 {
			leftNode := nodes.GetXY(currnode.x-1, currnode.y)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &leftNode)
		}
		// right
		if currnode.x < nodes.xmax-1 {
			rightNode := nodes.GetXY(currnode.x+1, currnode.y)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &rightNode)
		}
		// up
		if currnode.y > 0 {
			upNode := nodes.GetXY(currnode.x, currnode.y-1)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &upNode)
		}
		// down
		if currnode.y < nodes.ymax-1 {
			downNode := nodes.GetXY(currnode.x, currnode.y+1)
			nodes.nodes[index].neighbours = append(nodes.nodes[index].neighbours, &downNode)
		}
	}
	return nodes
}

func Day15() []int {
	nodesP1 := ReadInputP1()
	nodesP2 := ReadInputP2()
	return []int{
		PerformAStar(nodesP1),
		PerformAStar(nodesP2),
	}
}

func PrintFscore(fscores map[int]int, xmax int, ymax int) {
	for y := 0; y < ymax; y++ {
		row := ""
		for x := 0; x < xmax; x++ {
			fscore := fmt.Sprintf("%d", fscores[y*ymax+x])
			spaces := 5 - len(fscore)

			row += fscore + strings.Repeat(" ", spaces)
		}
		fmt.Println(row)
	}

}

func h(node *Node, nodes Nodes) int {
	return (nodes.xmax-node.x)*3 + (nodes.ymax-node.y)*3
}

func PerformAStar(nodes Nodes) int {
	start := nodes.GetXY(0, 0)
	end := nodes.GetXY(nodes.xmax-1, nodes.ymax-1)
	openSet := NodeQueue{}
	openSet = append(openSet, start)

	cameFrom := map[int]int{}

	fScore := map[int]int{}
	gScore := map[int]int{}
	for i := 0; i < nodes.cells; i++ {
		fScore[i] = 9999
		gScore[i] = 9999
	}
	fScore[0] = h(&start, nodes)
	gScore[0] = 0

	iteration := 0
	for len(openSet) > 0 {
		// Pop from the front of the list
		var current Node
		openSet, current = openSet.LowestFScore(fScore)
		current = nodes.nodes[current.n]

		iteration++
		if current.x == end.x && current.y == end.y {
			//do thing
			// fmt.Println(" ---- FScore ---- ")
			// PrintFscore(fScore, nodes.xmax, nodes.ymax)
			// fmt.Println(" ---- GScore ---- ")
			// PrintFscore(gScore, nodes.xmax, nodes.ymax)
			return gScore[end.n]
		}

		for _, neighbour := range current.neighbours {
			tentative_gScore := gScore[current.n] + neighbour.value
			if tentative_gScore < gScore[neighbour.n] {
				cameFrom[neighbour.n] = current.n
				gScore[neighbour.n] = tentative_gScore
				fScore[neighbour.n] = tentative_gScore + h(neighbour, nodes)
				openSet = openSet.Add(*neighbour)
			}
		}
	}
	return -1
}

func q15part2(nodes Nodes) int {
	return -1
}
