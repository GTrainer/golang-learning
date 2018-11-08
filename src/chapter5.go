package main

import (
	"fmt"
) 

func makePanic(x, y int) {
	if y == 0 {
		panic("y is 0")
	}
}

func makeRecover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	makePanic(0, 0)
	return nil
}

func makeRun() {
	fmt.Println("start to run ...")
	makeRecover()
	fmt.Println("continue to run ...")
}
