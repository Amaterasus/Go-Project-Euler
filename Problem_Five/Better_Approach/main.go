package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	answer := lcm(1, 20)
	t := time.Now()
	fmt.Println("The answer is", answer)


	elapsed := t.Sub(start)	
	fmt.Println("Time taken", elapsed)
}


func lcm(lowestNumber, highestNumber int) int {
	numbers := makeRange(lowestNumber, highestNumber)

	output := 1

	for _, v := range numbers {
		output = checkTimesTables(output, v)
	}

	return output
}


func checkTimesTables(number, divider int) int {
	i := 1

	flip := checkMultiple(number, divider, i)

	for !flip {
		i++
		flip = checkMultiple(number, divider, i)
	}

	return number * i
}

func checkMultiple(number, divider, i int) bool {
	return number * i % divider == 0 
}



func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}