package main

import "fmt"

const Alphabet_size = 26

type TrieNode struct {
	childNode    [Alphabet_size]*TrieNode //Array of child node, pointer to the node
	isEndOfWords bool
}

func createNode() *TrieNode { // node constructor -> returns a pointer; set of default values
	node := &TrieNode{} // create pointer to the new node
	node.isEndOfWords = false

	for i := 0; i < Alphabet_size; i++ { // assigning 0 value to the whole array
		node.childNode[i] = nil
	}

	return node
}

func insert(root *TrieNode, key string) { // root node -> to start a new key, new string to start
	currentNode := root // examine as traverse through trie; for each char look -> char represented in the array, if not create new node

	for i := 0; i < len(key); i++ { //looping till the length of inserting key
		index := key[i] - 'a'                    // ascii val of the char and use as index of the array
		if currentNode.childNode[index] == nil { //if no char after this node
			currentNode.childNode[index] = createNode() //create a new node
		}
		currentNode = currentNode.childNode[index] //if next node of the current node exists, go to the next node for checking
	}

	currentNode.isEndOfWords = true //end of word -> outside the loop
}

func search(root *TrieNode, key string) bool {
	currentNode := root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if currentNode.childNode[index] != nil { //if current node has child node
			currentNode = currentNode.childNode[index] //go to the next char (child node)
		} else {
			return false
		}
	}

	if currentNode != nil && currentNode.isEndOfWords {
		return true
	} else {
		return false
	}

}

func main() {
	words := []string{"bags", "baggage", "banner", "box", "cloths"}

	root := createNode()

	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}

	fmt.Println("box?", search(root, "box"))
	fmt.Println("cloths?", search(root, "cloths"))
	fmt.Println("baggage?", search(root, "baggage"))
	fmt.Println("ba?", search(root, "ba"))
}
