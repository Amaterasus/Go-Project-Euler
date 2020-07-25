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


func lcm(lowestNumber, highestNumber int64) int64 {
	numbers := makeRange(lowestNumber, highestNumber)

	var output int64 = 1

	for _, v := range numbers {
		output = checkTimesTables(output, v)
	}

	return output
}


func checkTimesTables(number, divider int64) int64 {
	var i int64 = 1

	flip := checkMultiple(number, divider, i)

	for !flip {
		i++
		flip = checkMultiple(number, divider, i)
	}

	return number * i
}

func checkMultiple(number, divider, i int64) bool {
	return number * i % divider == 0 
}



func makeRange(min, max int64) []int64 {
	a := make([]int64, max-min+1)
	for i := range a {
		a[i] = min + int64(i)
	}
	return a
}