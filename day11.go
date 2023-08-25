package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	items  []int64
	opFn   func(int64) int64
	testFn func(int64) int64
}

func getOpFn(opList []string) func(int64) int64 {
	return func(i int64) int64 {
		var a int64
		var b int64
		if opList[0] == "old" {
			a = i
		} else {
			a, _ = strconv.ParseInt(opList[0], 10, 64)
		}
		if opList[2] == "old" {
			b = i
		} else {
			b, _ = strconv.ParseInt(opList[2], 10, 64)
		}
		switch opList[1] {
		case "+":
			return a + b
		case "*":
			return a * b
		}
		return 0
	}
}

func getTestFn(divis int64, trueTarg int64, falseTarg int64) func(int64) int64 {
	return func(i int64) int64 {
		if i%divis == 0 {
			return trueTarg
		} else {
			return falseTarg
		}
	}
}

func scanMonkeys(filename string) ([]Monkey, int64) {
	filescanner, readFile := getFilescanner(filename)
	defer readFile.Close()
	monkeys := make([]Monkey, 0)
	var currMonkey Monkey
	var moduloConst int64 = 1

	for filescanner.Scan() {
		line := filescanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			currMonkey = Monkey{}
			currMonkey.items = make([]int64, 0)
		} else if strings.HasPrefix(line, "  Starting items:") {
			itemStrs := strings.Split(strings.Split(line, ": ")[1], ", ")
			for _, s := range itemStrs {
				i, _ := strconv.ParseInt(s, 10, 64)
				currMonkey.items = append(currMonkey.items, i)
			}
		} else if strings.HasPrefix(line, "  Operation:") {
			opList := strings.Split(strings.Split(strings.Split(line, ": ")[1], " = ")[1], " ")
			currMonkey.opFn = getOpFn(opList)
		} else if strings.HasPrefix(line, "  Test:") {
			testDivis, _ := strconv.ParseInt(strings.Split(line, ": divisible by ")[1], 10, 64)
			filescanner.Scan()
			line = filescanner.Text()
			trueTarg, _ := strconv.ParseInt(strings.Split(line, " throw to monkey ")[1], 10, 64)
			filescanner.Scan()
			line = filescanner.Text()
			falseTarg, _ := strconv.ParseInt(strings.Split(line, " throw to monkey ")[1], 10, 64)
			moduloConst *= testDivis
			currMonkey.testFn = getTestFn(testDivis, trueTarg, falseTarg)
		} else if line == "" {
			monkeys = append(monkeys, currMonkey)
		}
	}
	monkeys = append(monkeys, currMonkey)
	return monkeys, moduloConst
}

func countInspections(monkeys []Monkey, mod int64) (int64, int64) {
	totalInspections := make([]int64, len(monkeys))
	for round := 0; round < 10000; round++ {
		fmt.Println("round: ", round)
		totalItems := 0
		for i := 0; i < len(monkeys); i++ {
			totalItems += len(monkeys[i].items)
		}
		fmt.Println("    totalItems: ", totalItems)
		for i := 0; i < len(monkeys); i++ {
			for len(monkeys[i].items) > 0 {
				var item int64
				item, monkeys[i].items = monkeys[i].items[0], monkeys[i].items[1:]
				item = monkeys[i].opFn(item) % mod
				target := monkeys[i].testFn(item)
				monkeys[target].items = append(monkeys[target].items, item)
				totalInspections[i]++
			}
			fmt.Println(monkeys[i].items)
		}
		fmt.Println(totalInspections)
	}

	maxInspections := make([]int64, 2)
	var minOfMaxes int64
	maxInspections[0] = totalInspections[0]
	maxInspections[1] = totalInspections[1]
	for i := 2; i < len(totalInspections); i++ {
		if totalInspections[i] > maxInspections[0] && totalInspections[i] > maxInspections[1] {
			if maxInspections[0] > maxInspections[1] {
				minOfMaxes = 1
			} else {
				minOfMaxes = 0
			}
			maxInspections[minOfMaxes] = totalInspections[i]
		} else if totalInspections[i] > maxInspections[0] {
			maxInspections[0] = totalInspections[i]
		} else if totalInspections[i] > maxInspections[1] {
			maxInspections[1] = totalInspections[i]
		}
	}

	return maxInspections[0], maxInspections[1]
}

func day11_1() int64 {
	monkeys, mod := scanMonkeys("Inputs/day11.txt")
	a, b := countInspections(monkeys, mod)
	return a * b
}
