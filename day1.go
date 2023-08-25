package main

import (
	"math"
	"strconv"
)

func calories() []int {
	filescanner, readFile := getFilescanner("Inputs/day1-1.txt")
	var calories []int = make([]int, 0)
	calories = append(calories, 0)
	for filescanner.Scan() {
		row := filescanner.Text()
		if row == "" {
			calories = append(calories, 0)
		} else {
			calorieRow, err := strconv.Atoi(row)
			if err != nil {
				continue
			}
			calories[len(calories)-1] += calorieRow
		}
	}

	readFile.Close()
	return calories
}

func day1_1() int {
	var maxCal int = 0
	for _, cals := range calories() {
		if cals > maxCal {
			maxCal = cals
		}
	}
	return maxCal
}

func replaceMin(maxes []int, newVal int) []int {
	minval := math.MaxInt
	minIdx := -1
	for i, v := range maxes {
		if v < minval {
			minval = v
			minIdx = i
		}
	}

	if newVal > minval {
		maxes[minIdx] = newVal
	}

	return maxes
}

func day1_2() int {
	var maxCals []int = make([]int, 3)
	for _, cals := range calories() {
		replaceMin(maxCals, cals)
	}
	sum := 0
	for _, cals := range maxCals {
		sum += cals
	}
	return sum
}
