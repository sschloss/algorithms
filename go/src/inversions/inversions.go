package inversions

import (
	"math"
)

func MergeSort(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	mid := getMedian(n)
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	n := len(left) + len(right)
	result := make([]int, n)
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func getMedian(n int) int {
	return int(math.Floor(float64(n) / float64(2)))
}
