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

func (p *Puzzle) AnsOne() int {
	acc := 0
	line := 0
	ran := make(map[*Cmd]bool)

	for {
		// check if the address of cmd is in the set of run cmds
		if ran[&p.Cmds[line]] {
			break
		}
		// add to list of run commands
		ran[&p.Cmds[line]] = true

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
	return acc
}
