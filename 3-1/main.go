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

	var setPosition int
	groupItems := make(map[int]string)

	for scanner.Scan() {
		groupItems[setPosition] = scanner.Text()
		setPosition++

		if setPosition > 2 {
			hashMap := make(map[rune]int)
			sharedByTwo := make(map[rune]int)

			for _, v := range groupItems[0] {
				hashMap[v] += 1
			}

			for _, v := range groupItems[1] {
				if hashMap[v] > 0 {
					sharedByTwo[v] += 1
				}
			}

			for _, v := range groupItems[2] {
				if sharedByTwo[v] > 0 {
					prioritySum += priority(v)
					fmt.Println("Shared item is: ", string(v))
					break
				}
			}
			setPosition = 0
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
