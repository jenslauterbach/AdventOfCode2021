package main

import "testing"

func Test_part1(t *testing.T) {
	octopuses := loadOctopuses()
	want := 1655
	got := part1(octopuses)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}
