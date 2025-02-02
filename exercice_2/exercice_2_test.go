package exercice2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}
	fmt.Println("***Test BFS***")
	Bfs(graph, "A")
	fmt.Printf("\n\n")
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
	fmt.Println("***Test DFS***")
	DfsRecursive(graph, "A", visited)
	fmt.Printf("\n\n")
}


func TestBfsShortestPath(t *testing.T) {
    graph := map[string][]string{
        "A": {"B", "C"},
        "B": {"A", "D", "E"},
        "C": {"A", "F"},
        "D": {"B"},
        "E": {"B", "F"},
        "F": {"C", "E"},
    }

    tests := []struct {
        start     string
        goal      string
        expected  []string
        connected bool
    }{
        {"A", "F", []string{"A", "C", "F"}, true},
        {"A", "D", []string{"A", "B", "D"}, true},
        {"A", "E", []string{"A", "B", "E"}, true},
        {"D", "F", []string{"D", "B", "E", "F"}, true},
        {"A", "A", []string{"A"}, true},
        {"A", "G", nil, false}, // "G" n'existe pas dans le graphe
    }

	fmt.Println("**Test ShortestPath***")
    for _, test := range tests {
        path, connected := BfsShortestPath(graph, test.start, test.goal)
        fmt.Printf("Test de %s à %s: Chemin trouvé: %v, Connecté: %v\n", test.start, test.goal, path, connected)
        if !reflect.DeepEqual(path, test.expected) || connected != test.connected {
            t.Errorf("BfsShortestPath(%q, %q) = %v, %v; want %v, %v",
                test.start, test.goal, path, connected, test.expected, test.connected)
        }
    }
}