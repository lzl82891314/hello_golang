package main

import (
	"fmt"
	"os"
)

func main() {
	var msg string = "Hello World By Variable"
	var number int16 = 100
	fmt.Println(msg, number)

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
