package main

import (
	"fmt"
)

// zero values

func main() {
	var f float64
	var i int
	var b bool
	var s string

	my_school_account := 5
	println(my_school_account)

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i) // zero values
	f = 2.7
	i = 3
	fmt.Print("\n\n", f <= float64(i), "\n") // comparison (true/false)

}
