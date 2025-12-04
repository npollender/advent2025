/* npollender | advent2025 day 4 */
 
package main

import (
	"fmt"
	"strings"
)

func main() {
	/* part1 */
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var rolls int = 0

	rows := len(lines)    //vertical
	cols := len(lines[0]) //horizontal

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '@' {
				if isCorner(i, j, rows, cols) || hasAccess(i, j, rows, cols, grid) {
					rolls++
				}
			}
		}
	}
	fmt.Println(rolls)

	/* part 2 */
	rolls = 0
	for {
		var removed int = 0
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == '@' {
					if isCorner(i, j, rows, cols) || hasAccess(i, j, rows, cols, grid) {
						rolls++
						grid[i][j] = '.' //removed roll
						removed++
					}
				}
			}
		}
		if removed == 0 { //keep looping until no more rolls can be removed
			break
		}
	}
	fmt.Println(rolls)
}

func isCorner(x int, y int, rows int, cols int) bool {
	return (x == 0 && y == 0) ||
		   (x == 0 && y == cols-1) ||
		   (x == rows-1 && y == 0) ||
		   (x == rows-1 && y == cols-1)

}

func hasAccess(x int, y int, rows int, cols int, grid [][]byte) bool {
	var papers int = 0
	var dirs = [8][2]int{
		{-1, 0}, {1, 0}, //up + down
		{0, -1}, {0, 1}, // left + right
		{-1, -1}, {-1, 1}, //diagonals ...
		{1, -1}, {1, 1},
	}

	for _, offset := range dirs {
		xDir := x + offset[0]
		yDir := y + offset[1]

		if (xDir >= 0 && xDir < rows && yDir >= 0 && yDir < cols) &&
			 (grid[xDir][yDir] == '@') {
			papers++
		}
	}
	return papers < 4

}
