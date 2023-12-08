package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")

	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sumPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		cardInfo := strings.Split(line, "|")
		var winningNumbers string
		var myNumbers string
		for index, numberSection := range cardInfo {
			// winning numbers
			if index == 0 {
				numberSection := strings.TrimSpace(strings.Split(numberSection, ":")[1])
				winningNumbers = numberSection
			} else { // my numbers
				myNumbers = numberSection
				sumPoints += getPointsForCard(winningNumbers, myNumbers)
			}
		}
	}

	file.Close()
	fmt.Println("Sum: " + strconv.Itoa(sumPoints))
}

func getPointsForCard(winningNumbers string, myNumbers string) int {
	points := 0
	winningNumbersSplit := strings.Split(winningNumbers, " ")
	myNumbersSplit := strings.Split(myNumbers, " ")
	previousPointValue := 0
	for _, winningNumber := range winningNumbersSplit {
		winningNumber = strings.TrimSpace(winningNumber)
		for _, myNumber := range myNumbersSplit {
			myNumber = strings.TrimSpace(myNumber)
			if winningNumber != "" && myNumber != "" && winningNumber == myNumber {
				if previousPointValue == 0 {
					previousPointValue = 1
					points += previousPointValue
				} else {
					points += previousPointValue
					previousPointValue = previousPointValue * 2
				}
			}
		}
	}
	return points
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
