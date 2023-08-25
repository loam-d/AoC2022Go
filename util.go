package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFilescanner(inputFile string) (*bufio.Scanner, *os.File) {
	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	return filescanner, readFile
}
