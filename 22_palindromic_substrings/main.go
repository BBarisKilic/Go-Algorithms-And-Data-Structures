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

	fmt.Println(countSubstrings(userInput))
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

func countSubstrings(s string) int {
	totalPalindromic := 0

	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			reverseS := reverseString(s[i:j])

			if s[i:j] == reverseS && reverseS != "" {
				totalPalindromic += 1
			}
		}
	}

	return totalPalindromic
}

func reverseString(s string) string {
	chars := []byte(s)

	i := 0
	j := len(s) - 1

	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}

	return string(chars)
}
