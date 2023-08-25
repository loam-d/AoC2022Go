package main

func p1processLine(line string) int {
	opponentPlay := string(line[0])
	var response string = string(line[2])
	score := 0

	xmap := make(map[string]int)
	xmap["A"] = 3
	xmap["B"] = 0
	xmap["C"] = 6

	ymap := make(map[string]int)
	ymap["A"] = 6
	ymap["B"] = 3
	ymap["C"] = 0

	zmap := make(map[string]int)
	zmap["A"] = 0
	zmap["B"] = 6
	zmap["C"] = 3

	switch response {
	case "X":
		score += 1
		score += xmap[opponentPlay]
		break
	case "Y":
		score += 2
		score += ymap[opponentPlay]
		break
	case "Z":
		score += 3
		score += zmap[opponentPlay]
		break
	}

	return score
}

func p2processline(line string) int {
	opponentPlay := string(line[0])
	var response string = string(line[2])
	score := 0

	playscoremap := make(map[string]int)
	playscoremap["rock"] = 1
	playscoremap["paper"] = 2
	playscoremap["scissors"] = 3

	outcomescoremap := make(map[string]int)
	outcomescoremap["X"] = 0
	outcomescoremap["Y"] = 3
	outcomescoremap["Z"] = 6
	score += outcomescoremap[response]

	rockmap := make(map[string]string)
	rockmap["X"] = "scissors"
	rockmap["Y"] = "rock"
	rockmap["Z"] = "paper"

	papermap := make(map[string]string)
	papermap["X"] = "rock"
	papermap["Y"] = "paper"
	papermap["Z"] = "scissors"

	scissorsmap := make(map[string]string)
	scissorsmap["X"] = "paper"
	scissorsmap["Y"] = "scissors"
	scissorsmap["Z"] = "rock"

	var playerchoice string

	switch opponentPlay {
	case "A":
		playerchoice = rockmap[response]
	case "B":
		playerchoice = papermap[response]
	case "C":
		playerchoice = scissorsmap[response]
	}

	score += playscoremap[playerchoice]
	return score
}

func day2_1() int {
	filescanner, readFile := getFilescanner("Inputs/day2.txt")
	score := 0

	for filescanner.Scan() {
		score += p1processLine(filescanner.Text())
	}

	readFile.Close()
	return score
}

func day2_2() int {
	filescanner, readFile := getFilescanner("Inputs/day2.txt")
	score := 0

	for filescanner.Scan() {
		score += p2processline(filescanner.Text())
	}

	readFile.Close()
	return score
}
