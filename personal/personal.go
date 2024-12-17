package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// "fn453eru234wery234v32" =>> 453+234+234+32 // output := 953
func main() {
	str := "fn453eru234wer234yv33dsfergerf1e1r"

	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(str, -1)

	sum := 0
	// Convert each number string to an integer and add it to the sum
	for _, num := range numbers {
		val, _ := strconv.Atoi(num)
		sum += val
	}

	fmt.Printf("The sum of numbers in the string is: %d\n", sum)

}
