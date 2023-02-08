package bst

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-binary-search-tree/bst/interfaces"
)

// struct for hold single value as a node
type node struct {
	data  int
	left  *node
	right *node
}

// struct to hold whole nodes as a binary search tree and its functions
type binarySearchTree struct {
	root *node
}

// func to get new instance of binary search tree
func GetNewBinarySearchTree() interfaces.BinarySearchTreeInterface {
	return &binarySearchTree{}
}

// func to insert a value to binary search tree
func (bst *binarySearchTree) Insert(data int) {
	//fist check the root is nill then create root
	if bst.root == nil {
		bst.root = &node{data: data}
		return
	}

	//head is not nil then traverse throgh the binary search tree and add on it
	currentNode := bst.root

	for { // in here value is equal to currentNode then add new node on its right
		if data < currentNode.data { // data less than currentNode then check left
			if currentNode.left == nil { // and if its left is nill then create a node at left
				currentNode.left = &node{data: data}
				return
			} // left is not nil move to left
			currentNode = currentNode.left
		} else { // data is greater than currentNode the check right
			if currentNode.right == nil { // if right is nil then create a node at right
				currentNode.right = &node{data: data}
				return
			} // rigth is not nil then move to right
			currentNode = currentNode.right
		}
	}
}

// func to remove a value from the binary search tree
func (bst *binarySearchTree) Remove(value int) {
	//check the bst is empty or not
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	// find the node to delete if its present
	currentNode := bst.root

	for currentNode != nil {

		if value < currentNode.data { // value less than currentNode then check left

			// if value found at left then call helper to delete that node and assign there appropriate node
			if currentNode.left != nil && currentNode.left.data == value {
				currentNode.left = removeHelper(currentNode.left)
				return
			} // left value not the value then move to left
			currentNode = currentNode.left

		} else if value > currentNode.data { //value greater than currentNode then check right

			// if right is not nill and value found at right node then pass right node removeHelper to delet that node and assign there appropriate node
			if currentNode.right != nil && currentNode.right.data == value {
				currentNode.right = removeHelper(currentNode.right)
				return
			} //right not the value then move to right
			currentNode = currentNode.right

		} else { // only if root have the value then only currentNode value is same as the value to check
			// all other time we check left and right value in forward
			bst.root = removeHelper(bst.root)
			return
		}
	}
	fmt.Println("Can't find the value to remove")
}

// func to help remove that find appropriate node for replacement for delete node
// delete approach
// 1. -> go to right
// 2. -> then its last left
func removeHelper(removeNode *node) *node {
	removeNodeLeft := removeNode.left
	removeNodeRight := removeNode.right

	//if given node's right is nil then return currentNode's left
	if removeNode.right == nil {
		return removeNodeLeft
	} else if removeNode.right.left == nil { // removeNode's right left is nill then connect right and left and retun right
		removeNodeRight.left = removeNodeLeft // connect rigts' left as remove node's left and return right
		return removeNodeRight
	}

	// now removeNode's rightNode have left then find last left node
	prevNode := removeNodeRight // to hold previousNode

	currentNode := removeNodeRight.left // to find last left node

	for currentNode.left != nil { // loop until find last left node
		prevNode = currentNode         // find last prev node
		currentNode = currentNode.left // last node
	}
	//connect last left nodes right to previous node's left
	// chance last node have right node is present
	prevNode.left = currentNode.right

	// now the currentNode is last left node
	// so connect last node's left as removeNode's left and right as removeNode's right
	currentNode.left = removeNodeLeft
	currentNode.right = removeNodeRight

	return currentNode
}

// func to check a value is contain in binary search treee
func (bst *binarySearchTree) Contains(value int) bool {

	// traverse throght the tree in binary search model and check the value is there or not
	currentNode := bst.root

	for currentNode != nil {

		if value < currentNode.data { // value is less to currentNode value then move left
			currentNode = currentNode.left
		} else if value > currentNode.data { // value is greater then move to right
			currentNode = currentNode.right
		} else { // value is equal to the value then found the value and return true
			return true
		}
	}
	// head is nil(bst is empty) or value not found
	return false
}

// func to print all values in Inorder
// left -> root -> right
func (bst *binarySearchTree) Inorder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	InorderHelper(bst.root)
	fmt.Println()
}

// func to help recursion on Inorder
func InorderHelper(node *node) {

	if node == nil {
		return
	}

	//go to left
	InorderHelper(node.left)
	// print the node value at last
	fmt.Print(node.data, " ")
	// go to right
	InorderHelper(node.right)
}

// func to print bst values in preOrder
// root -> left -> right
func (bst *binarySearchTree) PreOrder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	preOrderHelper(bst.root)
	fmt.Println()
}

// func to help recursion for preOrder
func preOrderHelper(node *node) {

	if node == nil {
		return
	}

	//print the node value
	fmt.Print(node.data, " ")
	//then go to left
	preOrderHelper(node.left)
	//then go to right
	preOrderHelper(node.right)
}

// print values in postOrder
// left -> right -> root

func (bst *binarySearchTree) PostOrder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	postOrderHelper(bst.root)
	fmt.Println()
}

// func to help recursion for postOrder
func postOrderHelper(node *node) {
	if node == nil {
		return
	}

	// go to left
	postOrderHelper(node.left)
	// then go to left
	postOrderHelper(node.right)
	//last print value
	fmt.Print(node.data, " ")
}
