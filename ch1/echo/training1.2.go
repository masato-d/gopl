package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		//s += sep + "[" + string(i) + "]" + " " + arg
		//sep = "\n"
	}
	//fmt.Println(s)
}
