package day02

import (
	"reflect"
	"testing"
)

func TestPuzzInput_Import(t *testing.T) {
	type fields struct {
		Passwds []Password
	}
	type args struct {
		FileName string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValues []Password
		wantErr    bool
	}{
		{
			name:       "return passwd struct with good input",
			args:       args{"testinput"},
			wantValues: []Password{{PlcyChar: rune('a'), PlcyNum: [2]int{1, 3}, Password: "asdfasdf"}},
			wantErr:    false,
		},
		{
			name:       "return error when file doesn't exist",
			args:       args{"badfilename"},
			wantValues: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PuzzInput{
				Passwds: tt.fields.Passwds,
			}
			err := i.Import(tt.args.FileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("PuzzInput.Import() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(i.Passwds, tt.wantValues) {
				t.Errorf("PuzzInput.Import() = %v, want %v", i.Passwds, tt.wantValues)
			}
		})
	}
}

func TestPuzzInput_AnsTwoOne(t *testing.T) {
	type fields struct {
		Passwds []Password
	}
	tests := []struct {
		name   string
		fields fields
		wantC  int
	}{
		{
			"return correct number of instances",
			fields{
				[]Password{
					{PlcyChar: rune('a'), PlcyNum: [2]int{1, 3}, Password: "asdfasdf"},
					{PlcyChar: rune('c'), PlcyNum: [2]int{1, 3}, Password: "asdfcf"},
					{PlcyChar: rune('a'), PlcyNum: [2]int{1, 13}, Password: "aaasdf"},
					{PlcyChar: rune('a'), PlcyNum: [2]int{2, 8}, Password: "a"},
					{PlcyChar: rune('c'), PlcyNum: [2]int{2, 8}, Password: "aasdfasdf"},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PuzzInput{
				Passwds: tt.fields.Passwds,
			}
			if gotC := i.AnsTwoOne(); gotC != tt.wantC {
				t.Errorf("PuzzInput.AnsOne() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
