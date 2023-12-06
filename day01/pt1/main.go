package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	finalNum := 0

	for _, line := range fileLines {
		var numStr string
		lineLen := len(line)

		for char := 0; char < lineLen; {
			currentChar := string(line[char])
			_, err := strconv.Atoi(string(line[char]))
			if err == nil {
				numStr += currentChar
				break
			}
			char++
		}
		for char := lineLen - 1; char >= 0; {
			currentChar := string(line[char])
			_, err := strconv.Atoi(currentChar)
			if err == nil {
				numStr += currentChar
				break
			}
			char--
		}

		theRealNum, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		finalNum += theRealNum
	}

	fmt.Println(finalNum)
}
