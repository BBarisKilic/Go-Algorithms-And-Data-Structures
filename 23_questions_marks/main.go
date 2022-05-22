package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	userInput := getUserInput(reader)

	fmt.Println(questionsMarks(userInput))
}

func getUserInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	checkError(err)

	input = strings.TrimSpace(input)

	return input
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func questionsMarks(str string) (result string) {
	chars := []byte(str)

	totalNumberOfQuestionsMarks := 0
	result = "false"
	startIndex := -1
	endIndex := -1

	for index, char := range chars {
		if char >= 48 && char <= 57 {
			if startIndex == -1 {
				startIndex = index
			} else {
				endIndex = index
			}
		}

		if endIndex != -1 {
			firstNumber, err := strconv.Atoi(string(chars[startIndex]))

			checkError(err)

			secondNumber, err := strconv.Atoi(string(chars[endIndex]))

			if firstNumber+secondNumber == 10 {
				for i := startIndex + 1; i < endIndex; i++ {
					if string(chars[i]) == "?" {
						totalNumberOfQuestionsMarks++
					}
				}

				if totalNumberOfQuestionsMarks == 3 {
					result = "true"
				} else {
					return "false"
				}
			}

			totalNumberOfQuestionsMarks = 0
			startIndex = -1
			endIndex = -1
		}

	}

	return result
}
