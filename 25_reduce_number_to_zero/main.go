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

	input := convertStringToInt(userInput)

	fmt.Println(numberOfSteps(input))
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

func convertStringToInt(input string) int {
	number, err := strconv.Atoi(input)

	checkError(err)

	return number
}

func numberOfSteps(num int) int {
	steps := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}

		steps++
	}

	return steps
}
