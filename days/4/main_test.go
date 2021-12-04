package main

import "testing"

func Test_part1(t *testing.T) {
	numbers, boards := loadGameData()
	want := 74320
	got := part1(numbers, boards)
	if got != want {
		t.Errorf("part1(), got = %d, want = %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	numbers, boards := loadGameData()
	want := 17884
	got := part2(numbers, boards)
	if got != want {
		t.Errorf("part1(), got = %d, want = %d", got, want)
	}
}
