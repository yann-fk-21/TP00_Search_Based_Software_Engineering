package exercice_1

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println("*** Test results ****")
	sortedList := []int {1, 2, 3, 4, 5}
	min := 0
	max := len(sortedList) - 1
	target := 4
	result := BinarySearch(sortedList, target, min, max)
	fmt.Printf("Array: %v, Target: %v, Response: %d\n", sortedList, target, result)

	sortedList = []int {2, 4, 6, 8, 10, 12, 14, 16}
	min = 0
	max = len(sortedList) - 1
	target = 12
	result = BinarySearch(sortedList, target, min, max)
	fmt.Printf("Array: %v, Target: %v, Response: %d\n", sortedList, target, result)

	sortedList = []int {5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	min = 0
	max = len(sortedList) - 1
	target = 100
	result = BinarySearch(sortedList, target, min, max)
	fmt.Printf("Array: %v, Target: %v, Response: %d\n", sortedList, target, result)

	sortedList = []int {100, 200, 300, 400, 500, 600, 700}
	min = 0
	max = len(sortedList) - 1
	target = 800
	result = BinarySearch(sortedList, target, min, max)
	fmt.Printf("Array: %v, Target: %v, Response: %d\n", sortedList, target, result)
}