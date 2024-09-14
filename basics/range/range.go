package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var str string = "Hello"

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i, v := range str {
		fmt.Printf("str[%d] = %c\n", i, v)
	}
}
