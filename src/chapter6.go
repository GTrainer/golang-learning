package main

import (
	"fmt"
)

type Point struct {
	x, y float32
}

func (p *Point) scale(sca float32) {
	p.x *= sca;
	p.y *= sca;
}

func (p *Point) print() {
	fmt.Println(p.x)
	fmt.Println(p.y)
}

func makeReceiver() {
	p := &Point { 1.0, 1.0 }
	p.scale(2.0)
	p.print()
}
