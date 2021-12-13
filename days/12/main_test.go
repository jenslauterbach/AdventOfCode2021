package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	caves := loadCaves()
	want := 5157
	got := part1(caves)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	caves := loadCaves()
	want := 144309
	got := part2(caves)
	if got != want {
		t.Errorf("part2(): got = %d, want = %d", got, want)
	}
}
