package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)  // 입력받는 방법1
	name, err := in.ReadString('\n') // 입력받는 방법2
	fmt.Printf("Input your name : ")
	fmt.Println(name)
	fmt.Println(err)
}
