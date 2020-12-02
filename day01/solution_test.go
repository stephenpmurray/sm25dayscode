package day01

import (
	"reflect"
	"testing"
)

func Test_importInput(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []int
		wantErr    bool
	}{
		{
			name:       "successfully read contents of test file",
			args:       args{"testinput"},
			wantValues: []int{15, 1112, 2000, 20, 114, 7},
			wantErr:    false,
		},
		{
			name:       "return error when failing to open fstream",
			args:       args{"bad_filename"},
			wantValues: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := importInput(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("importInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("importInput() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func Test_checkList(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name       string
		args       args
		wantSummed [2]int
		wantErr    bool
	}{
		{
			name:       "return correct value pair",
			args:       args{[]int{15, 1112, 2000, 20, 114, 7}},
			wantSummed: [2]int{2000, 20},
			wantErr:    false,
		},
		{
			"returns error when a matching pair can't be found",
			args{[]int{}},
			[2]int{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSummed, err := checkList(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSummed, tt.wantSummed) {
				t.Errorf("checkList() = %v, want %v", gotSummed, tt.wantSummed)
			}
		})
	}
}
