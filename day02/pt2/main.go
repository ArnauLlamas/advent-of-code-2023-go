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

func powerOfCubesInGame(sets []string) int {
	red := 0
	green := 0
	blue := 0

	for _, set := range sets {
		pulls := strings.Split(set, ", ")

		for _, pull := range pulls {
			cubes := strings.Split(pull, " ")
			num, _ := strconv.Atoi(cubes[0])
			color := cubes[1]

			switch color {
			case "red":
				red = max(red, num)
			case "green":
				green = max(green, num)
			case "blue":
				blue = max(blue, num)
			}
		}
	}

	return red * green * blue
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	powerOfCubesSum := 0

	for _, line := range fileLines {
		game := strings.Split(line, ": ")
		sets := strings.Split(game[1], "; ")

		powerOfCubesSum += powerOfCubesInGame(sets)

	}

	fmt.Println(powerOfCubesSum)
}
