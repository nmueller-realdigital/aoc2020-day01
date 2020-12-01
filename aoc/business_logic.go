// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

// FindPair finds two int in a int slice that sum up to 2020.
func FindPair(data []int) (int, int) {
	for idx, xs := range data {
		for idy, ys := range data {
			if idx != idy && xs+ys == 2020 {
				return xs, ys
			}
		}
	}
	return 0, 0
}

// FindTriple finds three int in a int slice that sum up to 2020.
func FindTriple(data []int) (int, int, int) {
	for idx, xs := range data {
		for idy, ys := range data {
			for idz, zs := range data {
				if idx != idy && idy != idz && xs+ys+zs == 2020 {
					return xs, ys, zs
				}
			}
		}
	}
	return 0, 0, 0
}
