package main

type NodeCoords struct {
	X, Y int
}

func runeToValue(c rune) int {
	if c == 'S' {
		return 0
	} else if c == 'E' {
		return int('z' - 'a')
	} else {
		return int(c - 'a')
	}
}

func genrateGraph(filename string, isPart2 bool) (map[NodeCoords][]NodeCoords, []NodeCoords, NodeCoords) {
	filescanner, readFile := getFilescanner(filename)
	defer readFile.Close()
	graph := make(map[NodeCoords][]NodeCoords)
	values := make(map[NodeCoords]int)
	starts := make([]NodeCoords, 0)
	var end NodeCoords

	row := 0
	for filescanner.Scan() {
		line := filescanner.Text()
		for col, c := range line {
			nodeCoords := NodeCoords{X: row, Y: col}
			graph[nodeCoords] = make([]NodeCoords, 0)
			values[nodeCoords] = runeToValue(c)
			if row > 0 {
				upperNeighbor := NodeCoords{X: row - 1, Y: col}
				diff := values[upperNeighbor] - values[nodeCoords]
				if diff <= 1 {
					graph[nodeCoords] = append(graph[nodeCoords], upperNeighbor)
				}
				if diff >= -1 {
					graph[upperNeighbor] = append(graph[upperNeighbor], nodeCoords)
				}
			}
			if col > 0 {
				leftNeighbor := NodeCoords{X: row, Y: col - 1}
				diff := values[leftNeighbor] - values[nodeCoords]
				if diff <= 1 {
					graph[nodeCoords] = append(graph[nodeCoords], leftNeighbor)
				}
				if diff >= -1 {
					graph[leftNeighbor] = append(graph[leftNeighbor], nodeCoords)
				}
			}
			if c == 'S' || (c == 'a' && isPart2) {
				starts = append(starts, nodeCoords)
			}
			if c == 'E' {
				end = nodeCoords
			}
		}
		row++
	}

	return graph, starts, end
}

func coordsEqual(a NodeCoords, b NodeCoords) bool {
	return a.X == b.X && a.Y == b.Y
}

func BFS(graph map[NodeCoords][]NodeCoords, starts []NodeCoords, end NodeCoords) int {
	parent := make(map[NodeCoords]NodeCoords)
	explored := make(map[NodeCoords]struct{})
	startCoords := make(map[NodeCoords]struct{})

	q := make([]NodeCoords, 0)
	for _, start := range starts {
		q = append(q, start)
		explored[start] = struct{}{}
		startCoords[start] = struct{}{}
	}
	var head NodeCoords

	for len(q) > 0 {
		head, q = q[0], q[1:]
		if head == end {
			curr := head
			lenPath := 0
			_, isInStartCoods := startCoords[curr]
			for !isInStartCoods {
				curr = parent[curr]
				lenPath++
				_, isInStartCoods = startCoords[curr]
			}
			return lenPath
		}
		for _, neighbor := range graph[head] {
			_, isExplored := explored[neighbor]
			if isExplored {
				continue
			}
			q = append(q, neighbor)
			explored[neighbor] = struct{}{}
			parent[neighbor] = head
		}
	}
	return -1
}

func day12_1() int {
	graph, starts, end := genrateGraph("Inputs/day12.txt", true)
	return BFS(graph, starts, end)
}
