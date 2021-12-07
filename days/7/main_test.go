package main

import "testing"

func Test_part1(t *testing.T) {
	crabs := loadCrabs()
	want := 347449
	got := part1(crabs)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	crabs := loadCrabs()
	want := 98039527
	got := part2(crabs)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}
