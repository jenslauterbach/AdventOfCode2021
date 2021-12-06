package main

import "testing"

func Test_part1(t *testing.T) {
	fish := loadFish()

	want := 362666
	got := part1(fish)

	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}
