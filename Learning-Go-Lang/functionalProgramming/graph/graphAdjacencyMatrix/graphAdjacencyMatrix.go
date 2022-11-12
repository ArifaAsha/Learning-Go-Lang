package main

import "fmt"

type Graph struct {
	vertices []Vertex
}

type Vertex struct {
	key      int
	adjacent []Vertex
}

func (graph Graph) printGraph() {
	for _, v := range graph.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent { //vertices inside the adjacency list
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}

func (graph Graph) addvertex(n int) Graph {
	graph.vertices = append(graph.vertices, Vertex{key: n}) //append the new vertex to the vertices list in the graph
	return graph
}

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}

	fmt.Println(test)
	fmt.Println("----------------")

}
