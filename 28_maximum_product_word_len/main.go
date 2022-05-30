package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	userInput := getUserInput(reader)

	words := convertInputToWordArray(userInput)

	fmt.Println(maxProduct(words))
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

func convertInputToWordArray(input string) (wordArray []string) {
	for _, word := range strings.Fields(input) {
		wordArray = append(wordArray, word)
	}

	return wordArray
}

func maxProduct(words []string) (result int) {
	sortedWordMap := make(map[int]string)

	for i := 0; i < len(words); i++ {
		for j := 0; j <= i; j++ {
			if len(sortedWordMap[j]) <= len(words[i]) {
				for k := i; k > j; k-- {
					sortedWordMap[k] = sortedWordMap[k-1]
				}

				sortedWordMap[j] = words[i]
				break
			}
		}
	}

	for i := 0; i < len(words); i++ {
		for rest := i + 1; rest < len(words); rest++ {
			isSameLetter := false

			for j := 0; j < len(sortedWordMap[rest]); j++ {
				if strings.Contains(sortedWordMap[i], string(sortedWordMap[rest][j])) {
					isSameLetter = true
					break
				}
			}

			if !isSameLetter {
				product := len(sortedWordMap[i]) * len(sortedWordMap[rest])

				if product > result {
					result = product
				}
			}
		}
	}

	return result
}
