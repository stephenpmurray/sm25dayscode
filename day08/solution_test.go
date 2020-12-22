package day08

import (
	"reflect"
	"testing"
)

func TestPuzzle_ProcInput(t *testing.T) {
	type fields struct {
		Cmds []Cmd
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCmds []Cmd
		wantErr  bool
	}{
		{
			name: "return error when line does not match regex pattern",
			args: args{[]string{
				"asdf",
			}},
			wantErr: true,
		},
		{
			name: "return error when string cannot convert to int",
			args: args{[]string{
				"JJJ +asd",
			}},
			wantErr: true,
		},
		{
			name: "Correctly parse input",
			args: args{
				[]string{
					"acc -5",
					"nop +333",
					"acc +45",
				},
			},
			wantCmds: []Cmd{
				{"acc", -5},
				{"nop", 333},
				{"acc", 45},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Cmds: tt.fields.Cmds,
			}
			if err := p.ProcInput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.ProcInput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantCmds, p.Cmds) {
				t.Errorf("Puzzle.ProcInput() want:\n%v got\n%v", tt.wantCmds, p.Cmds)
			}
		})
	}
}

func TestPuzzle_AnsOne(t *testing.T) {
	type fields struct {
		Cmds []Cmd
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"return correct acc value",
			fields{
				Cmds: []Cmd{
					{"nop", 0},
					{"acc", 1},
					{"jmp", 4},
					{"acc", 3},
					{"jmp", -3},
					{"acc", -99},
					{"acc", 1},
					{"jmp", -4},
					{"acc", 6},
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Cmds: tt.fields.Cmds,
			}
			if got := p.AnsOne(); got != tt.want {
				t.Errorf("Puzzle.AnsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
