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

	input := convertStringToUint32(userInput)

	fmt.Println(countBits(input))
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

func convertStringToUint32(input string) uint32 {
	number, err := strconv.Atoi(input)

	checkError(err)

	return uint32(number)
}

func countBits(num uint32) (result int32) {
	binaryNum := convertUint32ToBinaryString(num)

	for i := 0; i < len(binaryNum); i++ {
		if string(binaryNum[i]) == "1" {
			result++
		}
	}

	return result
}

func convertUint32ToBinaryString(num uint32) string {
	return strconv.FormatInt(int64(num), 2)
}
