package main

import "testing"

func Test_part1(t *testing.T) {
	report := loadReport()

	var want uint64 = 3429254
	got := part1(report)
	if got != want {
		t.Errorf("part1(), got = %d, want = %d", got, want)
	}
}
