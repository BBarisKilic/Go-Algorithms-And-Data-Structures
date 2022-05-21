package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	inputFirst := getUserInput(scanner)
	inputSecond := getUserInput(scanner)

	firstArr := converToIntArray(inputFirst)
	secondArr := converToIntArray(inputSecond)

	l1 := []ListNode{}
	l2 := []ListNode{}

	for index, num := range firstArr {
		if index == 0 {
			l1 = append(l1, ListNode{Val: num})
		} else {
			l1 = append(l1, ListNode{Val: num, Next: &l1[index-1]})
		}
	}

	for index, num := range secondArr {
		if index == 0 {
			l2 = append(l2, ListNode{Val: num})
		} else {
			l2 = append(l2, ListNode{Val: num, Next: &l2[index-1]})
		}
	}

	response := addTwoNumbers(&l1[len(firstArr)-1], &l2[len(secondArr)-1])

	result := []int{}

	for {
		result = append(result, response.Val)

		if response.Next == nil {
			break
		}

		response = response.Next
	}

	fmt.Println(result)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numsInl1, numsInl2, smaller, greater := []int{}, []int{}, []int{}, []int{}

	for {
		numsInl1 = append(numsInl1, l1.Val)

		if l1.Next == nil {
			break
		}

		l1 = l1.Next
	}

	for {
		numsInl2 = append(numsInl2, l2.Val)

		if l2.Next == nil {
			break
		}

		l2 = l2.Next
	}

	if len(numsInl1) < len(numsInl2) {
		smaller = numsInl1
		greater = numsInl2
	} else {
		smaller = numsInl2
		greater = numsInl1
	}

	var index = 0

	for ; index < len(smaller); index++ {
		greater[index] = smaller[index] + greater[index]

		if greater[index] >= 10 {
			greater[index] = greater[index] % 10
			if index+1 == len(greater) {
				greater = append(greater, 1)
			} else {
				greater[index+1]++
			}
		}
	}

	for ; index < len(greater); index++ {
		if greater[index] >= 10 {
			greater[index] = greater[index] % 10

			if index+1 == len(greater) {
				greater = append(greater, 1)
			} else {
				greater[index+1]++
			}
		}
	}

	greater = reverseInts(greater)

	lResult := []ListNode{}

	for index, num := range greater {
		if index == 0 {
			lResult = append(lResult, ListNode{Val: num})
		} else {
			lResult = append(lResult, ListNode{Val: num, Next: &lResult[index-1]})
		}
	}

	return &lResult[len(greater)-1]
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
