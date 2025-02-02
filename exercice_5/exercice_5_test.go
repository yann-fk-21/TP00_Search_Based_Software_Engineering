package exercice_5

import (
	"fmt"
	"testing"
)

func TestKadaneAlgorithm(t *testing.T) {
	sequenceList := []int {23, 12, 1, 8, 9, 100}
	maxSum := KadaneAlgorithm(sequenceList)
	fmt.Printf("Array: %v, max_sum: %v\n", sequenceList, maxSum)

	sequenceList = []int {-2, -67, -10, -80, -200, -40}
	maxSum = KadaneAlgorithm(sequenceList)
	fmt.Printf("Array: %v, max_sum: %v\n", sequenceList, maxSum)

	sequenceList = []int {-1, 10, 25, 50, -67, -200}
	maxSum = KadaneAlgorithm(sequenceList)
	fmt.Printf("Array: %v, max_sum: %v\n", sequenceList, maxSum)
}