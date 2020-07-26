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
	
	i := highestNumber
	for true {
		if bruteForce(i, numbers) {
			break
		}
		i += highestNumber
	}
	return i
}

func bruteForce(testNumber int, numberRange []int) bool {

	for _ , v := range numberRange {
		if !divisible(testNumber, v) {
			return false
		}
	}
	return true
}

func divisible(testNumber, divider int) bool {
	if (testNumber % divider) == 0 {
		return true
	}
	return false
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}