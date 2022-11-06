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

// if already exist don't add new node
func ifExist(vertex []Vertex, x int) bool {
	// if i != len(vertex) {
	// 	if x == vertex[i].key {
	// 		return true
	// 	} else {
	// 		i++
	// 		ifExist(vertex.[i], x, i)
	// 	}
	// }
	for _, v := range vertex {
		if x == v.key {
			return true
		}
	}
	return false

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
	if ifExist(graph.vertices, n) {
		fmt.Printf("Vertex %v is not added because it is an existing key", n)
		return graph
	} else {
		graph.vertices = append(graph.vertices, Vertex{key: n}) //append the new vertex to the vertices list in the graph
		return graph
	}
}

// add edge for derected graph
func (graph Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := graph.get
	//check error => edge always exist and vertex doesn't exist
	//add edge
}

// getVertex returns a pointer to the vertex with a key integer
func (graph Graph) getVertex(k int) Vertex {
	for i, v := 
}

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}
	test = test.addvertex(9)
	test.printGraph()

}
