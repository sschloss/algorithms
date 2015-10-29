package inversions

import (
	"fmt"
	"inversions"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestCountInversions(t *testing.T) {

	ex1 := readFile("../../../data/IntegerArray.txt")
	count := 0
	inversions.MergeSort(ex1, &count)

	if count != 2407905288 {
		t.Error("Fail: %q", count)
	}
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
