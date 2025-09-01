package solutions

import (
	"fmt"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {
	hmap := map[int]int{}
	for _, num := range hand {
		hmap[num] = hmap[num] + 1
	}
	slices.Sort(hand)
	for _, num := range hand {
		if hmap[num] <= 0 {
			continue
		}
		count := hmap[num]
		for i := 1; i < groupSize; i++ {
			fmt.Print(i)
			if hmap[num+i] < count {
				return false
			}
			hmap[num+i] = hmap[num+i] - count
		}
		hmap[num] = 0
	}
	return true
}

func HandOfStraight() {
	fmt.Print(isNStraightHand([]int{}, 0))
}
