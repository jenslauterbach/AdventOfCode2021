package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	lines := loadLines()

	want := 323691
	got := part1(lines)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	lines := loadLines()

	want := 2858785164
	got := part2(lines)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}
