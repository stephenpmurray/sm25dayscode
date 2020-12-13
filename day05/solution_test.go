package day05

import (
	"reflect"
	"testing"
)

func TestPuzzle_Input(t *testing.T) {
	type fields struct {
		Passes []Pass
	}
	type args struct {
		FileName string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantPasses []string
		wantErr    bool
	}{
		{
			name:       "return correct string from a test file",
			args:       args{"testinput"},
			wantPasses: []string{"FFFFFFFRRR", "BBBBBBBLLL", "BFFFFFFRRR"},
			wantErr:    false,
		},
		{
			name:       "return error when filename is invalid",
			args:       args{"asdfasdf"},
			wantPasses: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Passes: tt.fields.Passes,
			}
			gotPasses, err := p.Input(tt.args.FileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.Input() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPasses, tt.wantPasses) {
				t.Errorf("Puzzle.Input() = %v, want %v", gotPasses, tt.wantPasses)
			}
		})
	}
}

func TestPuzzle_ReadPassports(t *testing.T) {
	type fields struct {
		Passes []Pass
	}
	type args struct {
		s []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantPass []Pass
		wantErr  bool
	}{
		{
			name:     "read string array correctly to Pass struct",
			args:     args{[]string{"FFFFFFFRRR", "BBBBBBBLLL", "BFFFFFFRRR"}},
			wantPass: []Pass{{Row: 0, Col: 7}, {Row: 127, Col: 0}, {Row: 64, Col: 7}},
			wantErr:  false,
		},
		{
			name:    "return error when string in input slice is empty",
			args:    args{[]string{""}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Passes: tt.fields.Passes,
			}

			if err := p.ReadPassports(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.ReadPassports() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.wantPass, p.Passes) {
				t.Errorf("Puzzle.ReadPassports() got %v want %v", p.Passes, tt.wantPass)
			}
		})
	}
}

func TestPuzzle_AnsOne(t *testing.T) {
	type fields struct {
		Passes []Pass
	}
	tests := []struct {
		name    string
		fields  fields
		wantMax int
		wantErr bool
	}{
		{
			name:    "return error when Passes[] is empty slice",
			wantMax: 0,
			wantErr: true,
		},
		{
			name:    "return correct max value",
			fields:  fields{[]Pass{{Row: 0, Col: 7}, {Row: 127, Col: 0}, {Row: 64, Col: 7}}},
			wantMax: 448,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Passes: tt.fields.Passes,
			}
			gotMax, err := p.AnsOne()
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.ansOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMax != tt.wantMax {
				t.Errorf("Puzzle.ansOne() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
