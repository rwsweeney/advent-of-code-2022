// 13380 -- Too high

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
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
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
		outcome := round[1]

		totalScore += shape[outcome]

		// Check to see if we should lose, draw, or win, and then choose our play accordingly.
		if outcome == "X" {
			switch opponentMove {
			case "A":
				totalScore += shape["C"]
			case "B":
				totalScore += shape["A"]
			case "C":
				totalScore += shape["B"]
			}
		} else if outcome == "Y" {
			totalScore += shape[opponentMove] + draw
		} else if outcome == "Z" {
			switch opponentMove {
			case "A":
				totalScore += shape["B"]
			case "B":
				totalScore += shape["C"]
			case "C":
				totalScore += shape["A"]
			}

			totalScore += won
		}

	}
	fmt.Println("Total score: ", totalScore)
}

// A rock
// B Paper
// C Scissors

// X lose
// Y Draw
// Z Win
