package main

import (
	"fmt"
	"strconv"
	"strings"
)

type d7dir struct {
	name       string
	childFiles []*d7file
	childDirs  []*d7dir
	parentDir  *d7dir
	size       int64
}

type d7file struct {
	name string
	size int64
}

func buildDirTree() d7dir {
	filescanner, readFile := getFilescanner("Inputs/day7.txt")
	var line []string
	var rootDir d7dir
	var currDir *d7dir
	for filescanner.Scan() {
		line = strings.Split(filescanner.Text(), " ")
		if line[0][0] == '$' {
			switch line[1] {
			case "cd":
				switch line[2] {
				case "/":
					rootDir = d7dir{
						name:      "/",
						parentDir: nil,
					}
					currDir = &rootDir
				case "..":
					if currDir.name != "/" {
						currDir = currDir.parentDir
					}
				default:
					targetName := line[2]
					for _, c := range currDir.childDirs {
						if c.name == targetName {
							currDir = c
							break
						}
					}
				}
				break
			case "ls":
				continue
			}
		} else if line[0][0] == 'd' {
			currDir.childDirs = append(currDir.childDirs, &d7dir{
				name:      line[1],
				parentDir: currDir,
			})
			if currDir.name == "/" {
				fmt.Println(currDir)
			}
		} else {
			// This is going to be a file
			size, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				panic(err)
			}
			currDir.childFiles = append(currDir.childFiles, &d7file{
				name: line[1],
				size: size,
			})
		}
	}
	readFile.Close()

	return rootDir
}

func (dir *d7dir) dirSize(dirSizes map[string]int64, sizesBelow100000 *[]int64) int64 {
	var localSize int64 = 0
	for _, f := range dir.childFiles {
		localSize += f.size
	}
	for _, d := range dir.childDirs {
		localSize += d.dirSize(dirSizes, sizesBelow100000)
	}

	if localSize <= 100000 {
		*sizesBelow100000 = append(*sizesBelow100000, localSize)
	}
	dirSizes[dir.name] = localSize
	dir.size = localSize

	return localSize
}

func (dir *d7dir) findMinimallyBigDirectory(atLeast int64) int64 {
	mini := dir.size
	if dir.size >= atLeast {
		for _, d := range dir.childDirs {
			checkedValue := d.findMinimallyBigDirectory(atLeast)
			if checkedValue >= atLeast && checkedValue < mini {
				mini = checkedValue
			}
		}
	}
	return mini
}

func day7_1() (int64, int64) {
	dirTree := buildDirTree()
	dirSizes := make(map[string]int64)
	sizesBelow100000 := make([]int64, 0)
	dirTree.dirSize(dirSizes, &sizesBelow100000)
	var sumTotal int64 = 0
	for _, s := range sizesBelow100000 {
		sumTotal += s
	}
	var atLeast int64 = dirTree.size - 40000000
	return sumTotal, dirTree.findMinimallyBigDirectory(atLeast)
}
