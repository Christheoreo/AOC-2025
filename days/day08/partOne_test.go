package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expect := 40
	got := SolvePartOne(exampleInput, 10)
	if got != expect {
		t.Fatalf("wanted (%v) but got %v", expect, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePartOne(input, 1000)
	}
}
