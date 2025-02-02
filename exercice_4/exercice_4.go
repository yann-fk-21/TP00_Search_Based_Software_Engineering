package exercice_4

import "sort"

type Interval struct {
	StartTime int
	EndTime   int
}

func sortedIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].StartTime < intervals[j].StartTime
	})
	return intervals
}

func removeRedundancy(intervals []Interval) []Interval {
	for i := 0; i < len(intervals); i++ {
		if i+1 < len(intervals) {
			isRedundancy := (intervals[i].StartTime == intervals[i+1].StartTime) &&
				(intervals[i].EndTime == intervals[i+1].EndTime)
			if isRedundancy {
				intervals = append(intervals[:i], intervals[i+1:]...)
			}
		}

	}
	return intervals
}

func MergedOverlappingIntervals(intervalList map[string][]Interval) []Interval {
	var mergedIntervalList []Interval
	for _, intervals := range intervalList {
		mergedIntervalList = append(mergedIntervalList, sortedIntervals(intervals)...)
	}
	mergedIntervalList = sortedIntervals(mergedIntervalList)
	mergedIntervalList = removeRedundancy(mergedIntervalList)

	return mergedIntervalList
}
