package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nmueller-realdigital/aoc2020-day01/aoc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	file, err := os.Open("data/input.txt")
	if err != nil {
		return err
	}
	data, err := aoc.ReadAndParseInput(file)
	if err != nil {
		return err
	}
	x, y := aoc.FindPair(data)
	fmt.Println(x, y, x*y)
	a, b, c := aoc.FindTriple(data)
	fmt.Println(a, b, c, a*b*c)
	return nil
}
