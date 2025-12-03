/* npollender | advent2025 day 2 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* part1 */
	ranges := strings.Split(strings.TrimSpace(input), ",")
	var id int64 = 0

	for _, interval := range ranges {
		bounds := strings.Split(interval, "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}

		for i := lower; i <= upper; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 == 0 { //odd number of digits will never have a repeating pair
				if findPattern(str, len(str)/2) {
					id += int64(i)
				}
			}
		}
	}
	fmt.Println(id)

	/* part 2 */
	id = 0
	for _, interval := range ranges {
		hits := make(map[int]bool)
		bounds := strings.Split(interval, "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}

		for i := lower; i <= upper; i++ {
			str := strconv.Itoa(i)
			if len(str) < 2 {
				continue
			}
			divs := getDivisors(len(str))
			for j := 0; j < len(divs); j++ {
				if findPattern(str, divs[j]) {
					if !hits[i] {
						id += int64(i)
						hits[i] = true //avoids same number with different pattern - for instance 2222,2222 and 22,22,22,22
					}
				}
			}
		}
	}
	fmt.Println(id)
}

func findPattern(str string, size int) bool {

	for i := size; i < len(str); i += size {
		if str[0:(size)] != str[i:(size+i)] {
			return false
		}
	}
	return true
}

func getDivisors(x int) []int {
	var divs []int
	for i := 1; i < x; i++ {
		if x%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}