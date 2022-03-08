package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	text := "23-ab-48-caba-56-haha"

	// For Test
	if testValidity(text) {
		fmt.Println("Valid")
	} else {
		fmt.Println("InValid")
	}

	fmt.Printf("Average Number = %d\n", averageNumber(text))

	fmt.Printf("wholeStory = %s\n", wholeStory(text))

}

// Check ASCII String
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// Function to check whether the input text is valid or invalid
func testValidity(texts string) bool {

	dt := strings.Split(texts, "-")

	// check sepertated numbers and string count
	if len(dt)%2 != 0 {
		return false
	}

	for index, value := range dt {
		if index%2 != 0 {
			// check if  assci
			if !(isASCII(value)) {
				return false
			}
		} else {

			// check interger
			_, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Not Number")
				return false
			}

		}
	}
	return true
}

// Function to calculate the average number from all the numbers
func averageNumber(texts string) int {

	// check validity
	if testValidity(texts) {

		dt := strings.Split(texts, "-")

		var total = 0
		var divisor = 0

		for index, value := range dt {
			if index%2 == 0 {
				intValue, _ := strconv.Atoi(value)
				total = total + intValue
				divisor++
			}
		}

		if divisor > 0 {
			return total / divisor
		}

		return 0
	} else {
		return 0
	}

}

// Function wholeStory that takes the string, and returns a text that is composed from all the text words separated
func wholeStory(texts string) string {

	var wholeStory = ""

	if testValidity(texts) {

		dt := strings.Split(texts, "-")

		for index, value := range dt {
			if index%2 != 0 {

				if index != 1 {
					wholeStory = wholeStory + " " + value

				} else {
					wholeStory = wholeStory + value
				}
			}
		}

		return wholeStory
	}

	return wholeStory
}
