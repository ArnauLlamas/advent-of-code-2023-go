package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filepath string) (fileLines []string, err error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	err = readFile.Close()
	if err != nil {
		return nil, err
	}

	return fileLines, nil
}

func strSliceToIntSlice(strSlice []string) (intSlice []int) {
	for _, str := range strSlice {
		int, err := strconv.Atoi(str)
		if err == nil {
			intSlice = append(intSlice, int)
		}
	}

	return intSlice
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	totalCardPoints := 0

	for _, line := range fileLines {
		cardNums := strings.Split(strings.Split(line, ":")[1], "|")

		winnerNumsStr := strings.Split(cardNums[0], " ")
		ownedNumsStr := strings.Split(cardNums[1], " ")

		winnerNums := strSliceToIntSlice(winnerNumsStr)
		ownedNums := strSliceToIntSlice(ownedNumsStr)

		cardPoints := 0
		for _, winnerNum := range winnerNums {
			if slices.Contains(ownedNums, winnerNum) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}
		totalCardPoints += cardPoints
	}

	fmt.Println(totalCardPoints)
}
