package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for range numbers {

		rand.Seed(time.Now().UnixNano())
		n := len(numbers)
		if n <= 1 {
			return
		}

		for i := n - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
