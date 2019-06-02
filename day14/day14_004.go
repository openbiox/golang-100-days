package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	// gopl.io/ch4/graph
	addEdge("a", "b")
	addEdge("c", "a")
	for k, v := range graph {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(hasEdge("c", "a")) // true
	fmt.Println(hasEdge("a", "b")) // true
	fmt.Println(hasEdge("c", "c")) // false
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
