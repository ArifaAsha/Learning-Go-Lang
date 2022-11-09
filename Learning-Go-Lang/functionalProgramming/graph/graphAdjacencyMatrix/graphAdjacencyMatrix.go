package main

import "fmt"

func addVertex(graph [][]int, v int) [][]int {

	// if(graph[i]!= v){

	// }else{
	// if(){
	graph2 := [][]int{}

	emptySlice := []int{}
	graph[0] = append(emptySlice, v)
	graph2 = append(graph2)
	return graph

}

// func findVertex(graph [][]int, from int) ([][]int, []int) {
// 	for _, col := range graph {
// 		for _, value := range col {
// 			if value == from {
// 				return graph, col
// 			}
// 		}
// 	}
// 	addVertex(graph, from)
// 	return graph, nil
// }

func findVertex(graph [][]int, from int, to int) ([][]int, []int) {
	for _, col := range graph {
		for _, value := range col {
			fmt.Println(value)
			// if value == from {
			// 	fmt.Println("hehe")
			// 	// col = append(col, to)
			// 	// fmt.Println(col)
			// }
		}

	}
	addVertex(graph, from)
	return graph, nil
}

// func addNewEdgeInGraph(graph [][]int, vertex []int, vertex2 []int) {
// 	for _, col := range graph {
// 		for _, value := range col {
// 			if value == from {
// 				append()
// 			}
// 		}
// 	}
// }

// func addEdge(graph [][]int, from int, to int) [][]int {
// 	graph, vertex := findVertex(graph, from)
// 	vertex2 := append(vertex, to)
// 	addNewEdgeInGraph(graph, vertex, vertex2)
// 	return graph
// }

// main function
func main() {

	// declaring a slice of slices of
	// type integer with a length of 3
	var graph = make([][]int, 4)
	graph = addVertex(graph, 0)
	graph = addVertex(graph, 1)
	graph = addVertex(graph, 2)
	graph = addVertex(graph, 3)

	fmt.Println(graph)

	// x := checkIfExist2D(graph, 0)
	// fmt.Println(graph)
	graph, x := findVertex(graph, 2, 3)
	// fmt.Println(graph)
	// fmt.Println(x)
	// graph = addEdge(graph, 2, 3)
	// fmt.Println(graph)

}
