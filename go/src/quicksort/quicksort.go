/*
	quicksort
*/
package quicksort

func Sort(a []int) {
	if len(a) < 2 {
		return
	}
	i := partition(a)
	Sort(a[:i])
	Sort(a[i+1:])
}

func partition(a []int) int {
	if len(a) < 2 {
		return 0
	}
	// pick the first element for the pivot for now
	i, p := 0, a[0]
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
	return (i - 1)
}
