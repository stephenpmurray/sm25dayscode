package day03

import (
	"bufio"
	"os"
)

const tree = rune('#')
const space = rune('.')

type Puzzle struct {
	Map [][]rune
}

// TODO: Use Interfaces to declare input() and Processline functions

func (p *Puzzle) readTreeLine(line string) {
	p.Map = append(p.Map, []rune(line))
}

func (p *Puzzle) Input(FileName string) (err error) {

	// get correct filepath
	var filePath string
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if _, err := os.Stat(pwd + "/" + FileName); os.IsNotExist(err) {
		filePath = pwd + "/day03/" + FileName
	} else {
		filePath = pwd + "/" + FileName
	}
	if err != nil {
		return err
	}

	//open fstream
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// read in lines of file
	s := bufio.NewScanner(file)
	if err != nil {
		return err
	}
	for s.Scan() {
		line := s.Text()
		p.readTreeLine(line)
	}

	return nil
}

(*Puzzle) func GetAnswer() {
	r = 0
	c = 0
	nTrees = 0
	oob = false // out of bounds

	for {
		
	}
}