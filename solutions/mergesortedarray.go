package solutions

// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if i >= 0 && j >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else if i >= 0 && j >= 0 && nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		k--
		i--
	}
}

func MergeSortedArray() {
	merge([]int{0}, 0, []int{1}, 1)
}
