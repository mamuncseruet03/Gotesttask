package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

	shortestWord, longestWord, averageWordLength, listOfAverageWord := storyStats(text)

	fmt.Printf("shortestWord = %s longestWord = %s averageWordLength =%d\n", shortestWord, longestWord, averageWordLength)

	fmt.Printf("list Of Average Words \n")

	for _, value := range listOfAverageWord {
		fmt.Printf("Item = %s\n", value)
	}
	fmt.Printf("Generate Strings \n")
	generate(true)

}

var randHelper = initialrandomhelper()

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
// difficulty : 4 out of 10, Estimate: 1 hour, Taken: 1 hour
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
// difficulty : 3 out of 10, Estimate: 1 hour, Taken: 1 hour
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
// difficulty : 4 out of 10, Estimate: 1 hour, Taken: 1 hour
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

/* Function storyStats that returns four things:
// difficulty : 6 out of 10, Estimate: 3 hours, Taken: 2.3 hours
-the shortest word
-the longest word
-the average word length
-the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
*/
func storyStats(texts string) (string, string, int, []string) { //
	var shortestWord = ""
	var longestWord = ""
	var averageWordLength = 0
	var listOfAverageWord []string

	if testValidity(texts) {
		dt := strings.Split(texts, "-")
		var divisor = 0
		var currentLength = 0
		var roundUp = 0
		var roundDown = 0

		shortestWord = dt[1]
		longestWord = dt[1]
		averageWordLength = 0

		for index, value := range dt {
			if index%2 != 0 {
				currentLength = len(value)
				if currentLength > len(longestWord) {
					longestWord = value
				}
				if currentLength < len(shortestWord) {
					shortestWord = value
				}
				averageWordLength = averageWordLength + currentLength
				divisor++
			}
		}
		if divisor != 0 {
			roundDown = averageWordLength / divisor
		}

		roundUp = roundDown
		if (averageWordLength % divisor) > 0 {
			roundUp++
		}
		if divisor != 0 {
			averageWordLength = averageWordLength / divisor
		}

		for index, value := range dt {
			if index%2 == 1 {
				if len(value) >= roundDown && len(value) <= roundUp {
					listOfAverageWord = append(listOfAverageWord, value)
				}
			}
		}
	}

	return shortestWord,
		longestWord,
		averageWordLength,
		listOfAverageWord

}

// function takes boolean flag and generates random correct strings if the parameter is true and random incorrect strings if the flag is false
// difficulty : 7 out of 10, Estimate: 6 hours, Taken: 5 hours

func generate(valid bool) []string {
	var strings []string
	for kn := 0; kn < 5; kn++ {
		k := 0
		numberOfText := randIntInRange(1, 9)

		var words []string
		word := ""

		var intNumbers []string
		inavlidType := randIntInRange(0, 3)
		/*
			0. string first , integer second
			1. 0 length string
			2. integer = float
			3. integer = string
		*/
		inavlidPsition := 0
		if numberOfText > 1 {
			inavlidPsition = randIntInRange(0, (numberOfText - 1))
		}

		for k = numberOfText + 1; k > 0; k-- {
			if inavlidPsition == (numberOfText-(k-1)) && inavlidType == 1 && !valid {
				word = ""
			} else {
				word = randString()
			}
			words = append(words, word)

			if inavlidPsition == (numberOfText-(k-1)) && inavlidType == 2 && !valid {
				word = getRandomInteger() + "." + getRandomInteger()
				intNumbers = append(intNumbers, word)
			} else if inavlidPsition == (numberOfText-(k-1)) && inavlidType == 3 && !valid {
				word = randString()
				intNumbers = append(intNumbers, word)
			} else {
				word = getRandomInteger()
				intNumbers = append(intNumbers, word)
			}
		}

		result := ""

		for k = 0; k < numberOfText; k++ {
			if k > 0 {
				result = result + "-"
			}
			if k == 0 && inavlidType == 0 && !valid {
				result = result + words[k] + "-" + intNumbers[k]
			} else {
				result = result + intNumbers[k] + "-" + words[k]
			}
		}

		println(result)

		strings = append(strings, result)
	}
	return strings
}

func randString() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const letterBytesChar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, getWordLength())
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	a := make([]byte, 1)
	for i := range a {
		a[i] = letterBytesChar[rand.Int63()%int64(len(letterBytesChar))]
	}
	return string(a) + string(b)
}

func getRandomInteger() string {
	const letterBytes = "0123456789"

	b := make([]byte, getIntLength())
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func getWordLength() int {
	return (randIntInRange(1, 10))
}

func getIntLength() int {
	return (randIntInRange(1, 7))
}

func randIntInRange(min int, max int) int {

	randHelper = randHelper + 39
	if randHelper > 99999999 {
		randHelper = randHelper / 99999
	}
	s3 := rand.NewSource(int64(randHelper))
	r3 := rand.New(s3)
	return (r3.Intn(max-min+1) + min)
}

func initialrandomhelper() int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(100) + 3)
}
