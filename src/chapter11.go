package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func testConvert() {
	var _10 = "10"

	a, _ := strconv.ParseInt(_10, 2, 32)
	b, _ := strconv.ParseInt(_10, 8, 32)
	c, _ := strconv.ParseInt(_10, 10, 32)
	d, _ := strconv.ParseInt(_10, 16, 32)

	fmt.Println(a, " ", b, " ", c, " ", d)

	_10 = strconv.Quote(_10)
	fmt.Println(_10)
}

func testSwitch() {
	switch 1 {
	case 1, 2:
		fmt.Println("1")
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
}

func testDefer() {
	var x = 1
	defer func () {
		fmt.Println(x)
	}()
	defer func () {
		fmt.Println(x + 1)
	}()

	x++
	
	defer func () {
		fmt.Println(x + 2)
	}()
}

func testString() {
	var s1 string = "你好"
	fmt.Println(len(s1))

	var s2 string = "Hello"
	fmt.Println(len(s2))
}

type Cat struct {
}

func (p Cat) eat() {
	fmt.Println("Cat::eat")
}

func (p *Cat) walk() {
	fmt.Println("Cat::walk")
}

func testFunc() {
	cat := &Cat{}
	cat.eat()
	cat.walk()

	eat := Cat.eat
	eat(*cat)

	walk := (*Cat).walk
	walk(cat)
}

func testGoExit() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			if x == 5 {
				runtime.Goexit()
			}
			fmt.Println(x)
		}(i)
	}

	time.Sleep(time.Second)
}
