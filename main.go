package main

import (
	"fmt"
	"time"
)

func calculate(t int, l int, index int, sum int) int {
	l += 1

	for i := 1; i <= index; i++ {
		if l < t {
			sum = calculate(t, l, i, sum)
		} else {
			sum += i
		}
	}

	return sum
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("Calculating cardinality for tuple of size %d\n", i)

		startTime := time.Now()

		res := calculate(i, 1, i, 0)

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)

		fmt.Printf("Card:  %d\n", res)
		fmt.Printf("Execution time: %d nanosecond\n\n", executionTime)
	}
}
