package main

import (
	"fmt"
	"sync"
	"time"
)

var doOnce sync.Once
var group sync.WaitGroup
func once() {
	fmt.Println("once")
}

func makeOnce() {
	doOnce.Do(once)
}

func testOnce() {
	go makeOnce()
	go makeOnce()
	go makeOnce()

	time.Sleep(1 * time.Second)
}

func testGroup() {
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			once()
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("finish")
}
