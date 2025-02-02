package exercice_1

func BinarySearch(sortedList []int, target int, min, max int) int {
	middle := min + (max-min)/2

	if min <= max {

		if sortedList[middle] == target {
			return middle
		}

		if sortedList[middle] > target {
			max = middle - 1
			min = 0
			return BinarySearch(sortedList, target, min, max)
		}

		if sortedList[middle] < target {
			max = len(sortedList) - 1
			min = middle + 1
			return BinarySearch(sortedList, target, min, max)
		}
	}

	return -1

}
