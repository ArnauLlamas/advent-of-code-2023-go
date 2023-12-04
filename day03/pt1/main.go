package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func isSymbol(input string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	return reg.MatchString(input)
}

func digitIsSurroundedBySymbol(currLine, prevLine, nextLine string) bool {
	for _, char := range currLine {
		if isSymbol(string(char)) {
			return true
		}
	}

	for _, char := range prevLine {
		if isSymbol(string(char)) {
			return true
		}
	}

	for _, char := range nextLine {
		if isSymbol(string(char)) {
			return true
		}
	}

	return false
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	sumOfNums := 0

	for i, line := range fileLines {
		prev := i - 1
		next := i + 1
		if prev < 0 {
			prev = 0
		}
		if next >= len(fileLines) {
			next = i
		}

		reg := regexp.MustCompile(`\d+`)
		matches := reg.FindAllStringIndex(line, -1)
		for _, match := range matches {
			lower := max(match[0]-1, 0)
			upper := min(match[1]+1, len(line))
			if digitIsSurroundedBySymbol(
				fileLines[i][lower:upper],
				fileLines[prev][lower:upper],
				fileLines[next][lower:upper],
			) {
				num, _ := strconv.Atoi(line[match[0]:match[1]])
				sumOfNums += num
			}
		}
	}

	fmt.Println(sumOfNums)
}
