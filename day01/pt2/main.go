package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/dlclark/regexp2"
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

func findFirstDigit(line string) (pos int, digit string) {
	lineLen := len(line)
	for char := 0; char < lineLen; {
		currentChar := string(line[char])
		_, err := strconv.Atoi(currentChar)
		if err == nil {
			return char, currentChar
		}
		char++
	}

	return -1, ""
}

func findLastDigit(line string) (pos int, digit string) {
	lineLen := len(line)
	for char := lineLen - 1; char >= 0; {
		currentChar := string(line[char])
		_, err := strconv.Atoi(currentChar)
		if err == nil {
			return char, currentChar
		}
		char--
	}

	return -1, "-1"
}

func digitStrToDigit(d string) string {
	numStrs := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	return numStrs[d]
}

func findFirstDigitStr(line string) (pos int, digit string) {
	reg := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")

	matches := reg.FindAllStringIndex(line, -1)
	if len(matches) == 0 {
		return -1, "-1"
	}

	return matches[0][0], digitStrToDigit(line[matches[0][0]:matches[0][1]])
}

func findLastDigitStr(line string) (pos int, digit string) {
	reg := regexp2.MustCompile("(?=(one|two|three|four|five|six|seven|eight|nine))", regexp2.RE2)
	m, _ := reg.FindStringMatch(line)

	if m == nil {
		return -1, "-1"
	}
	for m != nil {
		digit = m.Groups()[1].Captures[0].String()
		pos = m.Groups()[1].Captures[0].Index
		m, _ = reg.FindNextMatch(m)
	}

	return pos, digitStrToDigit(digit)
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	theFinalNum := 0

	for _, line := range fileLines {
		p1, n1 := findFirstDigit(line)
		p2, n2 := findLastDigit(line)

		p3, n3 := findFirstDigitStr(line)
		p4, n4 := findLastDigitStr(line)

		var numStr string
		if p1 != -1 && p3 == -1 ||
			p1 < p3 && p1 != -1 {
			numStr = n1
		} else {
			numStr = n3
		}

		if p2 >= p4 && p2 != -1 {
			numStr += n2
		} else {
			numStr += n4
		}

		theRealNum, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		theFinalNum += theRealNum
	}

	fmt.Println(theFinalNum)
}
