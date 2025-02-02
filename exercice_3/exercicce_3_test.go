package exercice3

import (
	"fmt"
	"testing"
)

func TestKnapsackSolution(t *testing.T) {
	items := []Item{
		{Value: 60, Weight: 10},
		{Value: 50, Weight: 4},
		{Value: 70, Weight: 2},
	}
	knapsackWeight := 5

	maxValue, selectedValue := KnapsackSolution(items, knapsackWeight)
	fmt.Printf("Items: %v\n", items)
	fmt.Printf("Poids du sac: %v\n", knapsackWeight)
	fmt.Printf("Valeur maximale: %d\n", maxValue)
	fmt.Printf("Selected items: %v\n\n", selectedValue)

	items = []Item{
		{Value: 60, Weight: 10},
		{Value: 100, Weight: 20},
		{Value: 120, Weight: 30},
	}
	knapsackWeight = 1

	maxValue, selectedValue = KnapsackSolution(items, knapsackWeight)
	fmt.Printf("Items: %v\n", items)
	fmt.Printf("Poids du sac: %v\n", knapsackWeight)
	fmt.Printf("Valeur maximale: %d\n", maxValue)
	fmt.Printf("Selected items: %v\n\n", selectedValue)

	items = []Item{
		{Value: 60, Weight: 10},
		{Value: 100, Weight: 20},
		{Value: 120, Weight: 30},
	}
	knapsackWeight = 100

	maxValue, selectedValue = KnapsackSolution(items, knapsackWeight)
	fmt.Printf("Items: %v\n", items)
	fmt.Printf("Poids du sac: %v\n", knapsackWeight)
	fmt.Printf("Valeur maximale: %d\n", maxValue)
	fmt.Printf("Selected items: %v\n\n", selectedValue)
}
