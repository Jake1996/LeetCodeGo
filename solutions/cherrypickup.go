package solutions

import (
	"fmt"
)

// https://leetcode.com/problems/cherry-pickup/description/

type cherryPick struct {
	grid       [][]int
	calculated [][][]int
	n          int
}

const MIN = -999999

func cherryPickup(grid [][]int) int {

	n := len(grid)
	calc := make([][][]int, n)

	for i := range n {
		calc[i] = make([][]int, n)
		for j := range n {
			calc[i][j] = make([]int, n)
			for k := range n {
				calc[i][j][k] = MIN
			}
		}
	}
	dat := &cherryPick{
		grid:       grid,
		n:          n,
		calculated: calc,
	}

	return max(0, helper(0, 0, 0, dat))
}

func helper(i1, j1, i2 int, data *cherryPick) int {
	j2 := i1 + j1 - i2

	// outside grid
	if i1 >= data.n || i2 >= data.n || j1 >= data.n || j2 >= data.n {
		return MIN
	}

	// No way forward
	if data.grid[i1][j1] == -1 || data.grid[i2][j2] == -1 {
		return MIN
	}

	if data.calculated[i1][j1][i2] != MIN {
		return data.calculated[i1][j1][i2]
	}

	if i1 == data.n-1 && j1 == data.n-1 && i2 == data.n-1 {
		data.calculated[i1][j1][i2] = data.grid[i1][j1]
		return data.calculated[i1][j1][i2]
	}

	ans := 0
	// Pick Cherries
	if i1 == i2 && j1 == j2 && data.grid[i1][j1] == 1 {
		ans = 1
	} else {
		ans += data.grid[i1][j1]
		ans += data.grid[i2][j2]
	}

	ans += max(
		helper(i1+1, j1, i2+1, data), // both go down
		helper(i1, j1+1, i2, data),   // both go right
		helper(i1+1, j1, i2, data),   // one goes down and one right
		helper(i1, j1+1, i2+1, data), // one goes right and one down
	)

	data.calculated[i1][j1][i2] = ans
	return ans
}

func CherryPickup() {
	out := cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}})
	fmt.Printf("%v", out)
}
