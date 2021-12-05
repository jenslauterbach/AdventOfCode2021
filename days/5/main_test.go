package main

import "testing"

func Test_part1(t *testing.T) {
	lines := loadLines()
	want := 6841
	if got := part1(lines); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	lines := loadLines()
	want := 19258
	if got := part2(lines); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
