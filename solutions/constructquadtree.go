package solutions

// https://leetcode.com/problems/construct-quad-tree/
type NodeQuad struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *NodeQuad
	TopRight    *NodeQuad
	BottomLeft  *NodeQuad
	BottomRight *NodeQuad
}

func construct(grid [][]int) *NodeQuad {
	return helperQuad(grid, 0, 0, len(grid))
}

func helperQuad(grid [][]int, x, y, s int) *NodeQuad {
	n := &NodeQuad{}
	if s == 1 {
		n.IsLeaf = true
		if grid[x][y] == 1 {
			n.Val = true
		}
		return n
	}
	topLeft := helperQuad(grid, x, y, s/2)
	topRight := helperQuad(grid, x, y+s/2, s/2)
	bottomLeft := helperQuad(grid, x+s/2, y, s/2)
	bottomRight := helperQuad(grid, x+s/2, y+s/2, s/2)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val &&
		bottomLeft.Val == bottomRight.Val &&
		topLeft.Val == bottomRight.Val {
		n.IsLeaf = true
		n.Val = topLeft.Val
	} else {
		n.TopLeft = topLeft
		n.TopRight = topRight
		n.BottomLeft = bottomLeft
		n.BottomRight = bottomRight
	}
	return n
}

func ConstructQuadTree() {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	_ = construct(grid)
}
