/*
	mergesort
*/
package mergesort

import (
	"fmt"
	"io/ioutil"
	"mergesort"
	"strconv"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {

	ex1 := readFile("../../../data/test.txt")
	ex2 := readFile("../../../data/test2.txt")
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

func readFile(filename string) []int {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	strarr := strings.Fields(string(dat))
	data := stringsToInts(strarr)
	return data
}

func stringsToInts(strarr []string) []int {
	var data []int
	for _, str := range strarr {
		d, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		} else {
			data = append(data, d)
		}
	}
	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
