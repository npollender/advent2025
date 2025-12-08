/* npollender | advent2025 day 6 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* part1 */
	lines := strings.Split(input, "\n")
	var grandTotal int = 0

	worksheet := make([][]string, len(lines))
	for i, line := range lines {
		worksheet[i] = strings.Fields(line)
	}

	rows := len(worksheet)
	cols := len(worksheet[0])

	for i := range cols {
		var total int = 0
		for j := 0; j < rows-1; j++ {
			num, _ := strconv.Atoi(worksheet[j][i])
			switch worksheet[rows-1][i] {
			case "+":
				total += num
			case "*":
				if total == 0 {
					total = 1
				}
				total *= num
			}
		}
		grandTotal += total
	}
	fmt.Println(grandTotal)

	/* part 2 */
	//we can't use the worksheet from part 1 since the data needs to be read differently
	grandTotal = 0
	var ops string = lines[len(lines)-1]

	for i := 0; i < len(ops); {
		var total int = 0
		var inc int = 0 //the space between the current operator and the next operator
		var nums []int
		for j := i + 1; j <= len(ops); j++ {
			if j == len(ops) || ops[j] == ' ' {
				inc++
			} else {
				break
			}
		}
		for j := 0; j < inc; j++ {
			var numStr string = ""
			for k := range lines {
				if k < len(lines)-1 && lines[k][i+j] != ' ' {
					numStr += string(lines[k][i+j])
				}
			}
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		for _, num := range nums {
			switch ops[i] {
			case '+':
				total += num
			case '*':
				if total == 0 {
					total = 1
				}
				total *= num
			}
		}
		grandTotal += total
		i += inc + 1 //next operator position
	}
	fmt.Println(grandTotal)

}
