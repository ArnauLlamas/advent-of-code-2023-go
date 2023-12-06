package main

import (
	"bufio"
	"fmt"
	"os"
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

func calculateTravelDist(timeButt int, timeRem int) int {
	return timeButt * timeRem
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	timesStr := strings.Split(strings.Split(fileLines[0], ":")[1], " ")
	distsStr := strings.Split(strings.Split(fileLines[1], ":")[1], " ")

	times := []int{}
	dists := []int{}
	for _, n := range timesStr {
		t, err := strconv.Atoi(n)

		if err == nil {
			times = append(times, t)
		}
	}
	for _, n := range distsStr {
		d, err := strconv.Atoi(n)

		if err == nil {
			dists = append(dists, d)
		}
	}

	optionsToWin := []int{}
	for race := 0; race < len(times); race++ {
		optionsToWinCurrentRace := 0

		for t := 0; t <= times[race]; t++ {
			if dists[race] < calculateTravelDist(t, times[race]-t) {
				optionsToWinCurrentRace++
			}
		}
		optionsToWin = append(optionsToWin, optionsToWinCurrentRace)
	}

	finalScore := 1
	for i := 0; i < len(optionsToWin); i++ {
		finalScore *= optionsToWin[i]
	}
	fmt.Println(finalScore)
}
