package main

import "fmt"

type Graph struct {
	vertices int
	adgeList [][]int
}

// func getGraph(matrix [][]int) *Graph {
// 	var me *Graph = &Graph{}
// 	me.vertices = len(matrix)
// 	me.adgeList = make([][]int, me.vertices)
// 	for i := 0; i < me.vertices; i++ {
// 		me.adgeList = append(me.adgeList, make([]int, 0))
// 	}
// 	me.makeAdjacencyList(matrix)
// 	return me
// }

func (graph Graph) getGraph(vertex [][]int) Graph {
	graph.vertices = len(vertex)
	graph.adgeList = make([][]int, graph.vertices)
	for i := 0; i < graph.vertices; i++ {
		graph.adgeList = append(graph.adgeList, make([]int, 0))
	}
	return graph
	// graph.makeAdjacencyList(vertex)
	// fmt.Println()
}

// Convert into adjacency list
// func (this *Graph) makeAdjacencyList(matrix [][]int) {
// 	for i := 0; i < this.vertices; i++ {
// 		for j := 0; j < this.vertices; j++ {
// 			if matrix[i][j] == 1 {
// 				this.addEdge(i, j)
// 			}
// 		}
// 	}
// }

// func (graph Graph) makeAdjacencyList(matrix [][]int) Graph {
// 	for i := 0; i < graph.vertices; i++ {
// 		for j := 0; j < graph.vertices; j++ {
// 			if matrix[i][j] == 1 {
// 				graph.addEdge(i, j)
// 			}
// 		}
// 	}
// 	return graph

// }

// func (this *Graph) addEdge(u, v int) {
// 	if u < 0 || u >= this.vertices ||
// 		v < 0 || v >= this.vertices {
// 		return
// 	}
// 	// Add node edge
// 	this.adgeList[u] = append(this.adgeList[u], v)
// }

// func (graph Graph) addEdge(u, v int) Graph {
// 	if u < 0 || u >= graph.vertices ||
// 		v < 0 || v >= graph.vertices {
// 		return graph
// 	}
// 	// Add node edge
// 	graph.adgeList[u] = append(graph.adgeList[u], v)
// }

// Display graph nodes and edges
// func (this Graph) printGraph() {
// 	fmt.Print("\n Graph Adjacency List ")
// 	for i := 0; i < this.vertices; i++ {
// 		fmt.Print(" \n [", i, "] :")
// 		// iterate edges of i node
// 		for j := 0; j < len(this.adgeList[i]); j++ {
// 			if j != 0 {
// 				fmt.Print(" → ")
// 			}
// 			fmt.Print(" ", this.adgeList[i][j])
// 		}
// 	}
// }

func (graph Graph) printGraph() {
	fmt.Print("\n Graph Adjacency List ")
	for i := 0; i < graph.vertices; i++ {
		fmt.Print(" \n [", i, "] :")
		// iterate edges of i node
		for j := 0; j < len(graph.adgeList[i]); j++ {
			if j != 0 {
				fmt.Print(" → ")
			}
			fmt.Print(" ", graph.adgeList[i][j])
		}
	}
}
func main() {
	var matrix = [][]int{
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 1, 0},
	}
	var g Graph
	y := g.getGraph(matrix)
	fmt.Println(y)
	// g.getGraph(matrix)
	// // Display graph element
	// g.printGraph()
}
