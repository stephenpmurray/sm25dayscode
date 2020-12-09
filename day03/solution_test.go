package day03

import (
	"reflect"
	"testing"
)

func TestPuzzle_readTreeLine(t *testing.T) {
	type fields struct {
		Map [][]rune
	}
	type args struct {
		line string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValues [][]rune
	}{
		{
			name:       "return success when line is appended to map array",
			args:       args{"..#"},
			wantValues: [][]rune{{'.', '.', '#'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Map: tt.fields.Map,
			}
			p.readTreeLine(tt.args.line)
			if !reflect.DeepEqual(p.Map, tt.wantValues) {
				t.Errorf("p.Map = %v \n want %v", p.Map, tt.wantValues)
			}
		})
	}
}

func TestPuzzle_Input(t *testing.T) {
	type fields struct {
		Map [][]rune
	}
	type args struct {
		FileName string
		ProcLine func(string)
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValues [][]rune
		wantErr    bool
	}{
		{
			name:    "return error when filename is bad",
			args:    args{FileName: "bad_filename"},
			wantErr: true,
		},
		{
			name:       "return success when file is read correctly",
			args:       args{"testinput", p.readTreeLine()}, // TODO: Mock the struct and method
			wantValues: [][]rune{{'#', '#', '.', '.'}, {'.', '#', '.', '#'}},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Map: tt.fields.Map,
			}
			if err := p.Input(tt.args.FileName, tt.args.ProcLine); (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.Input() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p.Map, tt.wantValues) {
				t.Errorf("p.Map = %v \n want %v", p.Map, tt.wantValues)
			}
		})
	}
}
