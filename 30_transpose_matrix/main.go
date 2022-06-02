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

	matrix := convertInputToMatrix(userInput)

	fmt.Println(transpose(matrix))

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

func convertInputToMatrix(input string) (matrix [][]int) {
	inputArr := strings.Fields(input)
	rowArr := []int{}

	var index = 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num, err := strconv.Atoi(inputArr[index])

			checkError(err)

			rowArr = append(rowArr, num)

			index++
		}

		matrix = append(matrix, [][]int{rowArr}...)
		rowArr = nil
	}

	return matrix
}

func transpose(matrix [][]int) (result [][]int) {
	tempArr := []int{}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			tempArr = append(tempArr, matrix[j][i])
		}
		result = append(result, [][]int{tempArr}...)
		tempArr = nil
	}

	return result
}
