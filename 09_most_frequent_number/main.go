package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6, 4, 4, 5, 5, 5, 5}

	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if m[arr[i]] == 0 {
			m[arr[i]] = 1
		} else {
			m[arr[i]] += 1
		}
	}

	var mostFrequentNumber, times int

	for value, key := range m {
		if key > times {
			times = key
			mostFrequentNumber = value
		}
	}

	fmt.Println(mostFrequentNumber)
}
