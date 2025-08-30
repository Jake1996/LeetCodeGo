package solutions

import (
	"fmt"
	"math"
	"slices"
)

// https://leetcode.com/problems/koko-eating-bananas/description/
func minEatingSpeed(piles []int, h int) int {
	max := slices.Max(piles)
	min := 1
	found := max
	for min <= max {
		mid := (max + min) / 2
		hours := totalHours(piles, mid)
		if hours <= h {
			found = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return found
}

func totalHours(piles []int, r int) int {
	h := 0

	for _, pile := range piles {
		h += int(math.Ceil(float64(pile) / float64(r)))
	}
	return h
}

func MinEatingSpeerd() {
	fmt.Print(minEatingSpeed([]int{312884470}, 968709470))
}
