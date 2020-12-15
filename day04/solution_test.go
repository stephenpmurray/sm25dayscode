package day04

import (
	"reflect"
	"testing"
)

func TestPuzzle_ReadPassports(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		inputs []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantRecord []Record
	}{
		{
			name: "Read input correctly to p.Record map",
			args: args{
				[]string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
					"byr:1937 iyr:2017 cid:147 hgt:183cm",
					"",
					"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
					"hcl:#cfa07d byr:1929",
				},
			},
			wantRecord: []Record{
				{"ecl": true, "pid": true, "eyr": true, "hcl": true, "byr": true, "iyr": true, "cid": true, "hgt": true},
				{"iyr": true, "ecl": true, "cid": true, "eyr": true, "pid": true, "hcl": true, "byr": true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			p.ReadPassports(tt.args.inputs)
			if !reflect.DeepEqual(p.Records, tt.wantRecord) {
				t.Errorf("Test_Puzzle_ReadPassports(): want \n %v \n got \n %v", tt.wantRecord, p.Records)
			}
		})
	}
}

func TestPuzzle_AnsOne(t *testing.T) {
	type fields struct {
		Records []Record
	}
	tests := []struct {
		name   string
		fields fields
		wantC  int
	}{
		{
			"Returns correct count for a given p.Records field",
			fields{
				[]Record{
					{"ecl": true, "pid": true, "eyr": true, "hcl": true, "byr": true, "iyr": true, "cid": true, "hgt": true},
					{"iyr": true, "ecl": true, "cid": true, "eyr": true, "pid": true, "hcl": true, "byr": true},
					{"ecl": true, "pid": true, "eyr": true, "hcl": true, "byr": true, "iyr": true, "hgt": true},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			if gotC := p.AnsOne(); gotC != tt.wantC {
				t.Errorf("Puzzle.AnsOne() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
