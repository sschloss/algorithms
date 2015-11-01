package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInts(filename string) []int {
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
