package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	expect := 6
	got := SolvePartTwo(exampleInput)
	if got != expect {
		t.Fatalf("wanted (%v) but got %v", expect, got)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePartTwo(input)
	}
}
