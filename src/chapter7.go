package main

import (
	"fmt"
)

type Person struct {
	name string
}

func whichType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case uint8:
		fmt.Println("uint8")
	case int8:
		fmt.Println("int8")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("unknown")
	}
}

func typeAssert() {
	var tpa interface{}

	tpa = Person{}
	var person = tpa.(Person)
	person.name = "xiaowei"

	fmt.Println(person.name)

	whichType(8)
	whichType(float32(8))
	whichType(float64(8))
	whichType(uint8(8))
}

