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

		val, _ := strconv.Atoi(string(first) + string(last))
		sum += val
	}
	file.Close()
	return sum
}

func main() {
	fmt.Println(day1())
}
