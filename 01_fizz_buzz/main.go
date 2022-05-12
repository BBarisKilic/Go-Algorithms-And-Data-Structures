package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fizzBuzz(n int32) {
	var i int32 = 1

	for ; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func main() {
	input, err := getUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	fizzBuzz(input)
}

func getUserInput() (int32, error) {
	reader := bufio.NewReader(os.Stdin)

	value, _, err := reader.ReadLine()

	if err != nil {
		return 0, err
	}

	num, err := convertInputToInt(string(value))

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
