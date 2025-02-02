package exercice2

import "testing"

func TestBFS(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}
	Bfs(graph, "A")
}

func TestDFS(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	visited := make(map[string]bool)
	
	DfsRecursive(graph, "A", visited)
}