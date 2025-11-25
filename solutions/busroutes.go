package solutions

import (
	"fmt"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stopToRoutes := make(map[int][]int)
	for i, route := range routes {
		for _, stop := range route {
			stopToRoutes[stop] = append(stopToRoutes[stop], i)
		}
	}

	visitedRoutes := make(map[int]bool)
	visitedStops := make(map[int]bool)
	queue := []int{}
	for _, route := range stopToRoutes[source] {
		queue = append(queue, route)
		visitedRoutes[route] = true
	}
	buses := 1
	visitedStops[source] = true
	for len(queue) > 0 {
		next := []int{}
		for _, route := range queue {
			for _, stop := range routes[route] {
				if stop == target {
					return buses
				}
				if visitedStops[stop] {
					continue
				}
				visitedStops[stop] = true
				for _, newRoute := range stopToRoutes[stop] {
					if !visitedRoutes[newRoute] {
						visitedRoutes[newRoute] = true
						next = append(next, newRoute)
					}
				}
			}
		}
		queue = next
		buses++
	}
	return -1
}

func NumBusesToDestination() {
	fmt.Print(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
}
