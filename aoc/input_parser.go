// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

// ReadAndParseInput reads challenge data from an io.Reader and returns its parsed contents.
func ReadAndParseInput(r io.Reader) ([]int, error) {
	csvReader := csv.NewReader(r)
	var result []int
	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		recordAsInt, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		result = append(result, recordAsInt)
	}
}
