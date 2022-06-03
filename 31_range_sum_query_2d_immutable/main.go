package main

import (
	"fmt"
)

type NumMatrix struct {
	x [][]int
}

func main() {

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)

	result1 := obj.SumRegion(2, 1, 4, 3)
	result2 := obj.SumRegion(1, 1, 2, 2)
	result3 := obj.SumRegion(1, 2, 2, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Constructor(matrix [][]int) (result NumMatrix) {
	result.x = matrix

	return result
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.x[i][j]
		}
	}

	return sum
}
