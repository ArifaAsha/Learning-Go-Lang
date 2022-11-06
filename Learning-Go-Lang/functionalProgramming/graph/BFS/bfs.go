package main

import "fmt"

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
		fmt.Println("Vertex ", v.key)
		for _, v := range v.adjacent { //vertices inside the adjacency list
			fmt.Printf("%v", v.key)
		}
	}
}

// func (graph Graph) printGraph() {
// 	for _, v := range graph.vertices {
// 		fmt.Println("Vertex ", v.key)
// 		// for _, v := range v.adjacent { //vertices inside the adjacency list
// 		// 	fmt.Printf("%v", v.key)
// 		// }
// 	}
// }

// func (graph Graph) print(i int) {
// 	if i <= len(graph.vertices) {
// 		traverseSlicesOfAdjcent(graph.vertices[i], i)
// 		fmt.Printf("%T", graph.vertices[i])
// 		i++
// 	}
// }

// func (graph Graph) traverseSlicesOfAdjcent(vertex []Vertex, i int) {
// 	if i <= len(vertex[i].adjacent) {
// 		fmt.Println(vertex.adjacent[i])
// 		i++
// 	}
// }

// add vertex to the graph
func (graph Graph) addvertex(n int) Graph {
	// v := Vertex{key: n}                        //vertex v with key k
	graph.vertices = append(graph.vertices, Vertex{key: n}) //append the new vertex to the vertices list in the graph
	return graph
}

func main() {
	test := Graph{}
	// g := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}
	test.printGraph()

}
