package solutions

import (
	"container/heap"
	"fmt"
	"slices"
)

type Course struct {
	CompleteBy int
}

// https://leetcode.com/problems/course-schedule-iii/

func scheduleCourse(courses [][]int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	time := 0
	slices.SortFunc(courses, func(a, b []int) int { return a[1] - b[1] })
	for _, course := range courses {
		if time+course[0] <= course[1] {
			heap.Push(&pq, &Item{
				Value: course[0],
				Additional: &Course{
					CompleteBy: course[1],
				},
			})
			time += course[0]
		} else if len(pq) > 0 {
			topVal := pq[0].Value
			if topVal > course[0] && time-topVal+course[0] <= course[1] {
				heap.Pop(&pq)
				heap.Push(&pq, &Item{
					Value: course[0],
					Additional: &Course{
						CompleteBy: course[1],
					},
				})
				time = time - topVal + course[0]
			}
		}
	}
	return len(pq)
}

func ScheduleCourse() {
	out := scheduleCourse([][]int{{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19}})
	fmt.Printf("%v", out)
}
