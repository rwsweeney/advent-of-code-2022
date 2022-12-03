package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prioritySum int

	for scanner.Scan() {
		middle := len(scanner.Text()) / 2
		sackA := scanner.Text()[0:middle]
		sackB := scanner.Text()[middle:]

		hashMap := make(map[rune]int)

		for _, v := range sackA {
			hashMap[v] += 1
		}

		for _, v := range sackB {
			if hashMap[v] > 0 {
				prioritySum += priority(v)
				fmt.Println("The shared item is : ", string(v), v)
				break
			}
		}
	}

	fmt.Println("The total priority is: ", prioritySum)

}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		// Subtract 38 to get priority for uppercase numbers
		return int(r) - 38
	} else {
		// Subtract 96 to get priority for lowercase numbers
		return int(r) - 96
	}
}
