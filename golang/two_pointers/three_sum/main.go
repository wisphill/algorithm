package main

import "fmt"

func three_sum(arr []int, total int) [][]int {
	res := make([][]int, 0)

	for index, value := range arr {
		remaining_total := total - value
		two_sum_arrays := two_sum(append(arr[:index], arr[index+1:]...), remaining_total)
		if len(two_sum_arrays) > 0 {
			for _, array := range two_sum_arrays {
				res = append(res, []int{value, array[0], array[1]})
			}
		}
	}

	return res
}

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
	triple_values := three_sum([]int{1, 2, 3, 8, 13, 3, 3, 2, 1}, 6)
	for _, values := range triple_values {
		fmt.Println(values[0], values[1], values[2])
	}
}
