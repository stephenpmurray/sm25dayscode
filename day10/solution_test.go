package day10

import (
	"testing"
)

func TestAnsOne(t *testing.T) {
	type args struct {
		sorted []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
		wantErr bool
	}{
		{
			"return correct answer from a sorted int slice",
			args{
				[]int{1, 1, 2, 3, 4, 7, 10, 12, 12, 13}, // 5 * 3
			},
			15,
			false,
		},
		{
			"return error when diff > 3",
			args{
				[]int{1, 2, 4, 7, 10, 14},
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns, err := AnsOne(tt.args.sorted)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnsOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAns != tt.wantAns {
				t.Errorf("AnsOne() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func TestcalcPaths(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"return p=1 for sequence of n=1",
			args{1},
			1,
		},
		{
			"return p=1 for sequence of n=2",
			args{2},
			1,
		},
		{
			"return p=2 for sequence of n=3",
			args{3},
			2,
		},
		{
			"return p=4 for sequence of n=4",
			args{4},
			4,
		},
		{
			"return p=13 for sequence of n=6",
			args{6},
			13,
		},
		{
			"return p=24 for sequence of n=7",
			args{7},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcPaths(tt.args.n); got != tt.want {
				t.Errorf("CalcPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnsTwo(t *testing.T) {
	type args struct {
		sorted []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"calculate correct product",
			args{[]int{
				1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19},
			},
			8,
		},

		{
			"calculate product for very short run",
			args{
				[]int{1, 2, 3, 4}, // zero is appended to sequence, so n=5
			},
			7,
		},
		{
			"calculate correct product for longer sequence",
			args{[]int{
				1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20,
				23, 24, 25, 28, 31, 32, 33, 34, 35, 38,
				39, 42, 45, 46, 47, 48, 49},
			},
			19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnsTwo(tt.args.sorted); got != tt.want {
				t.Errorf("AnsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
