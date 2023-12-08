package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := 0

	for scanner.Scan() {
		gameLine := scanner.Text()
		gameSets := strings.Split(gameLine, "; ")
		power := 0
		for index, gameInfo := range gameSets {
			if index == 0 {
				gameString := strings.TrimPrefix(gameInfo, "Game")
				gameNumString := strings.TrimSpace(gameString[:strings.Index(gameString, ":")])
				gameInfo = strings.TrimPrefix(gameInfo, "Game "+gameNumString+": ")
			}
			power = minimumCubesNeededPower(gameInfo)
		}
		sum += power
	}

	file.Close()
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func minimumCubesNeededPower(gameInfo string) int {
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
				if color == "red" && num > ALLOWED_RED {
					return false
				}
				if color == "green" && num > ALLOWED_GREEN {
					return false
				}
				if color == "blue" && num > ALLOWED_BLUE {
					return false
				}
			}
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
