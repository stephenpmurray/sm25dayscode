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

func TestPuzzle_CheckRun(t *testing.T) {
	MockCmds := []Cmd{
		{"nop", 0},
		{"acc", 1},
		{"jmp", 4},
		{"acc", 3},
		{"jmp", -3},
		{"acc", -99},
		{"acc", 1},
		{"jmp", -4},
		{"acc", 6},
	}
	type fields struct {
		Cmds []Cmd
	}
	tests := []struct {
		name     string
		fields   fields
		wantAcc  int
		wantRan  map[*Cmd]bool
		wantLine int
	}{
		{
			name: "return correct values fir standard case where boot seq loops",
			fields: fields{
				Cmds: MockCmds,
			},
			wantAcc: 5,
			wantRan: map[*Cmd]bool{
				&MockCmds[0]: true,
				&MockCmds[1]: true,
				&MockCmds[2]: true,
				&MockCmds[3]: true,
				&MockCmds[4]: true,
				&MockCmds[6]: true,
				&MockCmds[7]: true,
			},
			wantLine: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Cmds: tt.fields.Cmds,
			}
			acc, ran, line := p.CheckRun()
			if acc != tt.wantAcc {
				t.Errorf("Puzzle.CheckRun() acc = %v, want %v", acc, tt.wantAcc)
			}
			if !reflect.DeepEqual(ran, tt.wantRan) {
				t.Errorf("Puzzle.CheckRun() ran = %v, want %v", ran, tt.wantRan)
			}
			if line != tt.wantLine {
				t.Errorf("Puzzle.CheckRun() line = %v, want %v", line, tt.wantLine)
			}
		})
	}
}

func TestPuzzle_flipCmd(t *testing.T) {
	MockCmds := []Cmd{
		{"nop", 0},
		{"acc", 1},
		{"jmp", 4},
		{"acc", 3},
		{"jmp", -3},
		{"acc", -99},
		{"acc", 1},
		{"jmp", -4},
		{"acc", 6},
	}
	type fields struct {
		Cmds []Cmd
	}
	type args struct {
		line int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantCmd Cmd
	}{
		{
			"flip command from nop to jmp",
			fields{
				Cmds: MockCmds,
			},
			args{
				0,
			},
			Cmd{"jmp", 0},
		},
		{
			"flip command from jmp to nop",
			fields{
				Cmds: MockCmds,
			},
			args{
				7,
			},
			Cmd{"nop", -4},
		},
		{
			"don't flip command when acc",
			fields{
				Cmds: MockCmds,
			},
			args{
				3,
			},
			Cmd{"acc", 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Cmds: tt.fields.Cmds,
			}
			p.flipCmd(tt.args.line)
			if !reflect.DeepEqual(tt.wantCmd, p.Cmds[tt.args.line]) {
				t.Errorf("flipCmd() got %v want %v", tt.wantCmd, p.Cmds[tt.args.line])
			}
		})
	}
}

func TestPuzzle_AnsTwo(t *testing.T) {
	type fields struct {
		Cmds []Cmd
	}
	tests := []struct {
		name     string
		fields   fields
		wantEnds int
		wantAcc  int
		wantErr  bool
	}{
		{
			name: "return correct acc and line to flip",
			fields: fields{
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
			wantEnds: 8,
			wantAcc:  8,
			wantErr:  false,
		},
		{
			name: "return error when there is no solution",
			fields: fields{
				Cmds: []Cmd{
					{"nop", 1},
					{"acc", 1},
					{"jmp", 4},
					{"acc", 3},
					{"jmp", -3},
					{"acc", -99},
					{"jmp", -4},
					{"jmp", -5},
					{"acc", 6},
				},
			},
			wantEnds: 0,
			wantAcc:  0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Cmds: tt.fields.Cmds,
			}
			gotEnds, gotAcc, err := p.AnsTwo()
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.AnsTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEnds != tt.wantEnds {
				t.Errorf("Puzzle.AnsTwo() gotEnds = %v, want %v", gotEnds, tt.wantEnds)
			}
			if gotAcc != tt.wantAcc {
				t.Errorf("Puzzle.AnsTwo() gotAcc = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}
