package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("inputs/day1_1.txt")
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
	fmt.Println(sum)
}
