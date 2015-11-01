/*
	mergesort
*/
package mergesort

import (
	"fmt"
	"mergesort"
	"testing"
	"util"
)

func TestSort(t *testing.T) {

	ex1 := util.ReadInts("../../../data/test.txt")
	ex2 := util.ReadInts("../../../data/test2.txt")
	r1 := mergesort.Sort(ex1)
	r2 := mergesort.Sort(ex2)
	fail := false
	data := [][]int{r1, r2}
	for _, ex := range data {
		for i := 1; i < len(ex); i++ {
			if ex[i-1] > ex[i] {
				fail = true
				break
			}
		}
		if fail {
			t.Error("Fail: Error: Not sorted")
			break
		}
	}

	fmt.Println("IN1: ", ex1)
	fmt.Println("IN2: ", ex2)
	fmt.Println("R1: ", r1)
	fmt.Println("R2: ", r2)
}
