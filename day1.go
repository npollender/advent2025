/* npollender | advent2025 day 1 */
/* input var is a raw string literal in input.go */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* part1 */
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var dial int = 50
	var pw int = 0

	for _, line := range lines {
		rotation := line[0]
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		switch rotation {
		case 'R':
			dial += distance
		case 'L':
			dial -= distance
		}
		if dial%100 == 0 { //if mod 100 is 0, that means dial is at 0
			pw++
		}
	}
	fmt.Println(pw)

	/* part2 */
	var clicks int = 0
	dial = 50
	pw = 0
	for _, line := range lines {
		rotation := line[0]
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		switch rotation {
		case 'R':
			if dial % 100 == 0 {
				clicks = 100 //dial is set at 0, next click is a full turn
			} else if dial < 0 {
				clicks = abs(dial % 100)
			} else {
				clicks = 100 - (dial % 100)
			}
			dial += distance
		case 'L':
			if dial % 100 == 0 {
				clicks = 100
			} else if dial < 0 {
				clicks = 100 - abs(dial % 100)
			} else {
				clicks = dial % 100
			}
			dial -= distance
		}
		if distance >= clicks {
			pw += ((distance - clicks) / 100) + 1
		}
	}
	fmt.Println(pw)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}