package main

import (
	"os"
	"parex/internal/lib"
	"testing"
)

const testfilepath = "N:\\research\\testdata\\gb1.dd"

func Benchmark_Fls(b *testing.B) {
	run(b, 0)
}

func Benchmark_RecursiveFls(b *testing.B) {
	run(b, 1)
}

func Benchmark_RecursiveFcat(b *testing.B) {
	run(b, 2)
}

func run(b *testing.B, level int) {
	tfh, _ := os.Open(testfilepath)
	defer tfh.Close()
	for i := 0; i < b.N; i++ {
		err := lib.Explore(tfh, 0, 1)
		handle(err)
	}
}
