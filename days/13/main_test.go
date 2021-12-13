package main

import "testing"

func Test_part1(t *testing.T) {
	dots, folds := loadPage()

	want := 1
	got := part1(dots, folds)
	if got != want {
		t.Errorf("part1(): got = %d, want = %d", got, want)
	}
}

func Test_foldLeft(t *testing.T) {
	type args struct {
		x              int
		y              int
		foldCoordinate int
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			"ok",
			args{
				x:              622,
				y:              12,
				foldCoordinate: 492,
			},
			362,
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := foldLeft(tt.args.x, tt.args.y, tt.args.foldCoordinate)
			if got != tt.wantX {
				t.Errorf("foldLeft() got = %v, wantX %v", got, tt.wantX)
			}
			if got1 != tt.wantY {
				t.Errorf("foldLeft() got1 = %v, wantX %v", got1, tt.wantY)
			}
		})
	}
}

func Test_foldUp(t *testing.T) {
	type args struct {
		x              int
		y              int
		foldCoordinate int
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			"ok",
			args{
				x:              12,
				y:              542,
				foldCoordinate: 451,
			},
			12,
			360,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := foldUp(tt.args.x, tt.args.y, tt.args.foldCoordinate)
			if gotX != tt.wantX {
				t.Errorf("foldUp() gotX = %v, wantX %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("foldUp() gotY = %v, wantX %v", gotY, tt.wantY)
			}
		})
	}
}
