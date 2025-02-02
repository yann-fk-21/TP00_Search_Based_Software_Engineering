package exercice_5

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func KadaneAlgorithm (sequenceList []int) int {
   currentMax, globalMax := sequenceList[0], sequenceList[0]
   for i := 1; i < len(sequenceList); i++ {
		currentMax = maximum(sequenceList[i], sequenceList[i] + currentMax) 
	    globalMax = maximum(currentMax, globalMax)
   } 
   return globalMax
}
