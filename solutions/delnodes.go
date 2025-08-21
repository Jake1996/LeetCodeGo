package solutions

// https://leetcode.com/problems/delete-nodes-and-return-forest/description/

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDelete := make(map[int]bool)
	for _, d := range to_delete {
		toDelete[d] = true
	}
	var retList []*TreeNode
	if !toDelete[root.Val] {
		retList = append(retList, root)
	}
	res := delNodeHelper(root, toDelete)
	if len(res) > 0 {
		retList = append(retList, res...)
	}
	return retList
}

func delNodeHelper(root *TreeNode, to_delete map[int]bool) []*TreeNode {
	if root == nil {
		return nil
	}
	var retList []*TreeNode
	left := root.Left
	right := root.Right

	if to_delete[root.Val] {
		root.Left = nil
		root.Right = nil
		if left != nil && !to_delete[left.Val] {
			retList = append(retList, left)
		}
		if right != nil && !to_delete[right.Val] {
			retList = append(retList, right)
		}
	}
	if left != nil {
		le := delNodeHelper(left, to_delete)
		if len(le) > 0 {
			retList = append(retList, le...)
		}
		if to_delete[left.Val] {
			root.Left = nil
		}
	}
	if right != nil {
		re := delNodeHelper(right, to_delete)
		if len(re) > 0 {
			retList = append(retList, re...)
		}
		if to_delete[right.Val] {
			root.Right = nil
		}
	}
	return retList
}

func DelNodes() {
	delNodes(nil, []int{})
}
