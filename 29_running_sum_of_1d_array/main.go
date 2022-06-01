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

	nums := convertInputToIntArray(userInput)

	fmt.Println(runningSum(nums))
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

func convertInputToIntArray(input string) (intArr []int) {
	inputArray := strings.Fields(input)

	for _, item := range inputArray {
		intValue, err := strconv.Atoi(item)

		checkError(err)

		intArr = append(intArr, intValue)
	}

	return intArr
}

func runningSum(nums []int) (result []int) {
	for index, num := range nums {
		if index == 0 {
			result = append(result, num)
		} else {
			result = append(result, num+result[index-1])
		}
	}

	return result
}
