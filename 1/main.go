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
	var current int
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("Current tally: ", current)
			fmt.Println("Empty line -- continuing")
			current = 0
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Error converting string to number: ", err)
		}
		current += num
		fmt.Println(scanner.Text())
		if current > highest {
			highest = current
			fmt.Println("New highest: ", highest)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Errors encountered while scanning: ", err)
	}

	fmt.Println("The final highest is: ", highest)

}
