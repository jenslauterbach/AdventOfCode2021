package main

import (
	"reflect"
	"testing"
)

func Test_parsePatterns(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"ok",
			args{line: "ceg gedcfb ec eabfdg gcdabe baged cabgf gbaec fecagdb eacd | efcgbad adfecbg gec abgce"},
			[]string{"ceg", "gedcfb", "ec", "eabfdg", "gcdabe", "baged", "cabgf", "gbaec", "fecagdb", "eacd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePatterns(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOutput(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"ok",
			args{line: "ceg gedcfb ec eabfdg gcdabe baged cabgf gbaec fecagdb eacd | efcgbad adfecbg gec abgce"},
			[]string{"efcgbad", "adfecbg", "gec", "abgce"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOutput(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_difference(t *testing.T) {
	type args struct {
		a []rune
		b []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			"ok",
			args{
				a: []rune{'a', 'b', 'c'},
				b: []rune{'b', 'c'},
			},
			[]rune{'a'},
		},
		{
			"ok",
			args{
				a: []rune{'a', 'b', 'c'},
				b: []rune{'c'},
			},
			[]rune{'a', 'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		a []rune
		b [][]rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			"ok",
			args{
				a: []rune{'a', 'b', 'c', 'd'},
				b: [][]rune{
					{'a', 'c'},
					{'d'},
				},
			},
			[]rune{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.args.a, tt.args.b...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
