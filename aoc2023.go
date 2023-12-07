package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getScanner(filename string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file), file
}

func day1() int {
	scanner, file := getScanner("inputs/day1.txt")
	sum := 0
	for scanner.Scan() {
		var firstPos, lastPos int = -1, -1
		var first, last rune
		line := scanner.Text()

		numberNames := []string{
			"1", "2", "3",
			"4", "5", "6",
			"7", "8", "9",
			"one", "two", "three",
			"four", "five", "six",
			"seven", "eight", "nine",
		}
		for number, numberName := range numberNames {
			position := strings.Index(line, numberName)
			if position < 0 {
				continue
			}
			if position < firstPos || firstPos == -1 {
				first = rune(number%9) + rune('1')
				firstPos = position
			}
			lastPosition := strings.LastIndex(line, numberName)
			if lastPosition > lastPos {
				last = rune(number%9) + rune('1')
				lastPos = lastPosition
			}
		}

		val, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	}
	file.Close()
	return sum
}

func day2(day2Part1 bool) int {
	scanner, file := getScanner("inputs/day2.txt")
	defer file.Close()

	const maxRedCount, maxGreenCount, maxBlueCount int = 12, 13, 14
	var redCount, greenCount, blueCount int
	var gameHighestRedCount, gameHighestGreenCount, gameHighestBlueCount int
	var possible bool
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameInfo := strings.Split(line, ": ")
		gameSets := strings.Split(gameInfo[1], "; ")
		possible = true
		gameHighestRedCount, gameHighestGreenCount, gameHighestBlueCount = 0, 0, 0

		for _, set := range gameSets {
			redCount, greenCount, blueCount = 0, 0, 0
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				if strings.Contains(cube, "red") {
					redCount += cubeCount
				} else if strings.Contains(cube, "green") {
					greenCount += cubeCount
				} else if strings.Contains(cube, "blue") {
					blueCount += cubeCount
				}
			}
			gameHighestRedCount = max(redCount, gameHighestRedCount)
			gameHighestBlueCount = max(blueCount, gameHighestBlueCount)
			gameHighestGreenCount = max(greenCount, gameHighestGreenCount)

			if day2Part1 {
				if redCount > maxRedCount || greenCount > maxGreenCount || blueCount > maxBlueCount {
					possible = false
					break
				}
			}
		}

		if day2Part1 {
			if possible {
				gameId, err := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
				if err != nil {
					log.Fatal(err)
				}
				result += gameId
			}
		} else {
			result += gameHighestRedCount * gameHighestGreenCount * gameHighestBlueCount
		}
	}

	return result
}

func main() {
	fmt.Println(day2(false))
}
