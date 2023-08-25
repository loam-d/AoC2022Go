package main

func getPrio(c rune) int {
	runeInt := int(c)
	if runeInt >= 65 && runeInt <= 90 {
		return runeInt - int('A') + 27
	} else if runeInt >= 97 && runeInt <= 122 {
		return runeInt - int('a') + 1
	}
	return 0
}

func day3p1processLine(line string) int {
	firstCompartment := map[rune]struct{}{}
	for i, c := range line {
		if i < len(line)/2 {
			firstCompartment[c] = struct{}{}
		} else {
			if _, ok := firstCompartment[c]; ok {
				return getPrio(c)
			}
		}
	}
	return 0
}

func day3p2processGroup(group []string) int {
	badgeCandidates := map[rune]struct{}{}
	narrowedCandidates := map[rune]struct{}{}
	for i, s := range group {
		for _, c := range s {
			if i == 0 {
				badgeCandidates[c] = struct{}{}
			} else if i == 2 {
				if _, ok := narrowedCandidates[c]; ok {
					return getPrio(c)
				}
			} else {
				if _, ok := badgeCandidates[c]; ok {
					narrowedCandidates[c] = struct{}{}
				}
			}
		}
	}

	return 0
}

func day3_1() int {
	filescanner, readFile := getFilescanner("Inputs/day3.txt")
	prioSum := 0
	for filescanner.Scan() {
		prioSum += day3p1processLine(filescanner.Text())
	}

	readFile.Close()
	return prioSum
}

func day3_2() int {
	filescanner, readFile := getFilescanner("Inputs/day3.txt")
	group := make([]string, 0)
	prioSum := 0
	for filescanner.Scan() {
		group = append(group, filescanner.Text())
		if len(group) == 3 {
			prioSum += day3p2processGroup(group)
			group = make([]string, 0)
		}
	}

	readFile.Close()
	return prioSum
}
