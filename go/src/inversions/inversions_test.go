package inversions

import (
	"fmt"
	"inversions"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var (
	ex1 = []int{5, 3, 1, 4, 2}
	ex2 = []int{8, 1, 3, 6, 5, 4, 2, 7}
)

func TestMergeSort(t *testing.T) {
	r1 := inversions.MergeSort(ex1)
	r2 := inversions.MergeSort(ex2)

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
			t.Errorf("Fail: %q", "Bitches!")
			break
		}
	}

	fmt.Println("IN1: ", ex1)
	fmt.Println("IN2: ", ex2)
	fmt.Println("R1: ", r1)
	fmt.Println("R2: ", r2)

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

func fileIo() {

	dat, err := ioutil.ReadFile("test.txt")
	check(err)
	strarr := strings.Fields(string(dat))
	data := stringsToInts(strarr)
	for i, d := range data {
		fmt.Printf("Integer %d : %d\n", i, d)
	}

}

//func ReadInts(r io.Reader) ([]int, error) {
//	scanner := bufio.NewScanner(r)
//	scanner.Split(bufio.ScanWords)
//	var result []int
//	for scanner.Scan() {
//		x, err := strconv.Atoi(scanner.Text())
//		if err != nil {
//			return result, err
//		}
//		result = append(result, x)
//	}
//	return result, scanner.Err()
//}
