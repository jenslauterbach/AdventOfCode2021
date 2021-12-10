package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	m := loadHeightMap()
	want := 600
	got := part1(m)
	if got != want {
		t.Errorf("part1(): got = %v, want = %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	m := loadHeightMap()
	want := 987840
	got := part2(m)
	if got != want {
		t.Errorf("part1(): got = %v, want = %v", got, want)
	}
}
