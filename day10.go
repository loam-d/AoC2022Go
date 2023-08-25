package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processInstructions(filename string) []int {
	filescanner, readFile := getFilescanner(filename)
	defer readFile.Close()
	cycles := make([]int, 0)
	X := 1

	for filescanner.Scan() {
		rawCommand := strings.Split(filescanner.Text(), " ")
		switch rawCommand[0] {
		case "noop":
			cycles = append(cycles, X)
		case "addx":
			numToAdd, _ := strconv.Atoi(rawCommand[1])
			cycles = append(cycles, X)
			cycles = append(cycles, X)
			X += numToAdd
		}
	}

	return cycles
}

func getSigSum(cycles []int) int {
	sigsum := 0
	for i := 20; i <= 220; i = i + 40 {
		sigprod := i * cycles[i-1]
		sigsum += sigprod
	}

	return sigsum
}

const SCREEN_HEIGHT int = 6
const SCREEN_DEPTH int = 40

func constructScreen(cycles []int) [][]bool {
	screen := make([][]bool, 0)
	for i := 0; i < SCREEN_HEIGHT; i++ {
		screen = append(screen, make([]bool, 0))
		for j := 1; j <= SCREEN_DEPTH; j++ {
			screen[i] = append(screen[i], false)
		}
	}

	row := -1
	for cycleIdx := 0; cycleIdx < len(cycles); cycleIdx++ {
		if cycleIdx%SCREEN_DEPTH == 0 {
			row++
		}
		crtPos := cycleIdx % SCREEN_DEPTH
		currX := cycles[cycleIdx]
		for _, x := range []int{-1, 0, 1} {
			if crtPos == currX+x {
				screen[row][crtPos] = true
				break
			}
		}
	}

	return screen
}

func printSignals(screen [][]bool) {
	for i := 0; i < 6; i++ {
		printedScreen := ""
		for j := 1; j <= SCREEN_DEPTH; j++ {
			var char string
			if screen[i][j-1] {
				char = "██"
			} else {
				char = "  "
			}
			printedScreen += char
		}
		fmt.Println(printedScreen)
	}
}

func day10_1() (int, int) {
	cycles := processInstructions("Inputs/day10.txt")
	cyclesTest := processInstructions("Inputs/day10-test.txt")
	screen := constructScreen(cycles)
	printSignals(screen)
	return getSigSum(cycles), getSigSum(cyclesTest)
}
