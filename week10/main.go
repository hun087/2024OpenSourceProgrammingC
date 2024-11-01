package main

import (
	"fmt"
)

func main() {
	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 {
		isPrime = false
	} else {
		j := 3
		// for j <= int(math.Sqrt(float64(n))) {
		for j*j <= n {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j)
			j = j + 2
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
