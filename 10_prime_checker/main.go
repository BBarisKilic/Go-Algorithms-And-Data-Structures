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

	input, err := getUserInput(reader)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println(checkIfAnyOfPermutationNumberIsPrime(input))

}

func getUserInput(reader *bufio.Reader) (int, error) {
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

func convertInputToInt(input string) (int, error) {
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func checkIfAnyOfPermutationNumberIsPrime(num int) bool {

	for i := 1; i < num; i++ {
		if num%i == 0 {
			isPrime := checkIfNumberIsPrime(i)

			if isPrime {
				return true
			}
		}
	}

	return false

}

func checkIfNumberIsPrime(number int) bool {
	if number == 2 || number == 3 {
		return true
	} else if number <= 1 || number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}
