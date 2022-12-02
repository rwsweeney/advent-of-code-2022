package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// Additional score for results
	lost int = 0
	draw int = 3
	won  int = 6

	// Character differences in bytes to make A=X for rock paper scissors.
	charDifference = 23
)

func main() {

	// Score for playing each shape
	shape := map[string]int{
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}

	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("Failed to open file: ", err)
	}

	defer file.Close()
	if err != nil {
		log.Println("Error closing file: ", err)
	}

	scanner := bufio.NewScanner(file)

	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Split(line, " ")
		opponentMove := round[0]
		myMove := round[1]

		totalScore += shape[myMove]

		// Look for draw condition or the three win conditions else losing round
		if line[2]-23 == line[0] { // Compare the byte differences
			totalScore += draw
		} else if myMove == "X" && opponentMove == "C" {
			totalScore += won
		} else if myMove == "Y" && opponentMove == "A" {
			totalScore += won
		} else if myMove == "Z" && opponentMove == "B" {
			totalScore += won
		}

		fmt.Println("Total score: ", totalScore)
	}

}
