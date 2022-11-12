package main

import "fmt"

func root(tree []int, key int) []int {
	// if len(tree) != 0 {
	// 	fmt.Println("---Tree already had root--")
	// } else {
	// tree = append(tree, key)
	tree[0] = key
	// }
	return tree
}

func set_left(key int, parent int, tree []int) ([]int, int) {
	// trackParent = append(trackParent, parent)
	// if tree[parent] == 0 {
	// 	print("Can't set child at", (parent*2)+1, ", no parent found")
	// } else {
	index := (parent * 2) + 1
	tree[index] = key
	// }
	return tree, index
}

func set_right(key int, parent int, tree []int) ([]int, int) {
	// trackParent = append(trackParent, parent)
	// if tree[parent] == 0 {
	// 	print("Can't set child at", (parent*2)+2, ", no parent found")
	// } else {
	index := (parent * 2) + 2
	tree[index] = key
	// }
	return tree, index
}

// func printTree(tree []int, parentRoot int) {
// 	if len(tree) != 0 {
// 		fmt.Println(tree[(parentRoot*2)+1])
// 		if tree[(parentRoot*2)+1] != 0 && tree[(parentRoot*2)+1] <= 9 {
// 			fmt.Printf("left of %d is = ", parentRoot)
// 			fmt.Printf("%d\n", tree[(parentRoot*2)+1])
// 		}
// 		if tree[(parentRoot*2)+2] != 0 && tree[(parentRoot*2)+1] <= 9 {
// 			fmt.Printf("right of %d is = ", parentRoot)
// 			fmt.Printf("%d\n", tree[(parentRoot*2)+2])
// 		}
// 	} else {
// 		parentRoot++
// 		printTree(tree[1:], parentRoot)
// 	}

// }

func printTree(tree []int, eachNode int, indexOFNode int) {
	if len(tree) != 0 && eachNode != 0 {
		left := getLeftChild(tree, eachNode)

		if left != -1 {
			fmt.Printf("Left child of %d is = ", tree[])
			fmt.Println(left)
		}

		right := getRightChild(tree, eachNode)
		if right != -1 {
			fmt.Printf("Right child of %d is = ", eachNode)
			fmt.Println(right)
		}
	}
}

func traverseTree(tree []int, indexOfNode int) {
	if len(tree) == 0 {
		return
	} else {
		// fmt.Print(tree[0], "  ")
		printTree(tree, tree[0], indexOfNode)
		indexOfNode++
		traverseTree(tree[1:], indexOfNode)
	}
}

func getRightChild(tree []int, index int) int {
	// node is not null
	// and the result must lie within the number of nodes for a complete binary tree
	indexOfRightChild := (2 * index) + 2
	if len(tree) != 0 && indexOfRightChild <= 9 && tree[indexOfRightChild] != 0 {
		return tree[indexOfRightChild]
	}
	// right child doesn't exist
	return -1
}

func getLeftChild(tree []int, index int) int {
	// node is not null
	// and the result must lie within the number of nodes for a complete binary tree
	indexOfLeftChild := (2 * index) + 1
	if len(tree) != 0 && indexOfLeftChild <= 9 && tree[indexOfLeftChild] != 0 {
		return tree[indexOfLeftChild]
	}
	// right child doesn't exist
	return -1
}

// func inorder(tree []int, index int) {
// 	// checking for valid index and null node
// 	if index >= 0 && len(tree) != 0 {
// 		left := getLeftChild(tree, index)
// 		if left != -1 {
// 			fmt.Printf("Left child of %d is = ", tree[index])
// 			fmt.Println(tree[left])
// 			inorder(tree, left) //visiting left subtree
// 		}
// 		// fmt.Printf(" %d ", tree[index]) //visiting root
// 		right := getRightChild(tree, index)
// 		if right != -1 {
// 			fmt.Printf("Right child of %d is = ", tree[index])
// 			fmt.Println(tree[right])
// 			inorder(tree, right) // visiting right subtree
// 		}

// 	}
// }

// func printTree(tree []int, root int) {
// 	fmt.Printf("  ")
// 	fmt.Println(tree[root])
// 	fmt.Println("/   \\")
// 	fmt.Print(tree[(2*1)+1])
// 	fmt.Printf("   ")
// 	fmt.Print(tree[(2*1)+2])
// 	// fmt.Print("\\")

// }

// func printRight(tree []int, node int) {
// 	fmt.Println(tree[(2*1)+2])
// }

// func print_tree(tree []int) {
// 	for i := 0; i < len(tree); i++ {
// 		if tree[i] != 0 {
// 			fmt.Print(tree[i], "")
// 		} else {
// 			fmt.Print("-", "")
// 		}
// 		fmt.Println()

// 	}
// }

// func print_tree(tree []int) {
// 	for i := 0; i < len(tree); i++ {
// 		if tree[i] != 0 {
// 			fmt.Print(tree[i], "")
// 			if tree[(i*2)+1] != 0 {
// 				fmt.Printf("left of %d = %d", tree[i], tree[(i*2)+1])
// 			}

// 		} else {
// 			fmt.Print("-", "")
// 		}
// 		fmt.Println()

// 	}
// }

func main() {
	// noOfNodes := 10
	tree := make([]int, 10)
	indexOfNode := 0
	// root(tree, 3)

	// fmt.Println(tree)
	tree = root(tree, 1)
	tree, indexOfNode = set_left(2, 1, tree)
	tree, indexOfNode = set_right(3, 1, tree)

	tree, indexOfNode = set_left(4, 2, tree)
	tree, indexOfNode = set_right(5, 2, tree)

	tree, indexOfNode = set_left(6, 3, tree)
	tree, indexOfNode = set_right(7, 3, tree)

	// x := getRightChild(tree, 1)
	// y := getLeftChild(tree, 1)
	// fmt.Println(x)
	// fmt.Println(y)

	// traverseTree(tree)
	// fmt.Println()
	// fmt.Println(tree)
	// x := getLeftChild(tree, 2)
	// fmt.Println(x)

	fmt.Println(tree, 0)
	// traverseTree(tree)
	// inorder(tree, 1)

	// fmt.Println(tree)

	// traverseTree(tree)
	// printTree(tree, 1)
	// print_tree(tree)
	// printTree(tree, 0)

}
