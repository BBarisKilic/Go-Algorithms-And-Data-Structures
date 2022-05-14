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
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(firstFactorial(input))
}

func getUserInput(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	num, err := convertInputToInt(input)

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

func firstFactorial(num int) int {
	for i := num - 1; i > 0; i-- {
		num = num * i
	}

	return num
}
