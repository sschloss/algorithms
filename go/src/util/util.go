package util

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func ReadTable(filename string) [][]string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	lines := regexp.MustCompile(`\n+`).Split(string(dat), -1)
	// strings.Split appends an extra line so remove it
	nLines := len(lines) - 1
	t := make([][]string, nLines)
	for i := 0; i < nLines; i++ {
		line := strings.TrimSpace(lines[i])
		tokens := regexp.MustCompile(`\s+`).Split(line, -1)
		nTokens := len(tokens)
		t[i] = make([]string, nTokens)
		for j := 0; j < nTokens; j++ {
			t[i][j] = strings.TrimSpace(tokens[j])
		}
	}
	return t
}

func ReadIntsTable(filename string) [][]int {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	lines := regexp.MustCompile(`\n+`).Split(string(dat), -1)
	// strings.Split appends an extra line so remove it
	nLines := len(lines) - 1
	t := make([][]int, nLines)
	for i := 0; i < nLines; i++ {
		tokens := regexp.MustCompile(`\s+`).Split(lines[i], -1)
		nTokens := len(tokens) - 1
		t[i] = make([]int, nTokens)
		for j := 0; j < nTokens; j++ {
			token := strings.TrimSpace(tokens[j])
			val, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Format error")
				continue
			}
			t[i][j] = val
		}
	}
	return t
}

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
