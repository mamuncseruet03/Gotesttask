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
}

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
