package day04

import (
	"strings"
)

type Puzzle struct {
	Records []Record
}

type Record map[string]bool

func (p *Puzzle) ReadPassports(inputs []string) {
	r := Record{}
	var kvs []string

	f := func(c rune) bool {
		return c == rune(':') || c == rune(' ')
	}

	for _, line := range inputs {

		if line == "" {
			p.Records = append(p.Records, r)
			r = Record{}
			continue
		}

		kvs = strings.FieldsFunc(line, f)
		// fmt.Printf("%v", kvs)
		for idx, key := range kvs {
			if idx%2 == 0 {
				r[key] = true
			}
		}

	}
	// get last record:
	p.Records = append(p.Records, r)
}

func (p *Puzzle) AnsOne() (c int) {
	var flds int

	for _, r := range p.Records {
		flds = len(r)
		if (!r["cid"] && flds == 7) || flds == 8 {
			c += 1
		}

	}
	return c
}
