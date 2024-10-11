package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	fmt.Printf("Input your name : ")
	fmt.Println(name)
	fmt.Println(err)
}
