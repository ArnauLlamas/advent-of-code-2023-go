package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// config:
// 12 red cubes
// 13 green cubes
// 14 blue cubes

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

func isValidGame(sets []string, maxCount map[string]int) bool {
	for _, set := range sets {
		pulls := strings.Split(set, ", ")
		for _, pull := range pulls {
			colorCubes := strings.Split(pull, " ")
			numOfCubes, _ := strconv.Atoi(colorCubes[0])
			colorOfCubes := colorCubes[1]
			if maxCount[colorOfCubes] < numOfCubes {
				return false
			}
		}
	}

	return true
}

func main() {
	maxCount := map[string]int{"red": 12, "green": 13, "blue": 14}

	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	validGamesSum := 0

	for _, line := range fileLines {
		game := strings.Split(line, ": ")
		sets := strings.Split(game[1], "; ")

		if isValidGame(sets, maxCount) {
			gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
			validGamesSum += gameId
		}

	}

	fmt.Println(validGamesSum)
}
