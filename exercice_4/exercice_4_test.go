package exercice_4

import (
	"fmt"
	"testing"
)

func TestMergedOverlappingIntervals(t *testing.T) {
	intervalList := map[string][]Interval{
		"a": {{1, 3}},
		"b": {{5, 10}, {16, 20}},
	}
	response := MergedOverlappingIntervals(intervalList)
	fmt.Printf("intervals: %v\n", intervalList)
	fmt.Printf("Merged interval: %v\n\n", response)

	intervalList = map[string][]Interval{
		"a": {{1, 5}},
		"b": {{2, 4}, {3, 4}},
	}
	response = MergedOverlappingIntervals(intervalList)
	fmt.Printf("intervals: %v\n", intervalList)
	fmt.Printf("Merged interval: %v\n\n", response)

	intervalList = map[string][]Interval{
		"a": {{1, 4}, {2, 5}},
		"b": {{3, 6}, {5, 7}},
	}
	response = MergedOverlappingIntervals(intervalList)
	fmt.Printf("intervals: %v\n", intervalList)
	fmt.Printf("Merged interval: %v\n\n", response)

	intervalList = map[string][]Interval{}
	response = MergedOverlappingIntervals(intervalList)
	fmt.Printf("intervals: %v\n", intervalList)
	fmt.Printf("Merged interval: %v\n\n", response)

}
