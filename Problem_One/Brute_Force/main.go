package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	answer := sumOfThreeAndFiveMultiples(1000)
	t := time.Now()
	fmt.Println("The answer is", answer)

	elapsed := t.Sub(start)
	fmt.Println("Time taken", elapsed)
}

func sumOfThreeAndFiveMultiples(highestNumber int) int {
	numbers := makeRange(1, highestNumber - 1)
	sum := 0

	for _, v := range numbers {
		if v % 3 == 0 || v % 5 == 0 {
			sum += v
		}
	}

	return sum
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
