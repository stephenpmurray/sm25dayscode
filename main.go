package main

import (
	"fmt"
	"os"

	"githhub.com/stephenmurrayengineer/sm25dayscode/day01"
	"githhub.com/stephenmurrayengineer/sm25dayscode/day02"
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
}
