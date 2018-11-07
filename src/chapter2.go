package main


import (
	"flag"
	"fmt"
	"strings"
)


var n = flag.Bool("n", false, "new line")
var s = flag.String("s", " ", "separator")
func useFlag() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

type receiver func(data int)
