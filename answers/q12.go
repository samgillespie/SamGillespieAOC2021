package answers

import (
	"fmt"
	"strings"
)

type CaveSection struct {
	Name    string
	Exits   []*CaveSection
	BigCave bool
}

func Day12() []int {
	data := ReadInputAsStr(12)
	caveMap := make(map[string]CaveSection)
	for _, row := range data {
		rowSplit := strings.Split(row, "-")
		start, isfound := caveMap[rowSplit[0]]
		if isfound == false {

			caveMap[rowSplit[0]] = CaveSection{
				Name:    rowSplit[0],
				Exits:   make([]*CaveSection, 0),
				BigCave: rowSplit[0] == strings.ToUpper(rowSplit[0]),
			}
			start = caveMap[rowSplit[0]]
		}
		end, endfound := caveMap[rowSplit[1]]
		if endfound == false {
			caveMap[rowSplit[1]] = CaveSection{
				Name:    rowSplit[1],
				Exits:   make([]*CaveSection, 0),
				BigCave: rowSplit[1] == strings.ToUpper(rowSplit[1]),
			}
			end = caveMap[rowSplit[1]]
		}

		start.Exits = append(start.Exits, &end)
		end.Exits = append(end.Exits, &start)
		caveMap[rowSplit[0]] = start
		caveMap[rowSplit[1]] = end
	}
	return []int{q12part1(caveMap), q12part2(caveMap)}
}

func canTravel(path []CaveSection, newSection CaveSection) bool {
	// Returns True if the pathway is valid
	for _, elem := range path {
		if elem.BigCave == true {
			continue
		}
		if elem.Name == newSection.Name {
			return false
		}
	}
	return true
}

func canTravelP2(path []CaveSection, newSection CaveSection) bool {
	// Returns True if the pathway is valid
	visitedTimes := map[string]int{}
	visitedSingleCaveTwice := false
	for _, elem := range path {
		// Don't track Big Caves
		if elem.BigCave == true {
			continue
		}
		// Can travel to the same place
		_, visited := visitedTimes[elem.Name]
		if visited == false {
			visitedTimes[elem.Name] = 1
		} else {
			visitedTimes[elem.Name] += 1
			visitedSingleCaveTwice = true
		}
	}

	numVisits, visitedNewSection := visitedTimes[newSection.Name]
	if visitedNewSection == false {
		return true
	} else if numVisits == 1 && visitedSingleCaveTwice == true {
		return false
	} else if numVisits == 1 && visitedSingleCaveTwice == false {
		return true
	} else if numVisits == 2 {
		return false
	}
	panic("Shouldn't get here")
	return true
}

func PrintPaths(paths [][]CaveSection) {
	for _, path := range paths {
		PrintPath(path)
	}
}

func PrintPath(path []CaveSection) {
	pathstr := ""
	for _, pathElem := range path {
		pathstr += pathElem.Name + "-"
	}
	fmt.Println(pathstr)
}

func q12part1(caveMap map[string]CaveSection) int {
	start := caveMap["start"]
	activePaths := [][]CaveSection{[]CaveSection{start}}
	finishedPaths := [][]CaveSection{}
	for len(activePaths) > 0 {
		// Remove the first elementts
		path := activePaths[0]

		if len(activePaths) == 1 {
			activePaths = [][]CaveSection{}
		} else {
			activePaths = activePaths[1:]
		}
		lastPosition := path[len(path)-1]
		if lastPosition.Name == "end" {
			finishedPaths = append(finishedPaths, path)
			continue
		}

		// Add all the active paths
		pathsFound := 0
		for _, exit := range lastPosition.Exits {
			// Lookup in map.
			cave := caveMap[exit.Name]
			if canTravel(path, cave) == false {
				continue
			}
			newPath := append(path, cave)

			// Copy because golang jank
			newPathCopy := make([]CaveSection, len(newPath))
			copy(newPathCopy, newPath)
			activePaths = append(activePaths, newPathCopy)
			pathsFound++
		}

		// fmt.Println("V ACTIVE V")
		// PrintPaths(activePaths)
		// fmt.Println("V FINISHED V")
		// PrintPaths(finishedPaths)
		// fmt.Println("---- END CYCLE ---- ")
	}
	// PrintPaths(finishedPaths)
	return len(finishedPaths)
}

func q12part2(caveMap map[string]CaveSection) int {
	start := caveMap["start"]
	activePaths := [][]CaveSection{[]CaveSection{start}}
	finishedPaths := [][]CaveSection{}
	for len(activePaths) > 0 {
		// Remove the first elementts
		path := activePaths[0]

		if len(activePaths) == 1 {
			activePaths = [][]CaveSection{}
		} else {
			activePaths = activePaths[1:]
		}
		lastPosition := path[len(path)-1]
		if lastPosition.Name == "end" {
			finishedPaths = append(finishedPaths, path)
			continue
		}

		// Add all the active paths
		pathsFound := 0
		for _, exit := range lastPosition.Exits {
			// Lookup in map.
			if exit.Name == "start" {
				continue
			}
			cave := caveMap[exit.Name]
			if canTravelP2(path, cave) == false {
				continue
			}
			newPath := append(path, cave)

			// Copy because golang jank
			newPathCopy := make([]CaveSection, len(newPath))
			copy(newPathCopy, newPath)
			activePaths = append(activePaths, newPathCopy)
			pathsFound++
		}
	}
	// PrintPaths(finishedPaths)
	return len(finishedPaths)
}
