package main

import "fmt"

func main() {
	coun := []string{"UK",
		"China",
		"USA",
		"France",
		"New Zealand",
		"UK",
		"France"}

	var newCoun = make(map[string]bool)
	output := []string{}

	for _, str := range coun {
		if !newCoun[str] {
			output = append(output, str)
			newCoun[str] = true
		}
	}

	fmt.Println(len(output))
}
