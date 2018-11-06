package main


import (
	"fmt"
	"os"
)


func printArgs() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}	
}

func printArgsRange() {
	for _, arg = range os.Args[1:] {
		fmt.Println(arg)		
	}
}


func main() {
	fmt.Println("Hello, golang!");

	printArgs()
	printArgsRange()
}

