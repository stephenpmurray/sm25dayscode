package day02

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// read input from file
// get db entries
// get n (number entries)

type Password struct {
	PlcyChar rune
	PlcyNum  [2]int
	Password string
}

type PuzzInput struct {
	Passwds []Password
}

// TODO: this function is too long, split up
func (i *PuzzInput) Import(FileName string) (err error) {
	var filePath string

	pwd, _ := os.Getwd()
	// check both this folder and day02 folder for
	if _, err := os.Stat(pwd + "/" + FileName); os.IsNotExist(err) {
		filePath = pwd + "/day02/" + FileName
	} else {
		filePath = pwd + "/" + FileName
	}
	if err != nil {
		return err
	}

	// open fstream
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// read in and process lines of file
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		// TODO: put this into a function of it's own
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		splitLine := strings.FieldsFunc(line, f)

		numlow, _ := strconv.Atoi(splitLine[0])
		numhigh, _ := strconv.Atoi(splitLine[1])

		i.Passwds = append(i.Passwds, Password{
			PlcyChar: []rune(splitLine[2])[0],
			PlcyNum:  [2]int{numlow, numhigh},
			Password: splitLine[3],
		})
	}

	return nil
}

func (i *PuzzInput) AnsTwoOne() (c int) {
	c = 0

	for _, pass := range i.Passwds {
		nSubs := strings.Count(pass.Password, string(pass.PlcyChar))
		if nSubs >= pass.PlcyNum[0] && nSubs <= pass.PlcyNum[1] {
			c += 1
		}
	}

	return c
}

// TODO: I skipped tests for this function:
func (i *PuzzInput) AnsTwoTwo() (c int) {
	c = 0
	var first bool
	var secnd bool

	for _, pass := range i.Passwds {
		r := regexp.MustCompile(string(pass.PlcyChar))
		matches := r.FindAllIndex([]byte(pass.Password), -1)

		first = false
		secnd = false
		for _, n := range matches {
			if (n[0] + 1) == pass.PlcyNum[0] {
				first = true
			}
			if (n[0] + 1) == pass.PlcyNum[1] {
				secnd = true
			}
		}
		// fmt.Println(first, secnd)
		if (first || secnd) && !(first && secnd) {
			c += 1
		}
	}

	return c
}
