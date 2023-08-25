package main

func findSignalStart(signals string, windowSize int) int {
	for i := windowSize - 1; i < len(signals); i++ {
		window := signals[i-(windowSize-1) : i]
		isunique := make(map[rune]struct{})
		for _, c := range window {
			isunique[c] = struct{}{}
		}
		isunique[rune(signals[i])] = struct{}{}
		if len(isunique) == windowSize {
			return i + 1
		}
	}
	return -1
}

func day6_1() (int, int) {
	filescanner, readFile := getFilescanner("Inputs/day6.txt")
	var signals string
	for filescanner.Scan() {
		signals = filescanner.Text()
		break
	}
	readFile.Close()

	return findSignalStart(signals, 4), findSignalStart(signals, 14)
}
