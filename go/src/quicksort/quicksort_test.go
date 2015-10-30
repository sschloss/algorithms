/*
	quicksort
*/
package quicksort

import (
	"fmt"
	"quicksort"
	"testing"
)

func TestSort(t *testing.T) {

	var a = []int{3, 5, 7, 1, 2, 4, 6}
	fmt.Println("Input: ", a)
	quicksort.Sort(a)
	fmt.Println("Output: ", a)

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Error("Not sorted")
			fmt.Println(a)
		}
	}
}
