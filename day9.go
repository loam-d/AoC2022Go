package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func correctTLoc(hLoc Loc, tLoc Loc) Loc {
	diffX := hLoc.x - tLoc.x
	diffY := hLoc.y - tLoc.y

	if Abs(diffX) >= 2 {
		if diffX > 0 {
			diffX -= 1
		} else {
			diffX += 1
		}
		tLoc.x += diffX / Abs(diffX)
		if Abs(diffY) >= 1 {
			tLoc.y += (diffY) / Abs(diffY)
		}
	}

	// recalculate the y diff
	diffY = hLoc.y - tLoc.y
	if Abs(diffY) >= 2 {
		if diffY > 0 {
			diffY -= 1
		} else {
			diffY += 1
		}
		tLoc.y += (diffY) / Abs(diffY)
		if Abs(diffX) >= 1 {
			tLoc.x += (diffX) / Abs(diffX)
		}
	}

	return tLoc
}

func doMove(
	moveDir string,
	moveAmt int,
	hLoc Loc,
	tLoc Loc,
	visitedLocs map[Loc]struct{},
) (Loc, Loc) {
	for i := 0; i < moveAmt; i++ {
		switch moveDir {
		case "R":
			hLoc.x += 1
		case "L":
			hLoc.x -= 1
		case "U":
			hLoc.y += 1
		case "D":
			hLoc.y -= 1
		}
		tLoc = correctTLoc(hLoc, tLoc)
		fmt.Println(tLoc, " ", hLoc)
		visitedLocs[tLoc] = struct{}{}
	}
	return hLoc, tLoc
}

func doMoveLong(
	moveDir string,
	moveAmt int,
	locs *[]Loc,
	visitedLocs map[Loc]struct{},
) *[]Loc {
	for i := 0; i < moveAmt; i++ {
		switch moveDir {
		case "R":
			(*locs)[0].x += 1
		case "L":
			(*locs)[0].x -= 1
		case "U":
			(*locs)[0].y += 1
		case "D":
			(*locs)[0].y -= 1
		}
		for i := 1; i < len(*locs); i++ {
			(*locs)[i] = correctTLoc((*locs)[i-1], (*locs)[i])
			if i == len(*locs)-1 {
				visitedLocs[(*locs)[i]] = struct{}{}
			}
		}
	}
	return locs
}

func doMoves() map[Loc]struct{} {
	hLoc := loc(0, 0)
	tLoc := loc(0, 0)

	filescanner, readFile := getFilescanner("Inputs/day9.txt")
	visitedLocs := make(map[Loc]struct{})
	visitedLocs[tLoc] = struct{}{}
	for filescanner.Scan() {
		rawMove := strings.Split(filescanner.Text(), " ")
		moveDir := rawMove[0]
		moveAmt, _ := strconv.Atoi(rawMove[1])
		hLoc, tLoc = doMove(moveDir, moveAmt, hLoc, tLoc, visitedLocs)
	}
	readFile.Close()

	return visitedLocs
}

func doMovesLongRope(numSegs int) map[Loc]struct{} {
	ropeLocs := make([]Loc, 0)
	for i := 0; i < numSegs; i++ {
		ropeLocs = append(ropeLocs, loc(0, 0))
	}

	filescanner, readFile := getFilescanner("Inputs/day9.txt")
	visitedLocs := make(map[Loc]struct{})
	visitedLocs[ropeLocs[1]] = struct{}{}
	for filescanner.Scan() {
		rawMove := strings.Split(filescanner.Text(), " ")
		moveDir := rawMove[0]
		moveAmt, _ := strconv.Atoi(rawMove[1])
		ropeLocs = *doMoveLong(moveDir, moveAmt, &ropeLocs, visitedLocs)
	}
	readFile.Close()

	return visitedLocs

}

func day9_1() (int, int) {
	return len(doMoves()), len(doMovesLongRope(10))
}
