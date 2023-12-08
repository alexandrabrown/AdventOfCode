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
	sum := 0

	for scanner.Scan() {
		gameLine := scanner.Text()
		gameSets := strings.Split(gameLine, "; ")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for index, gameInfo := range gameSets {
			if index == 0 {
				gameString := strings.TrimPrefix(gameInfo, "Game")
				gameNumString := strings.TrimSpace(gameString[:strings.Index(gameString, ":")])
				gameInfo = strings.TrimPrefix(gameInfo, "Game "+gameNumString+": ")
			}
			getCubesNeeded(gameInfo, &maxRed, &maxGreen, &maxBlue)
		}
		sum += (maxRed * maxGreen * maxBlue)
	}

	file.Close()
	fmt.Println("Power: " + strconv.Itoa(sum))
}

func getCubesNeeded(gameInfo string, maxRed *int, maxGreen *int, maxBlue *int) {
	gameReveal := strings.Split(gameInfo, ", ")
	for _, game := range gameReveal {
		gameNumColors := strings.Split(game, " ")
		var num int
		var err error
		for index, gameNumColor := range gameNumColors {
			if index%2 == 0 {
				num, err = strconv.Atoi(gameNumColor)
				check(err)
			} else {
				color := gameNumColor
				if color == "red" && num > *maxRed {
					*maxRed = num
				}
				if color == "green" && num > *maxGreen {
					*maxGreen = num
				}
				if color == "blue" && num > *maxBlue {
					*maxBlue = num
				}
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
