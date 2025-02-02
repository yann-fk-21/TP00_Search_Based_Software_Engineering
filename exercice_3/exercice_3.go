package exercice3

type Item struct {
	Value  int
	Weight int
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func KnapsackSolution(items []Item, knapsackWeight int) (int, []Item) {
	dp := make([][]int, len(items)+1)

	for i := range dp {
		dp[i] = make([]int, knapsackWeight+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 1; j <= knapsackWeight; j++ {
			if items[i-1].Weight > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maximum(dp[i-1][j], items[i-1].Value+dp[i-1][j-items[i-1].Weight])
			}

		}
	}

	selectedItems := []Item{}
	i, j := len(items), knapsackWeight
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			selectedItems = append([]Item{items[i-1]}, selectedItems...) // Ajouter au début
			j -= items[i-1].Weight
		}
		i--
	}

	return dp[len(items)][knapsackWeight], selectedItems

}
