package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func day3() (int, int) {
	type PotentialPartPos struct {
		row   int
		start int
		end   int
	}
	type PotentialGearPos struct {
		row int
		col int
	}
	type gearValue struct {
		adjacentparts int
		gearRatio     int
	}

	scanner, file := getScanner("inputs/day3.txt")
	defer file.Close()
	rowNumber := 0

	var schematic [][]rune
	var potentialParts []PotentialPartPos
	var currentPart *PotentialPartPos
	inNumber := false
	potentialGears := make(map[PotentialGearPos]gearValue)

	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
		for colNumber, character := range line {
			if character >= '0' && character <= '9' && !inNumber {
				inNumber = true
				currentPart = new(PotentialPartPos)
				currentPart.row = rowNumber
				currentPart.start = colNumber
			} else if (character < '0' || character > '9') && inNumber {
				currentPart.end = colNumber - 1
				potentialParts = append(potentialParts, *currentPart)
				inNumber = false
			}
			if character == '*' {
				potentialGears[PotentialGearPos{
					rowNumber,
					colNumber,
				}] = gearValue{0, 1}
			}
		}
		if inNumber {
			currentPart.end = len(line) - 1
			potentialParts = append(potentialParts, *currentPart)
			inNumber = false
		}
		rowNumber++
	}

	partNumberSum := 0
	checkSymbol := func(c rune) bool { return c != '.' && (c < '0' || c > '9') }
	incrementGear := func(x int, y int, partNum int) {
		gear := potentialGears[PotentialGearPos{x, y}]
		gear.adjacentparts += 1
		gear.gearRatio *= partNum
		potentialGears[PotentialGearPos{x, y}] = gear
	}

	lastRow := len(schematic) - 1
	lastCol := len(schematic[0]) - 1
	var symbolFound bool

	for _, part := range potentialParts {
		symbolFound = false
		partNum, _ := strconv.Atoi(
			string(schematic[part.row][part.start : part.end+1]))

		if part.start > 0 {
			if checkSymbol(schematic[part.row][part.start-1]) {
				symbolFound = true
			}
			if schematic[part.row][part.start-1] == '*' {
				incrementGear(part.row, part.start-1, partNum)
			}
		}
		if part.end < lastCol {
			if checkSymbol(schematic[part.row][part.end+1]) {
				symbolFound = true
			}
			if schematic[part.row][part.end+1] == '*' {
				incrementGear(part.row, part.end+1, partNum)
			}
		}

		start := max(part.start-1, 0)
		end := min(part.end+1, lastCol)

		for i := start; i <= end; i++ {
			if part.row > 0 {
				if checkSymbol(schematic[part.row-1][i]) {
					symbolFound = true
				}
				if schematic[part.row-1][i] == '*' {
					incrementGear(part.row-1, i, partNum)
				}
			}
			if part.row < lastRow {
				if checkSymbol(schematic[part.row+1][i]) {
					symbolFound = true
				}
				if schematic[part.row+1][i] == '*' {
					incrementGear(part.row+1, i, partNum)
				}
			}
		}
		if symbolFound {
			partNumberSum += partNum
		}
	}

	gearSum := 0

	for _, potentialGear := range potentialGears {
		if potentialGear.adjacentparts == 2 {
			gearSum += potentialGear.gearRatio
		}
	}

	return partNumberSum, gearSum
}

func day4() int {
	scanner, file := getScanner("inputs/day4.txt")
	defer file.Close()

	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")
		gameParts := strings.Split(game[1], "|")
		gamePoints := 0

		var gameNumbers []int
		var winningNumbers []int

		for _, number := range strings.Fields(gameParts[0]) {
			num, _ := strconv.Atoi(number)
			gameNumbers = append(gameNumbers, num)
		}
		for _, number := range strings.Fields(gameParts[1]) {
			num, _ := strconv.Atoi(number)
			winningNumbers = append(winningNumbers, num)
		}

		for _, number := range gameNumbers {
			if slices.Contains(winningNumbers, number) {
				if gamePoints > 0 {
					gamePoints *= 2
				} else {
					gamePoints = 1
				}
			}
		}

		totalPoints += gamePoints
	}

	return totalPoints
}

func main() {
	fmt.Println(day4())
}
