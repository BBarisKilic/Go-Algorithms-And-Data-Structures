package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := getUserInput()

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println(firstReverse(input))
}

func getUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter a text: ")

	input, _, err := reader.ReadLine()

	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(input))

	return result, nil
}

func firstReverse(str string) string {

	len := len(str)

	chars := make([]byte, len)

	for i := 0; i < len; i++ {
		chars[i] = str[len-i-1]
	}

	str = string(chars)

	return str
}
