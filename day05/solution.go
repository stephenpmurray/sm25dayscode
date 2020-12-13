package day05

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct {
	Passes []Pass
}

type Pass struct {
	Row int
	Col int
}

// Get raw input from file
func Input(filePath string) (passes []string, err error) {

	//open fstream
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read in lines of file
	s := bufio.NewScanner(file)
	if err != nil {
		return passes, err
	}

	for s.Scan() {
		line := s.Text()
		passes = append(passes, line)
	}

	return passes, nil
}

func (p *Puzzle) ReadPassports(s []string) (err error) {

	for _, line := range s {
		// convert chars to binary
		if line == "" {
			return errors.New("s []string contains empty strings")
		}
		r, err := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line, "F", "0"), "B", "1")[:7], 2, 8)
		c, err := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line, "L", "0"), "R", "1")[7:], 2, 8)
		if err != nil {
			return err
		}
		p.Passes = append(p.Passes, Pass{Row: int(r), Col: int(c)})
	}

	return nil
}

func (p *Puzzle) AnsOne() (max int, err error) {
	i := 0
	if len(p.Passes) == 0 {
		err := errors.New("Puzzle.Passes[] is empty")
		return 0, err
	}

	// get product of row and column
	for _, pass := range p.Passes {
		i = pass.Row*8 + pass.Col
		if i > max {
			max = i
		}
	}
	return max, nil
}

func (p *Puzzle) AnsTwo() int {
	var ids []int

	for _, pass := range p.Passes {
		ids = append(ids, pass.Row*8+pass.Col)
	}
	sort.Ints(ids)

	for idx, curr := range ids[:(len(ids) - 1)] {
		if (ids[idx+1] - curr) > 1 {
			return curr + 1
		}
	}

	// sort.Slice(p.Passes, func(i, j int) bool {
	// 	return p.Passes[i].Row < p.Passes[j].Row
	// })
	return 0
}
