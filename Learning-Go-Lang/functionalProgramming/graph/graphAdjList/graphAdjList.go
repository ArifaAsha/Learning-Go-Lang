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
// func (graph Graph) printGraph() {
// 	for _, v := range graph.vertices {
// 		fmt.Printf("\nVertex %v: ", v.key)
// 		for _, v := range v.adjacent { //vertices inside the adjacency list
// 			fmt.Printf("%v ", v.key)
// 		}
// 	}
// 	fmt.Println()
// }

// print graph
func (graph Graph) returnKey(i int) {
	if i == len(graph.vertices) {
		// fmt.Println("shesh")
		return
	} else {
		key := graph.vertices[i].key
		findVertex := graph.getVertex(key)
		fmt.Print("Vertex ", findVertex.key, ": ")
		iterateOverAdjacent(findVertex, 0)
		// iterateOverAdjacentSlice(findVertex, 0)
		fmt.Println()
		i++
		graph.returnKey(i)

	}
}

// func iterateOverAdjacent(vertex Vertex, i int) {
// 	if i == len(vertex.adjacent) {
// 		return
// 	} else {
// 		printKeyOfTheVertex(vertex.adjacent[i])
// 		i++
// 		iterateOverAdjacent(vertex.adjacent[i], i)
// 	}
// }

// ===========
func iterateOverAdjacent(vertex Vertex, i int) {
	if i == len(vertex.adjacent) {
		return
	} else {
		if len(vertex.adjacent) == 0 {
			printKeyOfTheVertex(vertex)
		} else {
			//iterate over the list
			getListOfVertex(vertex.adjacent, vertex, 0)

		}
	}
}

func getListOfVertex(vertexList []Vertex, vertex Vertex, i int) {
	// fmt.Println("when Entering: ", len(vertexList))
	if i == len(vertexList) {
		return
	} else {
		vertex = vertexList[i]
		printKeyOfTheVertex(vertex)
		fmt.Print(" ")
		// fmt.Println("i = ", i)
		if i+1 != len(vertexList) {
			i++
			getListOfVertex(vertexList, vertexList[i], i)
		}

	}
}

// func iterateOverAdjacent(vertex Vertex, i int) {
// 	if i == len(vertex.adjacent) {
// 		return
// 	} else {
// 		if len(vertex.adjacent) == 0 {
// 			printKeyOfTheVertex(vertex)
// 		} else {
// 			//iterate over the list

// 		}
// 	}
// }

// func iterateOverAdjacent(vertex Vertex, i int) {
// 	if i == len(vertex.adjacent) {
// 		return
// 	} else {
// 		if len(vertex.adjacent) == 0 {
// 			printKeyOfTheVertex(vertex)
// 		} else {
// 			vertex = vertex.adjacent[i]
// 			printKeyOfTheVertex(vertex.adjacent[i])
// 			i++
// 			iterateOverAdjacent(vertex, i)
// 			// vertex = vertex.adjacent[i]
// 			// printKeyOfTheVertex(vertex)
// 			// i++
// 			// iterateOverAdjacent(vertex, i)

// 			// iterateOverAdjacent(vertex, i)
// 		}
// 		// printKeyOfTheVertex(vertex.adjacent[i].key)
// 		// i++
// 		// iterateOverAdjacent(vertex.adjacent[i], i)
// 	}
// }

func printKeyOfTheVertex(vertex Vertex) {
	fmt.Print(vertex.key)
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
	fromVertex.adjacent = append(graph.getVertex(from).adjacent, toVertex)
	graph = graph.replaceAdacencyList(from, fromVertex)
	return graph
}

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}

	g := test.addEdge(1, 2) //add 2 in adj list of vertex 1
	g2 := g.addEdge(2, 3)
	g3 := g2.addEdge(2, 4)
	g4 := g3.addEdge(1, 3)

	g4.returnKey(0)

	fmt.Println()
	fmt.Println("----------------")

}
