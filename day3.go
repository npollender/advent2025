/* npollender | advent2025 day 3 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* part1 */
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var joltage int = 0

	for _, line := range lines {
		var battery1 int = 0
		var battery2 int = 0
		var idx int = 0
		for i := 0; i < len(line)-1; i++ { //do not consider the last battery for first of the pair
			current := int(line[i] - '0')
			if current > battery1 {
				battery1 = current
				idx = i
			}
		}
		for i := idx + 1; i < len(line); i++ {
			current := int(line[i] - '0')
			if current > battery2 {
				battery2 = current
			}
		}
		joltageStr := strconv.Itoa(battery1) + strconv.Itoa(battery2)
		joltageNum, _ := strconv.Atoi(joltageStr) //til - you can use an _ instead of an err var
		joltage += joltageNum
	}
	fmt.Println(joltage)

	/* part 2 */
	joltage = 0

	for _, line := range lines {
		var batterySize int = 12
		var batteryBank string
		var idx int = 0
		for batterySize > 0 {
			var battery int = 0
			for i := idx; i <= len(line)-batterySize; i++ { //make sure we always have a battery bank that is 12 digits long
				current := int(line[i] - '0')
				if current > battery {
					battery = current
					idx = i + 1
				}
			}
			batteryBank += strconv.Itoa(battery) //different approach for part 2, concat highest batteries each iteration
			batterySize--
		}
		joltageNum, _ := strconv.Atoi(batteryBank)
		joltage += joltageNum
	}
	fmt.Println(joltage)
}
