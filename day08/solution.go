package day08

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Puzzle struct {
	Cmds []Cmd
}

type Cmd struct {
	inst string
	val  int
}

func (p *Puzzle) ProcInput(input []string) (err error) {
	var matches [][]string
	var i int

	r := regexp.MustCompile(`(\w{3}) ([\+|\-]\d+)`)

	for idx, line := range input {
		matches = r.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			return errors.New("No matches at input line " + strconv.Itoa(idx))
		}
		i, err = strconv.Atoi(matches[0][2])
		if err != nil {
			return err
		}
		p.Cmds = append(p.Cmds, Cmd{
			inst: matches[0][1],
			val:  i,
		})
	}
	return nil
}

func (p *Puzzle) CheckRun() (acc int, ran map[*Cmd]bool, line int) {
	// function to check a run through of command list, returning the accumulator (acc)
	// value when the first repeated command occurs, as well as the list of run
	// commands (ran). Also return the line at which the checker exits, which
	// could be the last instruction

	l := len(p.Cmds) - 1
	ran = make(map[*Cmd]bool)

	for {
		// check if the address of cmd is in the set of run cmds
		if ran[&p.Cmds[line]] {
			break
		}
		// add to list of run commands
		ran[&p.Cmds[line]] = true

		// if last row of boot, return:
		if line == l {
			if p.Cmds[line].inst == "acc" {
				acc += p.Cmds[line].val
			}
			return acc, ran, line
		}

		switch p.Cmds[line].inst {
		case "acc":
			acc += p.Cmds[line].val
			line += 1
		case "nop":
			line += 1
		case "jmp":
			line += p.Cmds[line].val
			// can this go OOB?
		default:
			fmt.Println("invalid function type at line ", line)
			os.Exit(1)
		}
	}
	return acc, ran, line
}

func (p *Puzzle) flipCmd(line int) {
	switch p.Cmds[line].inst {
	case "nop":
		p.Cmds[line].inst = "jmp"
	case "jmp":
		p.Cmds[line].inst = "nop"
	default:
		return
	}
	return
}

func (p *Puzzle) AnsTwo() (ends, acc int, err error) {
	// brute-force approach

	l := len(p.Cmds) - 1

	// get the full list of instructions:
	_, ran, _ := p.CheckRun()

	// flip instructions if they're in the list and see if it completes
	// see what happens
	for i, _ := range p.Cmds {
		if ran[&p.Cmds[i]] {
			p.flipCmd(i)
			acc, _, ends = p.CheckRun()
			if ends == l {
				return ends, acc, nil
			} else {
				p.flipCmd(i)
			}
		}
	}
	err = errors.New("Error in Puzzle.AnsTwo() Only loops")
	return 0, 0, err
}
