package solutions

import "fmt"

// Not working
type TimeNode struct {
	start     int
	end       int
	leftNode  *TimeNode
	rightNode *TimeNode
}

type MyCalendar struct {
	root *TimeNode
}

func ConstructorCal() MyCalendar {
	return MyCalendar{}
}

// https://leetcode.com/problems/my-calendar-i/
func (mc *MyCalendar) Book(startTime int, endTime int) bool {
	if mc.root == nil {
		mc.root = &TimeNode{start: startTime, end: endTime}
		return true
	}
	return mc.insert(mc.root, &TimeNode{start: startTime, end: endTime})
}

func (mc *MyCalendar) insert(cur *TimeNode, newNode *TimeNode) bool {
	if newNode.start >= cur.end {
		if cur.rightNode == nil {
			cur.rightNode = newNode
			return true
		}
		return mc.insert(cur.rightNode, newNode)
	} else if newNode.end <= cur.start {
		if cur.leftNode == nil {
			cur.leftNode = newNode
			return true
		}
		return mc.insert(cur.leftNode, newNode)
	}
	return false
}

func MyCalendarBook() {
	mc := ConstructorCal()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(15, 25))
	fmt.Println(mc.Book(20, 30))
}
