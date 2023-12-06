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

func addCards(cards map[int]int, cardN int, r int) {
	cards[cardN] += r
}

func main() {
	fileLines, err := readFile("../inputs.txt")
	if err != nil {
		panic(err)
	}

	totalCards := make(map[int]int)

	for _, line := range fileLines {
		card := strings.Split(line, ":")
		cardNums := strings.Split(card[1], "|")
		cardGame := strings.Split(card[0], " ")
		cardGameNum, _ := strconv.Atoi(cardGame[len(cardGame)-1])

		addCards(totalCards, cardGameNum, 1)

		winnerNumsStr := strings.Split(cardNums[0], " ")
		ownedNumsStr := strings.Split(cardNums[1], " ")

		winnerNums := strSliceToIntSlice(winnerNumsStr)
		ownedNums := strSliceToIntSlice(ownedNumsStr)

		ownedCards := totalCards[cardGameNum]

		matches := 0
		for _, winnerNum := range winnerNums {
			if slices.Contains(ownedNums, winnerNum) {
				matches++
			}
		}

		for matches > 0 {
			addCards(totalCards, cardGameNum+matches, ownedCards)
			matches--
		}
	}

	totalCardsSum := 0
	for _, cards := range totalCards {
		totalCardsSum += cards
	}

	fmt.Println(totalCardsSum)
}
