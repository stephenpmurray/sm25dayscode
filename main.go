package main

import (
	"fmt"
	"os"

	"githhub.com/stephenmurrayengineer/sm25dayscode/day01"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day02"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day05"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day06"
)

func main() {

	day01.RunDay01()

	var day2 day02.PuzzInput
	err := day2.Import("input")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Day Two Part One: ", day2.AnsTwoOne())
	fmt.Println("Day Two Part One: ", day2.AnsTwoTwo())
	/*
		var p day03.Puzzle
		p.Input("input")
	*/
	var p5 day05.Puzzle
	passes, err := day05.Input("./day05/input")
	p5.ReadPassports(passes)
	max, _ := p5.AnsOne()
	fmt.Println("Day5: Maximum value is", max)
	fmt.Println("Day5: Your seat id is", p5.AnsTwo())

	//day 6 answers
	input, err := day05.Input("./day06/input")
	if err != nil {
		os.Exit(1)
	}
	runes := day06.ProcInputSets(input)
	fmt.Println("Day6 Part 1 is", day06.AnsOne(runes))
	msets, ppl := day06.ProcInputFull(input)
	ans, err := day06.AnsTwo(runes, msets, ppl)
	if err == nil {
		fmt.Println("Day6 Part 2 is", ans)
	}

}
