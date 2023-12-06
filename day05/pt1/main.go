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

func getSeeds(line string) (seeds []int) {
	seedsStr := strings.Split(line, ":")[1]
	seedsNumStr := strings.Split(seedsStr, " ")

	for _, seedNumStr := range seedsNumStr {
		seedNum, err := strconv.Atoi(seedNumStr)

		if err == nil {
			seeds = append(seeds, seedNum)
		}
	}

	return seeds
}

func getMapLine(line string) (almanacMap []int) {
	lineStr := strings.Split(line, " ")
	for _, str := range lineStr {
		int, err := strconv.Atoi(str)

		if err == nil {
			almanacMap = append(almanacMap, int)
		}
	}

	return almanacMap
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	almanac := make(map[string][][]int)
	var seeds []int

	order := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	orderPos := 0
	for i, line := range fileLines {
		if i == 0 {
			seeds = getSeeds(line)
			continue
		}

		if line == "" {
			continue
		}

		if orderPos < len(order) && strings.HasPrefix(line, order[orderPos]) {
			orderPos++
			continue
		}

		currentMap := order[orderPos-1]
		almanac[currentMap] = append(almanac[currentMap], getMapLine(line))
	}

	finalTarget := -1
	for _, seed := range seeds {
		target := -1
		source := seed
		for _, almanacMap := range order {

			for _, almanacMapLine := range almanac[almanacMap] {
				if source >= almanacMapLine[1] && source <= almanacMapLine[1]+almanacMapLine[2]-1 {
					target = almanacMapLine[0] + source - almanacMapLine[1]
				}
			}
			if target == -1 {
				target = source
			}

			source = target
		}
		if finalTarget == -1 {
			finalTarget = target
		} else {
			finalTarget = min(finalTarget, target)
		}
	}

	fmt.Println(finalTarget)
}
