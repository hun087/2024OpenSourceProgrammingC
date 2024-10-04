package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13              // var i int = 13
	var f float64 = 12.9 // f := 12.9
	c1 := 'Z'            // 90
	c2 := '김'            //44608, home. Unicode

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f * %f\n", i, f, i*f) // invalid operation: i * f (mismatched types int and float64)
	// fmt.Printf("%d * %f * %f\n", i, f, float64(i)*f)   // 강제변환 conversion 실수타입
	fmt.Printf("%d * %f * %d\n", i, f, i*int(f)) // 강제변환 conversion 정수타입
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2)) // i는 바뀌지 않음
}
