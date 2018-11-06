package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func printArgs() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}	
}

func printArgsRange() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)		
	}
}

func printArgsJoin() {
	fmt.Println(strings.Join(os.Args[1:], ";"))
}

func dup() {
	counter := make(map[string]int)
	for _, filename := range os.Args[1:] {
		text, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("open ", filename, " failed!")
			continue;
		}
		for _, line := range strings.Split(string(text), "\n") {
			counter[line]++
		}
	}
	for line, num := range counter {
		if num > 1 {
			fmt.Println(line, " dup!")
		}
	}
}

func main() {
	fmt.Println("Hello, golang!");

	//printArgs()
	//printArgsRange()
	//printArgsJoin()
	dup()
}

