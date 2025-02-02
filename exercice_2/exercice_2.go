package exercice2

import (
	"fmt"
)

func Bfs(graph map[string][]string, start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			fmt.Print(node, " ")
			visited[node] = true

			
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	fmt.Println(visited)
}

func DfsRecursive(graph map[string][]string, node string, visited map[string]bool) {
	if visited[node] {
		return
	}

	fmt.Print(node, " ")
	visited[node] = true

	for _, neighbor := range graph[node] {
		DfsRecursive(graph, neighbor, visited)
	}

	fmt.Println(visited)
}