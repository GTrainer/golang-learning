package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("error")
	}
}

func Benchmark_add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
