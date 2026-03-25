package main

func bubble_sort(arr []int) []int {
	res := arr
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i] = res[i] + res[j]
				res[j] = res[i] - res[j]
				res[i] = res[i] - res[j]
			}
		}
	}

	return res
}

func main() {
	sorted_arr := bubble_sort([]int{1, 3, 52, 3, 2, 7, 8, 0, 1, 4, 3, 2, 9})
	for _, value := range sorted_arr {
		println(value)
	}
}
