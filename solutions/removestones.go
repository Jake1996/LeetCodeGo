package solutions

import "fmt"

// DSU struct represents the Disjoint Set Union data structure.
type DSU struct {
	parent []int
}

// NewDSU initializes a new DSU structure with n elements.
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSU{parent: parent}
}

// Find returns the representative of the set that the element i belongs to.
// It uses path compression for optimization.
func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

// Union merges the sets containing elements i and j.
func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootJ] = rootI
	}
}

// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
func removeStones(stones [][]int) int {
	n := len(stones)
	dsu := NewDSU(n)

	// Use maps to group stones by row and column
	rowMap := make(map[int]int)
	colMap := make(map[int]int)

	for i, stone := range stones {
		x, y := stone[0], stone[1]

		// Check if this row has been seen before
		if _, ok := rowMap[x]; !ok {
			rowMap[x] = i
		} else {
			// If it has, union the current stone with the first stone in that row
			dsu.Union(i, rowMap[x])
		}

		// Check if this column has been seen before
		if _, ok := colMap[y]; !ok {
			colMap[y] = i
		} else {
			// If it has, union the current stone with the first stone in that column
			dsu.Union(i, colMap[y])
		}
	}

	// Count the number of connected components
	numComponents := 0
	for i := 0; i < n; i++ {
		if dsu.parent[i] == i {
			numComponents++
		}
	}

	// The maximum number of removable stones is the total number of stones
	// minus the number of connected components.
	return n - numComponents
}

func RemoveStones() {
	fmt.Print(removeStones([][]int{}))
}
