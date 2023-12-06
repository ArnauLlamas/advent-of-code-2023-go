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

	timesStr := strings.Split(strings.Split(fileLines[0], ":")[1], "")
	distsStr := strings.Split(strings.Split(fileLines[1], ":")[1], "")

	var timeStr string
	var distStr string
	for _, n := range timesStr {
		_, err := strconv.Atoi(n)

		if err == nil {
			timeStr += n
		}
	}
	for _, n := range distsStr {
		_, err := strconv.Atoi(n)

		if err == nil {
			distStr += n
		}
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)

	optionsToWin := 0

	for t := 0; t <= time; t++ {
		if dist < calculateTravelDist(t, time-t) {
			optionsToWin++
		}
	}

	fmt.Println(optionsToWin)
}
