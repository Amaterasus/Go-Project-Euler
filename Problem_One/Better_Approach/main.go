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
	threes := sumOfMultiples(3, highestNumber)
	fives := sumOfMultiples(5, highestNumber)
	fifteens := sumOfMultiples(15, highestNumber)
	return threes + fives - fifteens
}

func sumOfMultiples(multiple, highestNumber int) int {

	sum := 0

	for i := multiple; i < highestNumber; i += multiple {
		sum += i
	}
	return sum
}