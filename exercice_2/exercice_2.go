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

func BfsShortestPath(graph map[string][]string, start, goal string) ([]string, bool) {
    if start == goal {
        return []string{start}, true
    }

    
    queue := [][]string{{start}}
    visited := map[string]bool{start: true}

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]

        node := path[len(path)-1]

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
               
                newPath := append([]string{}, path...)
                newPath = append(newPath, neighbor)

                
                if neighbor == goal {
                    return newPath, true
                }

                
                queue = append(queue, newPath)
                visited[neighbor] = true
            }
        }
    }

    return nil, false
}