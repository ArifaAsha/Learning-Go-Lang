package main

import "fmt"

// https://www.youtube.com/watch?v=bSZ57h7GN2w

// Graph structure => adjacency list
type Graph struct {
	vertices []Vertex //list of vertices
}

// vertex structure => a graph vertex
type Vertex struct {
	key      int
	adjacent []Vertex
}

// print adjacent list of each vertex
func (graph Graph) printGraph() {
	for _, v := range graph.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent { //vertices inside the adjacency list
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}

// add vertex to the graph
func (graph Graph) addvertex(n int) Graph {
	graph.vertices = append(graph.vertices, Vertex{key: n}) //append the new vertex to the vertices list in the graph
	return graph
}

// find vertex to add adjacent
func (graph Graph) getVertex(k int) Vertex {
	for i, v := range graph.vertices {
		if v.key == k {
			return graph.vertices[i]
		}
	}
	return graph.vertices[0]
}

// Replace the new adjacent nodes list with the previous -> making immutable
func (graph Graph) replaceAdacencyList(k int, vertex Vertex) Graph {
	for i, v := range graph.vertices {
		if v.key == k {
			graph.vertices[i].adjacent = vertex.adjacent
		}
	}
	return graph
}

func (graph Graph) addEdge(from, to int) Graph {
	fromVertex := graph.getVertex(from)
	toVertex := graph.getVertex(to)
	// fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	fromVertex.adjacent = append(graph.getVertex(from).adjacent, toVertex)
	// graph.getVertex(from) = fromVertex.adjacent
	graph = graph.replaceAdacencyList(from, fromVertex)
	// fmt.Println(fromVertex)
	return graph
}

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}
	// test = test.addvertex(9)

	g := test.addEdge(1, 2)

	g2 := g.addEdge(2, 3)
	g2.printGraph()
	// g2.printGraph()

	fmt.Println()
	fmt.Println("----------------")

}
