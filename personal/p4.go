package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	// var output [][]string

	for i := 0; i < len(input); i++ {
		ss := sortString(input[i])
		fmt.Println(ss)
	}
}

func sortString(str string) string {
	r := []rune(str)

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
