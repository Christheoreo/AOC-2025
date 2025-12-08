package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expect := 21
	got := SolvePartOne(exampleInput)
	if got != expect {
		t.Fatalf("wanted (%v) but got %v", expect, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePartOne(input)
	}
}
