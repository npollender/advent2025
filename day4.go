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
	if (x == 0 && y == 0) ||
		(x == 0 && y == cols-1) ||
		(x == rows-1 && y == 0) ||
		(x == rows-1 && y == cols-1) {
		return true
	}
	return false
}

func hasAccess(x int, y int, rows int, cols int, grid [][]byte) bool {
	var papers int = 0
	if x > 0 && grid[x-1][y] == '@' { //down
		papers++
	}
	if x < rows-1 && grid[x+1][y] == '@' { //up
		papers++
	}
	if y > 0 && grid[x][y-1] == '@' { //left
		papers++
	}
	if y < cols-1 && grid[x][y+1] == '@' { //right
		papers++
	}
	if x > 0 && y > 0 && grid[x-1][y-1] == '@' { //diagonals
		papers++
	}
	if x > 0 && y < cols-1 && grid[x-1][y+1] == '@' {
		papers++
	}
	if x < rows-1 && y > 0 && grid[x+1][y-1] == '@' {
		papers++
	}
	if x < rows-1 && y < cols-1 && grid[x+1][y+1] == '@' {
		papers++
	}
	if papers < 4 {
		return true
	}
	return false
}
