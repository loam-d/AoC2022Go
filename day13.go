package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Nary struct {
	val    int
	sl     []Nary
	hasVal bool
}

func (n Nary) String() string {
	if n.hasVal {
		return strconv.Itoa(n.val)
	}
	str := "["
	for i, v := range n.sl {
		curr := v.String()
		if i > 0 {
			str += ", "
		}
		str += curr
	}
	str += "]"
	return str
}

func (n Nary) getNaryValue() (int, []Nary, bool) {
	if n.hasVal {
		return n.val, nil, true
	}
	return 0, n.sl, false
}

func parseLine(line string) Nary {
	res := Nary{hasVal: false, sl: make([]Nary, 0)}
	currSlice := &res.sl
	parentStack := make([]*[]Nary, 0)
	parentStack = append(parentStack, currSlice)
	currDigit := ""
	if len(line) == 2 {
		return res
	}
	for _, c := range line[1:] {
		switch c {
		case '[':
			newSlice := Nary{hasVal: false, sl: make([]Nary, 0)}
			currSlice = &newSlice.sl
			parentStack = append(parentStack, currSlice)
		case ']':
			if currDigit != "" {
				appendDigit, _ := strconv.Atoi(currDigit)
				valNary := Nary{hasVal: true, val: appendDigit}
				*currSlice = append(*currSlice, valNary)
				currDigit = ""
			}
			if len(parentStack) > 1 {
				toAppend := Nary{hasVal: false, sl: *currSlice}
				*parentStack[len(parentStack)-2] = append(*parentStack[len(parentStack)-2], toAppend)
				currSlice = parentStack[len(parentStack)-2]
			}
			parentStack = parentStack[:len(parentStack)-1]
		case ',':
			if currDigit != "" {
				appendDigit, _ := strconv.Atoi(currDigit)
				valNary := Nary{hasVal: true, val: appendDigit}
				*currSlice = append(*currSlice, valNary)
				currDigit = ""
			}
		default:
			currDigit += string(c)
		}
	}
	return res
}

func processPairs(filename string) ([]Nary, []Nary) {
	filescanner, readFile := getFilescanner(filename)
	defer readFile.Close()

	left := make([]Nary, 0)
	right := make([]Nary, 0)

	lineNum := 0
	for filescanner.Scan() {
		line := filescanner.Text()
		if lineNum%3 == 0 {
			left = append(left, parseLine(line))
		} else if lineNum%3 == 1 {
			right = append(right, parseLine(line))
		}
		lineNum += 1
	}

	return left, right
}

func isInRightOrder(l Nary, r Nary) int {
	switch l.hasVal {
	case true:
		switch r.hasVal {
		case true:
			// Base case, both items are integers. Return value determines behavior
			if l.val < r.val {
				return 1
			} else if l.val > r.val {
				return -1
			} else {
				return 0
			}
		case false:
			// L is a value, and R is a list
			// We need to convert the value to a list then compare again
			lcmp := Nary{hasVal: false, sl: []Nary{{hasVal: true, val: l.val}}}
			return isInRightOrder(lcmp, r)
		}
	case false:
		switch r.hasVal {
		case true:
			// R is a value, and L is a list
			// We need to convert the value to a list then compare again
			rcmp := Nary{hasVal: false, sl: []Nary{{hasVal: true, val: r.val}}}
			return isInRightOrder(l, rcmp)
		case false:
			// Both are lists, check first check lens, then check elements until you can't
			for i := 0; ; i++ {
				if len(l.sl) == i && len(r.sl) == i {
					return 0
				} else if len(l.sl) == i {
					return 1
				} else if len(r.sl) == i {
					return -1
				}
				check := isInRightOrder(l.sl[i], r.sl[i])
				if check != 0 {
					return check
				}
			}
		}
	}
	return 0
}

func countOrders(left []Nary, right []Nary) int {
	res := 0
	for k := 0; k < len(left); k++ {
		l := left[k]
		r := right[k]
		if isInRightOrder(l, r) == 1 {
			res += k + 1
		}
	}
	return res
}

func mergeSortedPackets(left []Nary, right []Nary) []Nary {
	res := make([]Nary, 0)
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			res = append(res, right[0])
			right = right[1:]
			continue
		}
		if len(right) == 0 {
			res = append(res, left[0])
			left = left[1:]
			continue
		}
		l := left[0]
		r := right[0]
		if isInRightOrder(l, r) == 1 {
			res = append(res, l)
			left = left[1:]
		} else {
			res = append(res, r)
			right = right[1:]
		}
	}
	return res
}

func insertDivider(res []Nary) []Nary {
	divided := make([]Nary, 0)
	two := parseLine("[[2]]")
	six := parseLine("[[6]]")
	twoIn := false
	sixIn := false
	for i, n := range res {
		if i == 0 {
			if isInRightOrder(two, n) == 1 {
				divided = append(divided, two)
				twoIn = true
			}
			if isInRightOrder(six, n) == 1 {
				divided = append(divided, six)
				sixIn = true
			}
			divided = append(divided, n)
		} else if i == len(res)-1 {
			divided = append(divided, n)
			if !twoIn && isInRightOrder(n, two) == 1 {
				divided = append(divided, two)
				twoIn = true
			}
			if !sixIn && isInRightOrder(n, six) == 1 {
				divided = append(divided, six)
				sixIn = true
			}
		} else {
			if !twoIn && isInRightOrder(res[i-1], two) == 1 && isInRightOrder(two, n) == 1 {
				divided = append(divided, two)
				twoIn = true
			}
			if !sixIn && isInRightOrder(res[i-1], six) == 1 && isInRightOrder(six, n) == 1 {
				divided = append(divided, six)
				sixIn = true
			}
			divided = append(divided, n)
		}
	}
	return divided
}

func findAndMultiplyDividerIndices(divided []Nary) int {
	mult := 1
	two := "[[2]]"
	six := "[[6]]"
	for i, n := range divided {
		fmt.Println(n.String())
		if n.String() == two {
			fmt.Println("========================")
			mult *= i + 1
		}
		if n.String() == six {
			fmt.Println("========================")
			mult *= i + 1
		}
	}
	return mult
}

func day13_1() (int, int) {
	left, right := processPairs("Inputs/day13.txt")
	merged := mergeSortedPackets(left, right)
	sort.Slice(merged, func(i, j int) bool {
		return isInRightOrder(merged[i], merged[j]) == 1
	})
	divided := insertDivider(merged)
	fmt.Println(divided)
	return countOrders(left, right), findAndMultiplyDividerIndices(divided)
}
