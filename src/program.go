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


func main() {
	fmt.Println("Hello, golang!");

	printArgs()
}

