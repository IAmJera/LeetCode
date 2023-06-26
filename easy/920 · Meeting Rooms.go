package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func main() {
	fmt.Println(CanAttendMeetings([]*Interval{{0, 30}, {15, 20}, {5, 10}}))
	fmt.Println(CanAttendMeetings([]*Interval{{5, 8}, {9, 15}}))
}
func CanAttendMeetings(intervals []*Interval) bool {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End > intervals[i+1].Start {
			return false
		}
	}
	return true
}
