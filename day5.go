/* npollender | advent2025 day 5 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ingredientRange struct {
	min, max int
}

func main() {
	/* part1 */
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var blank bool = false
	var ingredientsRange []ingredientRange
	var ingredients []int
	var valid int = 0

	for _, line := range lines { //separate ingredient ranges and valid ingredients
		if !blank && line == "" {
			blank = true
		}

		if !blank {
			parts := strings.Split(line, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			ingredientsRange = append(ingredientsRange, ingredientRange{min, max})
		} else {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	for _, ingredient := range ingredients {
		for _, ingredientRange := range ingredientsRange {
			if ingredient >= ingredientRange.min && ingredient <= ingredientRange.max {
				valid++
				break
			}
		}
	}
	fmt.Println(valid)

	/* part 2 */
	valid = 0
	var min, max int = 0, 0
	var unprocessed []ingredientRange

	for {
		for _, ingredientRange := range ingredientsRange {
			var inRange bool = false
			if min == 0 && max == 0 {
				min = ingredientRange.min
				max = ingredientRange.max
				continue
			}
			min, max, inRange = getNewRange(min, max, ingredientRange)
			if !inRange {
				unprocessed = append(unprocessed, ingredientRange)
			}
		}
		if len(unprocessed) == 0 { //all ranges processed
			break
		} else if len(unprocessed) == len(ingredientsRange) { //no progress made on previous iteration, sum up current range and reprocess
			valid += max - min + 1
			max, min = 0, 0
		}
		ingredientsRange = unprocessed //reprocess all the unprocessed ranges
		unprocessed = nil
	}
	valid += max - min + 1
	fmt.Println(valid)
}

// there's probably a better way to do this but this method is the most visual for me
func getNewRange(min int, max int, ingredientRange ingredientRange) (int, int, bool) {
	if ingredientRange.min >= min && ingredientRange.max <= max { //current range is exclusively within current min/max, no change required but inRange is true
		return min, max, true
	}
	if (ingredientRange.min < min && ingredientRange.max < min) || (ingredientRange.min > max && ingredientRange.max > max) { //current range is completely before or after current min/max, no change required
		return min, max, false
	}
	if ingredientRange.min <= min && ingredientRange.max >= max { //current range completely encompasses current min/max
		return ingredientRange.min, ingredientRange.max, true
	} else if ingredientRange.min <= min && ingredientRange.max >= min { //current range overlaps left side of current min/max
		return ingredientRange.min, max, true
	} else if ingredientRange.min <= max && ingredientRange.max >= max { //current range overlaps right side of current min/max
		return min, ingredientRange.max, true
	} else {
		return min, max, false //should not be reachable
	}
}
