package main

import (
	"strconv"
	"strings"
)

func addCoordToRockMap(rockMap map[int][]string, toAdd Loc) {
	if _, ok := rockMap[toAdd.x]; !ok {
		rockMap[toAdd.x] = make([]string, 0)
	}
	for i := 0; i < toAdd.y; i++ {
		if i == len(rockMap[toAdd.x]) {
			rockMap[toAdd.x] = append(rockMap[toAdd.x], ".")
		}
		if rockMap[toAdd.x][i] != "#" {
			rockMap[toAdd.x][i] = "."
		}
	}

	if toAdd.y < len(rockMap[toAdd.x]) {
		rockMap[toAdd.x][toAdd.y] = "#"
	} else {
		rockMap[toAdd.x] = append(rockMap[toAdd.x], "#")
	}
}

func addFloor(rockMap map[int][]string, k int, floorDepth int) {
	for i := len(rockMap[k]); i <= floorDepth; i++ {
		toAppend := "."
		if i == floorDepth {
			toAppend = "#"
		}
		rockMap[k] = append(rockMap[k], toAppend)
	}
}

func addFloors(rockMap map[int][]string, floorDepth int) {
	for k := range rockMap {
		addFloor(rockMap, k, floorDepth)
	}
}

func processRockMap(filename string) (map[int][]string, int) {
	filescanner, readFile := getFilescanner(filename)
	defer readFile.Close()

	rockMap := make(map[int][]string)
	paths := make([][]Loc, 0)

	for filescanner.Scan() {
		line := filescanner.Text()
		rawCoords := strings.Split(line, " -> ")
		paths = append(paths, make([]Loc, 0))
		for _, rawCoord := range rawCoords {
			nums := strings.Split(rawCoord, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			rockLoc := loc(x, y)
			paths[len(paths)-1] = append(paths[len(paths)-1], rockLoc)
		}
	}

	floorDepth := 0
	for _, rockCoords := range paths {
		if len(rockCoords) == 1 {
			addCoordToRockMap(rockMap, rockCoords[0])
			continue
		}
		for i := 1; i < len(rockCoords); i++ {
			coord1 := rockCoords[i-1]
			coord2 := rockCoords[i]
			if coord1.y > floorDepth {
				floorDepth = coord1.y
			}
			if coord2.y > floorDepth {
				floorDepth = coord2.y
			}
			iter := 1
			if coord1.x == coord2.x {
				if coord2.y < coord1.y {
					iter = -iter
				}
				for j := coord1.y; j != coord2.y; j += iter {
					addCoordToRockMap(rockMap, loc(coord1.x, j))
				}
			} else {
				if coord2.x < coord1.x {
					iter = -iter
				}
				for j := coord1.x; j != coord2.x; j += iter {
					addCoordToRockMap(rockMap, loc(j, coord1.y))
				}
			}
			addCoordToRockMap(rockMap, coord2)
		}
	}

	// Do the floor for part 2
	addFloors(rockMap, floorDepth+2)
	return rockMap, floorDepth + 2
}

func processSand(rockMap map[int][]string, source Loc, floorDepth int) bool {
	// Returns true if the sand is at rest and false if it flows into the abyss
	j := source.x
	if _, ok := rockMap[j]; ok {
		if rockMap[j][source.y] == "o" {
			return false
		}
	}

OUTER:
	for i := source.y; ; i++ {
		if rockMap[j][i] == "." {
			continue
		}
		if _, ok := rockMap[j]; !ok {
			// No map at j means we're falling forever
			return false
		}
		if i >= len(rockMap[j]) {
			// exceeding the length of a map means there's no rocks to hit and we fall forever
			return false
		}
		if rockMap[j][i] == "o" || rockMap[j][i] == "#" {
			// First check if we can fall left or right
			for k := -1; k < 2; k = k + 2 {
				if _, ok := rockMap[j+k]; ok {
					if i >= len(rockMap[j+k]) {
						// We're going into an undefined part of the map which means we fall forever
						// should never happen in part 2
						return false
					}
					if !(rockMap[j+k][i] == "o" || rockMap[j+k][i] == "#") {
						j = j + k
						continue OUTER
					}
				} else {
					if floorDepth <= 0 {
						return false
					}
					rockMap[j+k] = make([]string, 0)
					addFloor(rockMap, j+k, floorDepth)
					k = k - 2
				}
			}
			// Otherwise stack up
			rockMap[j][i-1] = "o"
			return true
		}
	}
}

func processAllSand(rockMap map[int][]string, floorDepth int) int {
	source := loc(500, 0)
	numSettled := 0
	for processSand(rockMap, source, floorDepth) {
		numSettled++
	}
	return numSettled
}

func day14_1() int {
	rockMap, floorDepth := processRockMap("Inputs/day14.txt")
	return processAllSand(rockMap, floorDepth)
}
