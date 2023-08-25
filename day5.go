package main

import (
	"strconv"
	"strings"
)

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func filterSlice(s []rune, filteredRunes []string) []rune {
	filteredCol := make([]rune, 0)
	for _, c2 := range s {
		willContinue := false
		for _, c3 := range filteredRunes {
			if string(c2) == c3 {
				willContinue = true
				break
			}
		}
		if willContinue {
			continue
		}
		filteredCol = append(filteredCol, c2)
	}
	return filteredCol
}

func prepStacks(numStacks int) [][]string {
	stacks := make([][]string, 0)
	for i := 0; i < numStacks; i++ {
		stacks = append(stacks, make([]string, 0))
	}
	return stacks
}

func parseMove(move string) []int {
	moveSlice := strings.Split(move, " ")
	moveInts := make([]int, 0)
	amt, _ := strconv.Atoi(moveSlice[1])
	from, _ := strconv.Atoi(moveSlice[3])
	to, _ := strconv.Atoi(moveSlice[5])
	moveInts = append(moveInts, amt)
	moveInts = append(moveInts, from-1)
	moveInts = append(moveInts, to-1)
	return moveInts
}

func executeMove(move string, stacks [][]string) [][]string {
	moveInts := parseMove(move)
	for i := 0; i < moveInts[0]; i++ {
		from := moveInts[1]
		to := moveInts[2]
		var val string
		val, stacks[from] = stacks[from][len(stacks[from])-1], stacks[from][:len(stacks[from])-1]
		stacks[to] = append(stacks[to], val)
	}
	return stacks
}

func executeMoveP2(move string, stacks [][]string) [][]string {
	moveInts := parseMove(move)
	amt := moveInts[0]
	from := moveInts[1]
	to := moveInts[2]
	var val []string
	val, stacks[from] = stacks[from][len(stacks[from])-amt:len(stacks[from])], stacks[from][:len(stacks[from])-amt]
	for _, v := range val {
		stacks[to] = append(stacks[to], v)
	}
	return stacks
}

func loadStacks(stacks [][]string) [][]string {
	filescanner, readFile := getFilescanner("Inputs/day5-1.txt")
	windowSize := 3
	for filescanner.Scan() {
		line := filescanner.Text()
		line = line + "   "
		currCol := make([]rune, 0)
		filteredChars := []string{" ", "[", "]"}
		for i, c := range line {
			if i%(windowSize+1) == 0 {
				colVal := string(filterSlice(currCol, filteredChars))
				if colVal != "" {
					idx := (i / (windowSize + 1))
					stacks[idx-1] = append(stacks[idx-1], colVal)
				}
				currCol = make([]rune, 0)
			} else {
				currCol = append(currCol, c)
			}
		}
	}

	for i, s := range stacks {
		stacks[i] = stacks[i][:len(stacks[i])-1]
		stacks[i] = reverseSlice(s)
	}

	readFile.Close()
	return stacks
}

func getTopOfStacks(stacks [][]string) []string {
	tops := make([]string, 0)
	for _, stack := range stacks {
		tops = append(tops, stack[len(stack)-1])
	}
	return tops
}

func day5_1() string {
	stacks := prepStacks(9)
	stacks = loadStacks(stacks)
	filescanner, readFile := getFilescanner("Inputs/day5-2.txt")
	for filescanner.Scan() {
		move := filescanner.Text()
		stacks = executeMove(move, stacks)
	}
	readFile.Close()
	return strings.Join(getTopOfStacks(stacks), "")
}

func day5_2() string {
	stacks := prepStacks(9)
	stacks = loadStacks(stacks)
	filescanner, readFile := getFilescanner("Inputs/day5-2.txt")
	for filescanner.Scan() {
		move := filescanner.Text()
		stacks = executeMoveP2(move, stacks)
	}
	readFile.Close()
	return strings.Join(getTopOfStacks(stacks), "")
}
