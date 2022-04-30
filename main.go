package main

import (
	"combinations/combos"
)

func main() {
	set := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	// just testing out how the "generator" like version preforms against the
	// non-generator.
	allFromGenerator(set)
	combos.All(set)
}

func allFromGenerator(set []string) (subsets [][]string) {
	for n := 0; n < len(set); n++ {

		abort := make(chan struct{})
		ch := combos.CombinationsGenerator(abort, set, n)
		for subset := range ch {
			subsets = append(subsets, subset)
		}
		close(abort)
	}

	return subsets
}
