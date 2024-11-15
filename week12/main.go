package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	fmt.Println(scores[1], scores[0]) //, scores[3])  // invalid argument: index 3 out of bounds
	fmt.Printf("%#v\n", scores)       // 배열 리터럴 형태로 보여줌
	fmt.Println(scores)

	// var dates [3]time.Time
	// dates[1] = time.Unix(1947200001, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1947200001, 0), // need comma
	// } // 원소 3개 배열 리터럴로 초기화 후  실행
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947200001, 0)}
	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Printf("%#v\n", dates)
	fmt.Println(dates)
}
