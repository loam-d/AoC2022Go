package main

import (
	"strconv"
	"strings"
)

func containsCompleteOverlap(pair1 []int, pair2 []int) bool {
	contains := false

	contains = pair1[0] >= pair2[0] &&
		pair1[0] <= pair2[1] &&
		pair1[1] >= pair2[0] &&
		pair1[1] <= pair2[1]
	contains = contains || (pair2[0] >= pair1[0] &&
		pair2[0] <= pair1[1] &&
		pair2[1] >= pair1[0] &&
		pair2[1] <= pair1[1])
	return contains
}

func containsPartialOverlap(pair1 []int, pair2 []int) bool {
	contains := false

	contains = (pair1[0] >= pair2[0] &&
		pair1[0] <= pair2[1]) ||
		(pair1[1] >= pair2[0] &&
			pair1[1] <= pair2[1])
	contains = contains || ((pair2[0] >= pair1[0] &&
		pair2[0] <= pair1[1]) ||
		(pair2[1] >= pair1[0] &&
			pair2[1] <= pair1[1]))
	return contains
}

func day4processLine(line string, containsFn func([]int, []int) bool) int {
	pairs := strings.Split(line, ",")
	rawpair1 := strings.Split(pairs[0], "-")
	pair1 := make([]int, 0)
	for _, s := range rawpair1 {
		ints, _ := strconv.Atoi(s)
		pair1 = append(pair1, ints)
	}
	rawpair2 := strings.Split(pairs[1], "-")
	pair2 := make([]int, 0)
	for _, s := range rawpair2 {
		ints, _ := strconv.Atoi(s)
		pair2 = append(pair2, ints)
	}

	contains := containsFn(pair1, pair2)

	if contains {
		return 1
	} else {
		return 0
	}
}

func day4_1() int {
	filescanner, readFile := getFilescanner("Inputs/day4.txt")
	numOverlapPairs := 0
	for filescanner.Scan() {
		numOverlapPairs += day4processLine(filescanner.Text(), containsCompleteOverlap)
	}

	readFile.Close()
	return numOverlapPairs
}

func day4_2() int {
	filescanner, readFile := getFilescanner("Inputs/day4.txt")
	numOverlapPairs := 0
	for filescanner.Scan() {
		numOverlapPairs += day4processLine(filescanner.Text(), containsPartialOverlap)
	}

	readFile.Close()
	return numOverlapPairs
}
