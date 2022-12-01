package day1

import (
	"testing"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"
)

var testList = utils.ReadFileLinesToSlice("../../config/day1/input_test")

func TestSolve1(t *testing.T) {
	want := 24000
	got, err := Solve1(testList)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	want := 45000
	got, err := Solve2(testList)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
