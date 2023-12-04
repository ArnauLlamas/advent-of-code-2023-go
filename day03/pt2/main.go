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

func digitsSurroundingGear(x, y int, fileLines []string) []int {
	reg := regexp.MustCompile(`\d+`)
	var ret []int

	for _, line := range fileLines {
		matches := reg.FindAllStringIndex(line, -1)
		for _, match := range matches {
			if match[0]-1 <= y && match[1] >= y {
				num, _ := strconv.Atoi(line[match[0]:match[1]])
				ret = append(ret, num)
			}
		}
	}
	return ret
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	sumOfGearRatios := 0

	for i, line := range fileLines {
		prev := i - 1
		next := i + 2
		if prev < 0 {
			prev = 0
		}
		if next >= len(fileLines) {
			next = len(fileLines)
		}

		reg := regexp.MustCompile(`[*]`)
		matches := reg.FindAllStringIndex(line, -1)
		for _, match := range matches {
			nums := digitsSurroundingGear(i, match[0], fileLines[prev:next])
			if len(nums) == 2 {
				sumOfGearRatios += nums[0] * nums[1]
			}
		}
	}

	fmt.Println(sumOfGearRatios)
}
