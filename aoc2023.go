package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1_1() int {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var first, last rune
		for _, letter := range scanner.Text() {
			if letter >= rune('0') && letter <= rune('9') {
				if first == 0 {
					first = letter
					last = letter
				} else {
					last = letter
				}
			}
		}
		val, _ := strconv.Atoi(string(first) + string(last))
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(day1_1())
}
