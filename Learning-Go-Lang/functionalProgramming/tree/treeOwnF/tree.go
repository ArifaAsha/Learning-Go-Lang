package main

import "fmt"

// set root of the tree
func root(tree []int, key int) []int {
	tree[0] = key
	return tree
}

// left = (2*parent)+1
func set_left(key int, parent int, tree []int) []int {
	tree[(parent*2)+1] = key
	return tree
}

// right = (2*parent)+2
func set_right(key int, parent int, tree []int) []int {
	tree[(parent*2)+2] = key
	return tree
}

func getRightChild(tree []int, index int) int {
	indexOfRightChild := (2 * index) + 2
	if len(tree) != 0 && indexOfRightChild <= 9 && tree[indexOfRightChild] != 0 { //tree[indexOfRightChild]  = 0 => no node is assigned
		return tree[indexOfRightChild]
	}
	return -1
}

func getLeftChild(tree []int, index int) int {
	indexOfLeftChild := (2 * index) + 1
	if len(tree) != 0 && indexOfLeftChild <= 9 && tree[indexOfLeftChild] != 0 {
		return tree[indexOfLeftChild]
	}
	return -1
}

func traverseTree(actualTree []int, tree []int, indexOfNode int, indexAtActualTree int) {
	if len(tree) == 0 {
		return
	} else {
		// indexOfNode = 0
		indexAtActualTree++

		if indexOfNode < 9 {
			val := tree[0]
			printTree(actualTree, tree, val, indexAtActualTree)
			traverseTree(actualTree, tree[1:], indexOfNode, indexAtActualTree)

		} else {
			return
		}
	}
}

func printTree(actualTree []int, tree []int, valAtIndex int, indexAtActualTree int) {
	if valAtIndex != 0 {
		left := getLeftChild(actualTree, actualTree[indexAtActualTree])
		right := getRightChild(actualTree, actualTree[indexAtActualTree])

		if left != -1 && left < 10 { // -1 = no left child
			fmt.Printf("Left child of %d is = ", actualTree[indexAtActualTree])
			fmt.Println(left)
		}
		if right != -1 && left < 10 { // -1 = no right child
			fmt.Printf("Right child of %d is = ", actualTree[indexAtActualTree])
			fmt.Println(right)
		}
	}

}

func main() {

	tree := make([]int, 10)

	tree = root(tree, 1)
	tree = set_left(2, 1, tree)
	tree = set_right(3, 1, tree)

	tree = set_left(4, 2, tree)
	tree = set_right(5, 2, tree)

	tree = set_left(6, 3, tree)
	tree = set_right(7, 3, tree)

	fmt.Print("\n---------------------\n")
	fmt.Print("Tree : ")
	fmt.Print(tree, "\n---------------------\n")

	// printTree(tree, 0)
	traverseTree(tree, tree, -1, -1)
	fmt.Println()

}
