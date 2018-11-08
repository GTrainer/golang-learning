package main

import (
	"fmt"
)

const (
	kMonday = 1 << iota
	kTursday
	kThirdsday
)

const (
	kError1 = iota
	kError2 = iota
)

func printConst() {
	fmt.Println(kMonday)
	fmt.Println(kTursday)
	fmt.Println(kThirdsday)
	fmt.Println(kError1)
	fmt.Println(kError2)
}

