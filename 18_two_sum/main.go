package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	inputNums := getUserInput(scanner)
	inputTarget := getUserInput(scanner)

	nums := converToIntArray(inputNums)
	target := convertToInt(inputTarget)

	fmt.Println(twoSum(nums, target))
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

func converToIntArray(s string) []int {
	nums := strings.Fields(s)

	numArr := []int{}

	for _, num := range nums {
		intNum, err := strconv.Atoi(num)

		checkError(err)

		numArr = append(numArr, intNum)
	}

	return numArr
}

func convertToInt(s string) int {
	num, err := strconv.Atoi(s)

	checkError(err)

	return num
}

func twoSum(nums []int, target int) []int {

	indexOne, indexTwo := calculate(nums, target)

	return []int{indexOne, indexTwo}
}

func calculate(nums []int, target int) (int, int) {
	for mainIndex, num := range nums {
		for secondaryIndex := len(nums) - 1; secondaryIndex > mainIndex; secondaryIndex-- {
			if nums[secondaryIndex] == target-num {
				return mainIndex, secondaryIndex
			}
		}
	}

	return 0, 0
}
