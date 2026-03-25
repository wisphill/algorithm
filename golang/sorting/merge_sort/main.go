package main

import "fmt"

func merge_sort(arr []int, startIndex int, endIndex int) []int {
	if startIndex == endIndex {
		return []int{arr[startIndex]}
	}

	middleIndex := (startIndex + endIndex) / 2
	sorted_arr1 := merge_sort(arr, startIndex, middleIndex)
	sorted_arr2 := merge_sort(arr, middleIndex+1, endIndex)
	res := merge_sorted_array(sorted_arr1, sorted_arr2)
	return res
}

func merge_sorted_array(sorted_arr1 []int, sorted_arr2 []int) []int {
	res := make([]int, 0, len(sorted_arr1)+len(sorted_arr2))
	i, j := 0, 0
	for i < len(sorted_arr1) && j < len(sorted_arr2) {
		if sorted_arr1[i] <= sorted_arr2[j] {
			res = append(res, sorted_arr1[i])
			i++
		} else {
			res = append(res, sorted_arr2[j])
			j++
		}
	}

	res = append(res, sorted_arr1[i:]...)
	res = append(res, sorted_arr2[j:]...)
	return res
}

func main() {
	arr := []int{1, 3, 52, 3, 2, 7, 8, 0, 1, 4, 3, 2, 9}
	sorted_arr := merge_sort([]int{1, 3, 52, 3, 2, 7, 8, 0, 1, 4, 3, 2, 9}, 0, len(arr)-1)
	fmt.Println("sorted array: ", sorted_arr)
}
