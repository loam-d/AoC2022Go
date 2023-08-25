package main

import (
	"strconv"
	"sync"
)

type Loc struct {
	x, y int
}

func day8processLines() [][]int {
	filescanner, readFile := getFilescanner("Inputs/day8.txt")
	trees := make([][]int, 0)
	for filescanner.Scan() {
		raw := filescanner.Text()
		treeRow := make([]int, 0)
		for _, c := range raw {
			num, _ := strconv.Atoi(string(c))
			treeRow = append(treeRow, num)
		}
		trees = append(trees, treeRow)
	}
	readFile.Close()

	return trees
}

func loc(x int, y int) Loc {
	return Loc{x: x, y: y}
}

func findVisible(trees [][]int) map[Loc]struct{} {
	visibleMap := make(map[Loc]struct{})
	for i := 0; i < 2; i++ {
		//from the left and right
		for j := 0; j < len(trees); j++ {
			maxvis := -1
			for k := 0; k < len(trees[0]); k++ {
				if trees[j][k] > maxvis {
					visibleMap[loc(j, k)] = struct{}{}
					maxvis = trees[j][k]
				}
			}
			maxvis2 := -1
			for k := len(trees[0]) - 1; k >= 0; k-- {
				if trees[j][k] > maxvis2 {
					visibleMap[loc(j, k)] = struct{}{}
					maxvis2 = trees[j][k]
				}
			}
		}
		// from the top and bottom
		for k := 0; k < len(trees[0]); k++ {
			maxvis := -1
			for j := 0; j < len(trees); j++ {
				if trees[j][k] > maxvis {
					visibleMap[loc(j, k)] = struct{}{}
					maxvis = trees[j][k]
				}
			}
			maxvis2 := -1
			for j := len(trees) - 1; j >= 0; j-- {

				if trees[j][k] > maxvis2 {
					visibleMap[loc(j, k)] = struct{}{}
					maxvis2 = trees[j][k]
				}
			}
		}
	}

	return visibleMap
}

func calculateScenicScore(location Loc, trees [][]int) int {
	starti := location.x
	startj := location.y
	treehouseHeight := trees[starti][startj]
	totals := make([]int, 0)

	totals = append(totals, 0)
	for i := starti + 1; i < len(trees); i++ {
		totals[len(totals)-1]++
		if trees[i][startj] >= treehouseHeight {
			break
		}
	}

	totals = append(totals, 0)
	for i := starti - 1; i >= 0; i-- {
		totals[len(totals)-1]++
		if trees[i][startj] >= treehouseHeight {
			break
		}
	}

	totals = append(totals, 0)
	for j := startj + 1; j < len(trees[0]); j++ {
		totals[len(totals)-1]++
		if trees[starti][j] >= treehouseHeight {
			break
		}
	}

	totals = append(totals, 0)
	for j := startj - 1; j >= 0; j-- {
		totals[len(totals)-1]++
		if trees[starti][j] >= treehouseHeight {
			break
		}
	}

	product := 1
	for _, in := range totals {
		product *= in
	}

	return product
}

func findMaxScenicScore(trees [][]int) int {
	maxScore := -1
	var wg sync.WaitGroup
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			wg.Add(1)
			// fix i and j in place so we can use them in parallel
			ig := i
			jg := j
			go func() {
				defer wg.Done()
				currLoc := loc(ig, jg)
				currScore := calculateScenicScore(currLoc, trees)
				if currScore > maxScore {
					maxScore = currScore
				}
			}()
		}
	}

	wg.Wait()
	return maxScore
}

func day8_1() (int, int) {
	trees := day8processLines()
	visibleMap := findVisible(trees)
	return len(visibleMap), findMaxScenicScore(trees)
}
