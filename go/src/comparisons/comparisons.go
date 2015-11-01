/*
	comparisons
	comparing pivot choices for the quicksort algorithm
*/
package comparisons

import (
	"math"
)

type PTYPE string

func Sort(a []int, pt PTYPE, count *int) {
	if len(a) < 2 {
		return
	}
	i := 0
	switch pt {
	case "median":
		i = partition(a, median, count)
	case "last":
		i = partition(a, last, count)
	default:
		i = partition(a, first, count)
	}

	Sort(a[:i], pt, count)
	Sort(a[i+1:], pt, count)
}

func partition(a []int, pf func(a []int) int, count *int) int {
	if len(a) < 2 {
		return 0
	}

	*count += len(a) - 1

	// select pivot
	pi := pf(a)
	// save the pivot value
	p := a[pi]
	// save value in first position
	x := a[0]
	// put the pivot in the first position
	a[0] = p
	// put the saved first element at the pivot index
	a[pi] = x

	i, j := 1, 1
	for j < len(a) {
		if p > a[j] {
			tmp := a[i]
			a[i] = a[j]
			a[j] = tmp
			i++
		}
		j++
	}
	tmp := a[0]
	a[0] = a[i-1]
	a[i-1] = tmp
	return i - 1
}

func first(a []int) int {
	return 0
}

func last(a []int) int {
	return len(a) - 1
}

func median(a []int) int {
	if len(a) < 2 {
		return 0
	}
	start, end := 0, len(a)-1
	mid := int(math.Floor(float64(start+end) / float64(2)))

	var parr = []int{a[start], a[mid], a[end]}
	insertionSort(parr)
	m := 0
	switch parr[1] {
	case a[start]:
		m = start
	case a[mid]:
		m = mid
	default:
		m = end
	}

	return m
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			tmp := a[i]
			j := i
			for j > 0 && a[j-1] > a[j] {
				a[j] = a[j-1]
				a[j-1] = tmp
				j--
			}
			a[j] = tmp
		}
	}
}
