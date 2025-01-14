package day19

import (
	"fmt"
	"strings"
)

type Solution struct {
	Towels  []Towel
	Designs []Design
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 19:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Possible designs: %d\n", PossibleDesigns(s.Designs, s.Towels))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Total combinations: %d\n", TotalCombinations(s.Designs, s.Towels))

	return nil
}

// PossibleDesigns returns the set of all designs that have one or more solutions with these towels.
func PossibleDesigns(designs []Design, towels []Towel) int {
	possibleDesigns := 0
	for _, design := range designs {
		if isPossible(design, towels) {
			possibleDesigns++
		}
	}
	return possibleDesigns
}

// TotalCombinations returns the total number of possible towel combinations that meet the given designs.
func TotalCombinations(designs []Design, towels []Towel) int {
	totalCombinations := 0
	cache := map[Design]int{}
	for _, design := range designs {
		combinations := combinationsFor(design, towels, cache)
		if combinations == 0 {
			continue // impossible design
		}
		totalCombinations += combinations
	}
	return totalCombinations
}

func isPossible(design Design, towels []Towel) bool {
	for _, towel := range towels {
		if towel == design {
			return true
		}
		if strings.HasPrefix(design, towel) && isPossible(design[len(towel):], towels) {
			return true // ok: towel meets design
		}
	}
	return false
}

func combinationsFor(design Design, towels []Towel, cache map[Design]int) (result int) {
	if cached, ok := cache[design]; ok {
		return cached
	}

	candidateTowels := make([]Towel, 0, len(towels))
	for _, towel := range towels {
		if strings.Contains(design, towel) {
			candidateTowels = append(candidateTowels, towel)
		}
	}

	for _, towel := range candidateTowels {
		if !strings.HasPrefix(design, towel) {
			continue // not a candidate
		}

		if towel == design {
			// ok: this towel completes the design
			result++
			continue
		}

		// candidate: search for matching children
		result += combinationsFor(design[len(towel):], candidateTowels, cache)
	}

	cache[design] = result
	return result
}
