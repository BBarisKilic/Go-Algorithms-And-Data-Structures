//! "Greedy Algorithm" is not an appropriate answer to the "Coin Change" problem.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	coinsInput := getUserInput(reader)
	amountInput := getUserInput(reader)

	coins := convertToIntArray(coinsInput)
	amount := convertToInt(amountInput)

	fmt.Println(coinChange(coins, amount))
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

func convertToIntArray(s string) []int {
	nums := strings.Fields(s)

	intArray := []int{}

	for _, num := range nums {
		intNum, err := strconv.Atoi(num)

		checkError(err)

		intArray = append(intArray, intNum)
	}

	return intArray
}

func convertToInt(s string) int {
	num, err := strconv.Atoi(s)

	checkError(err)

	return num
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	var result = -1

	for index := 1; index < len(coins)+1; index++ {
		result = calculate(coins, amount, len(coins)-index, 0)

		if result != -1 {
			break
		}
	}

	return result
}

func calculate(coins []int, amount int, index int, result int) int {

	if index < 0 {
		return -1
	}

	times := amount / coins[index]
	rest := amount % coins[index]

	result += times

	if rest == 0 {
		return result
	}

	amount -= coins[index] * times
	index--

	return calculate(coins, amount, index, result)
}
