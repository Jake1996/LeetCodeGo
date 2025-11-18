package solutions

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

const MOD = 1000000007

type MaxProduct struct {
	maxProd  int
	totalSum int
}

func maxProduct(root *TreeNode) int {
	p := &MaxProduct{
		totalSum: totalSum(root),
	}
	p.CalculateMax(root)
	return p.maxProd % MOD
}

func (m *MaxProduct) CalculateMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := m.CalculateMax(root.Left)
	right := m.CalculateMax(root.Right)
	thisTree := left + right + root.Val
	m.maxProd = max(m.maxProd, (m.totalSum-thisTree)*thisTree)
	return thisTree
}

func totalSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + totalSum(root.Left) + totalSum(root.Right)
}

func MaxProductSplittedTree() {
	out := maxProduct(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}})
	fmt.Printf("%v", out)
}
