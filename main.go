package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"githhub.com/stephenmurrayengineer/sm25dayscode/day01"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day02"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day03"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day04"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day05"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day06"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day08"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day09"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day10"
)

func main() {

	day01.RunDay01()

	// day 2
	var day2 day02.PuzzInput
	err := day2.Import("input")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Day 2 Part 1: ", day2.AnsTwoOne())
	fmt.Println("Day 2 Part 1: ", day2.AnsTwoTwo())

	// day 3 answers
	input, err := day05.Input("./day03/input")
	if err != nil {
		log.Fatal(err)
	}
	var p3 day03.Puzzle
	p3.ReadAllTreeLines(input)
	fmt.Println("Day3 Part 1: ", p3.Ans(1, 3), "trees")
	p2Ins := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	prod := p3.Ans(p2Ins[0][0], p2Ins[0][1])
	for _, pair := range p2Ins[1:] {
		prod *= p3.Ans(pair[0], pair[1])
	}
	fmt.Println("Day3 Part 2: ", prod, "trees")

	// day 4 answers
	input, err = day05.Input("./day04/input")
	if err != nil {
		log.Fatal(err)
	}
	var p4 day04.Puzzle
	p4.ReadPassports(input)
	// fmt.Println("", p4.Records[287])
	fmt.Println("Day4 Part 1: ", p4.AnsOne(), "valid passports")
	ans, err := p4.AnsTwo()
	if err != nil {
		panic(err)
	}
	fmt.Println("Day4 Part 2: ", ans, "valid passports")

	// day 5 answers
	var p5 day05.Puzzle
	passes, err := day05.Input("./day05/input")
	p5.ReadPassports(passes)
	max, _ := p5.AnsOne()
	fmt.Println("Day5: Maximum value is", max)
	fmt.Println("Day5: Your seat id is", p5.AnsTwo())

	//day 6 answers
	input, err = day05.Input("./day06/input")
	if err != nil {
		log.Fatal(err)
	}
	runes := day06.ProcInputSets(input)
	fmt.Println("Day6 Part 1 is", day06.AnsOne(runes))
	msets, ppl := day06.ProcInputFull(input)
	ans, err = day06.AnsTwo(runes, msets, ppl)
	if err == nil {
		fmt.Println("Day6 Part 2 is", ans)
	}

	// day 8 answers
	input, err = day05.Input("./day08/input")
	if err != nil {
		log.Fatal(err)
	}
	var p8 day08.Puzzle
	err = p8.ProcInput(input)
	if err != nil {
		log.Fatal(err)
	}
	ans, _, _ = p8.CheckRun()
	fmt.Println("Day 8 part 1 is ", ans)
	_, acc, err := p8.AnsTwo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 8 part 2 ends with accumulator", acc)

	// day 9 answers
	input, err = day05.Input("./day09/input")
	if err != nil {
		log.Fatal(err)
	}
	ints, err := day09.ProcInput(input)
	if err != nil {
		log.Fatal(err)
	}
	ans, err = day09.AnsOne(25, ints)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 9 Part 1:", ans)
	start, end := day09.FindRun(ans, ints)
	min, max := day09.GetMinMax(start, end, ints)
	fmt.Println("Day 9 Part 2:", min+max)

	// day 10 answers
	input, err = day05.Input("./day10/input")
	if err != nil {
		log.Fatal(err)
	}
	ints, err = day09.ProcInput(input)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })
	ans, err = day10.AnsOne(ints)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 10 Part 1:", ans)
	fmt.Println("Day 10 Part 2:", day10.AnsTwo(ints))

}
