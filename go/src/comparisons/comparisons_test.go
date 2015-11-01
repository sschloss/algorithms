/*
	comparisons
	test for comparing pivot choices for the quicksort algorithm
*/
package comparisons

import (
	"comparisons"
	"flag"
	"fmt"
	"os"
	"testing"
	"util"
)

var (
	data []int
)

func TestMain(m *testing.M) {
	flag.Parse()
	data = util.ReadInts("../../../data/QuickSort.txt")
	status := m.Run()
	os.Exit(status)
}

func TestSortFirst(t *testing.T) {
	a := make([]int, len(data))
	copy(a, data)
	count := 0
	comparisons.Sort(a, "first", &count)
	fmt.Println("First Count: ", count)
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Error("Not sorted")
			fmt.Println(a)
		}
	}
}
func TestSortLast(t *testing.T) {
	a := make([]int, len(data))
	copy(a, data)
	count := 0
	comparisons.Sort(a, "last", &count)
	fmt.Println("Last Count: ", count)
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Error("Not sorted")
			fmt.Println(a)
		}
	}
}
func TestSortMedian(t *testing.T) {
	a := make([]int, len(data))
	copy(a, data)
	count := 0
	comparisons.Sort(a, "median", &count)
	fmt.Println("Median Count: ", count)
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Error("Not sorted")
			fmt.Println(a)
		}
	}
}
