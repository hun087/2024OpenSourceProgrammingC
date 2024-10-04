package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 13
	f := 12.9
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f * %f\n", i, f, i*f) //.\main.go:14:37: invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f * %f\n", i, f, float64(i)*f) // 강제변환 conversion
	fmt.Println(reflect.TypeOf(i))
}
