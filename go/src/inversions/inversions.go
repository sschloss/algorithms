package inversions

import (
	"math"
)

func MergeSort(data []int, count *int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	mid := getMedian(n)
	left := MergeSort(data[:mid], count)
	right := MergeSort(data[mid:], count)
	return merge(left, right, count)
}

func merge(left, right []int, count *int) []int {
	n := len(left) + len(right)
	result := make([]int, n)
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			*count += len(left[i:])
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
