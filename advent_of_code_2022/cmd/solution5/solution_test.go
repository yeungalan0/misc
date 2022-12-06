package main

import (
	"testing"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"
)

var testList = utils.ReadFileLinesToSlice("input_test")

func TestSolve1(t *testing.T) {
	want := "CMZ"
	got, err := solve1(testList)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	want := "MCD"
	got, err := solve2(testList)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
