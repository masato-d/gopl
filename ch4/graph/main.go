package main

import "fmt"

var graph = make(map[string]map[string]bool)

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

func main() {
	fmt.Println(graph)
	addEdge("from1", "to1")
	fmt.Println(graph)
	fmt.Println(hasEdge("from1", "to1"))
	addEdge("begin", "end")
	fmt.Println(graph)
	fmt.Println(hasEdge("begin", "end"))
	addEdge("from1", "to2")
	fmt.Println(graph)
	fmt.Println(hasEdge("from2", "to2"))
}