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

	firstInput, err := getUserInput(reader)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	secondInput, err := getUserInput(reader)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("Greatest Common Factore of %d and %d is %d\n", firstInput, secondInput, getGreatestCommonFactor(firstInput, secondInput))

}

func getUserInput(reader *bufio.Reader) (int32, error) {
	fmt.Print("Please enter a number: ")

	input, _, err := reader.ReadLine()

	if err != nil {
		return 0, err
	}

	num, err := convertInputToInt(string(input))

	if err != nil {
		return 0, err
	}

	return num, nil
}

func convertInputToInt(input string) (int32, error) {
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return int32(num), nil
}

func getGreatestCommonFactor(num1 int32, num2 int32) int32 {
	var temp int32

	for num2 != 0 {
		temp = num1 % num2
		num1 = num2
		num2 = temp
	}

	return num1
}
