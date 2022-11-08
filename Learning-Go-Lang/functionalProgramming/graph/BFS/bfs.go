package main

import "fmt"

// https://www.youtube.com/watch?v=bSZ57h7GN2w

// Graph structure => adjacency list
type Graph struct {
	noOfVertices int
	vertices     []Vertex //list of vertices

}

// vertex structure => a graph vertex
type Vertex struct {
	key      int
	adjacent []Vertex
}

// if already exist don't add new node
// func ifExist(vertex []Vertex, x int) bool {
// 	// if i != len(vertex) {
// 	// 	if x == vertex[i].key {
// 	// 		return true
// 	// 	} else {
// 	// 		i++
// 	// 		ifExist(vertex.[i], x, i)
// 	// 	}
// 	// }
// 	for _, v := range vertex {
// 		if x == v.key {
// 			return true
// 		}
// 	}
// 	return false

// }

// func ifExist(vertex []Vertex, x int, index int) bool {
// 	if len(vertex) == 0 {
// 		return false
// 	} else if vertex[index].key == x {
// 		return true
// 	} else {
// 		print(vertex[index].key)
// 		index++
// 		ifExist(vertex[1:], x, index)
// 	}
// 	return false
// }

// func printV(vertex []Vertex, toAddAsVertex int, index int) bool {
// 	// firstIndexOfSlice := 0
// 	// fmt.Println(vertex)
// 	if len(vertex) != 0 {
// 		if vertex[0].key != toAddAsVertex {

// 			// fmt.Println("key ", vertex[index].key)
// 			// fmt.Println("to add ", toAddAsVertex)
// 			fmt.Println("index = ", index)
// 			index += 1
// 			printV(vertex[index:], toAddAsVertex, index)

// 		} else if vertex[index].key == toAddAsVertex {
// 			fmt.Println("hehe")
// 			return true
// 		}
// 	} else {
// 		return false
// 	}

// 	return false
// }

// func printV(vertex []Vertex, toAddAsVertex int, index int) {
// 	// fmt.Println("index = ", index)
// 	// if len(vertex) == 0 {
// 	// 	return
// 	// }
// 	// fmt.Println(len(vertex))
// 	if index < len(vertex) {
// 		fmt.Println("slice = ", vertex)
// 		fmt.Println("Index = ", index)
// 		index += 1

// 		printV(vertex[index:], toAddAsVertex, index)
// 	}
// }

// func printV(vertex []Vertex, index int, length int) {
// 	if index < 5 {
// 		fmt.Println(vertex)
// 		index++
// 		// if index == length-1 {
// 		// 	fmt.Println("hehe")
// 		// 	fmt.Println(vertex[1:index])
// 		// } else {
// 		printV(vertex[index:5], index, length)
// 		// }
// 	}
// }

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
	// if ifExist(graph.vertices, n) {
	// if ifExist(graph.vertices, n, 0, false) {
	// fmt.Printf("Vertex %v is not added because it is an existing key", n)
	// return graph
	// } else {
	graph.vertices = append(graph.vertices, Vertex{key: n}) //append the new vertex to the vertices list in the graph
	return graph
	// }
}

// add edge for derected graph
// func (graph *Graph) AddEdge(from, to int) {
// 	//get vertex
// 	fromVertex := graph.getVertex(from)
// 	toVertex := graph.getVertex(to)
// 	//check error => edge always exist and vertex doesn't exist
// 	//add edge
// 	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
// }

// // getVertex returns a pointer to the vertex with a key integer
// func (graph *Graph) getVertex(k int) *Vertex {
// 	for i, v := range graph.vertices {
// 		if v.key == k {
// 			return graph.vertices[i]
// 		}
// 	}
// 	return nil
// }
// ----------------
// func (graph Graph) getVertex(k int) Vertex {
// 	for i, v := range graph.vertices {
// 		if v.key == k {
// 			return graph.vertices[i]
// 		}
// 	}
// 	return graph.vertices[0]
// }
// ------------------------

func (graph Graph) getVertex(k int) Vertex {
	for i, v := range graph.vertices {
		if v.key == k {
			return graph.vertices[i]
		}
	}
	return graph.vertices[0]
}

func (graph Graph) addEdge(fromVertex int, toVertex int) Graph {
	fromVertexV := graph.getVertex(fromVertex)
	graph.vertices[fromVertex].adjacent = append(fromVertexV.adjacent[fromVertex], toVertex)
	graph.printGraph()
	return graph
}

// func (this *Graph) connect(start, last int) {
// 	var edge *AjlistNode = getAjlistNode(last)
// 	if this.node[start].next == nil {
// 		this.node[start].next = edge
// 	} else {
// 		// Add edge at the end
// 		this.node[start].last.next = edge
// 	}
// 	// Get last edge
// 	this.node[start].last = edge
// }

//  Handling the request of adding new edge
// func(this *Graph) addEdge(start, last int) {
//     if start >= 0 && start < this.size &&
//         last >= 0 && last < this.size {
//         // Safe connection
//         this.connect(start, last)
//     } else {
//         // When invalid nodes
//         fmt.Println("\nHere Something Wrong")
//     }
// }

// func (graph Graph) addEdge(from, to int) Graph {
// 	// get vertex
// 	fromVertex := graph.getVertex(from)
// 	toVertex := graph.getVertex(to)

// 	// fmt.Println(fromVertex, toVertex)
// 	//check error
// 	// add edge
// 	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

// 	// fmt.Println(graph.vertices[])

// 	// fmt.Println(fromVertex.adjacent)
// 	return graph
// }

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test = test.addvertex(i)
	}
	// test = test.addvertex(9)
	test.printGraph()
	test.addEdge(1, 2)
	fmt.Println("----------------")
	// g.printGraph()
	// ifExist(test.vertices, 1, 0)
	// printV(test.vertices, 5, 0)
	// fmt.Println("-------", len(test.vertices))

}
