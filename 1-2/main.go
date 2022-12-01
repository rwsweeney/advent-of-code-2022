// Incorrect guess: 203085 -- Too low
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Println("Can't open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var highest int
	var secondHighest int
	var thirdHighest int

	var current int
	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > highest {
				thirdHighest = secondHighest
				secondHighest = highest
				highest = current
				fmt.Println("New highest: ", highest)
			} else if current > secondHighest {
				secondHighest = current
				fmt.Println("New second highest: ", secondHighest)
			} else if current > thirdHighest {
				thirdHighest = current
				fmt.Println("New third highest: ", thirdHighest)
			}
			fmt.Println("Current tally: ", current)
			fmt.Println("Highest: ", highest)
			fmt.Println("Second Highest: ", secondHighest)
			fmt.Println("Third Highest: ", thirdHighest)
			fmt.Println("Empty line -- continuing\n")
			current = 0
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Error converting string to number: ", err)
		}
		current += num
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Errors encountered while scanning: ", err)
	}

	fmt.Println("Highest: ", highest)
	fmt.Println("Second Highest: ", secondHighest)
	fmt.Println("Third Highest: ", thirdHighest)

	fmt.Println("The sum of the top 3: ", highest+secondHighest+thirdHighest)

}
