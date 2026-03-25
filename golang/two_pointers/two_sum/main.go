package main

import "fmt"

func two_sum(arr []int, total int) [][]int {
	// dictionay
	dict := make(map[int]int)
	res := make([][]int, 0)
	for _, value := range arr {
		remaining_value := total - value
		// mark the value is existed in the array
		dict[value] = 1

		if _, found := dict[remaining_value]; found {
			res = append(res, []int{value, remaining_value})
		}
	}

	return res
}

func main() {
	couple := two_sum([]int{1, 2, 3, 8, 13, 3, 3, 2, 1, 4}, 5)
	for _, values := range couple {
		fmt.Println(values[0], values[1])
	}
}
