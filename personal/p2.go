package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4, 5}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Printf("(%d,%d)\n", a[i], b[j])
		}
	}
	// output := (1,3), (1,4), (2,3) (2,4)
}
