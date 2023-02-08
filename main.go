package main

import "github.com/nikhilnarayanan623/ds-binary-search-tree/bst"

func main() {

	b := bst.GetNewBinarySearchTree()

	b.Insert(5)
	b.Insert(3)
	b.Insert(78)
	b.Insert(1)
	b.Insert(45)
	b.Insert(53)
	b.Insert(33)
	b.Inorder()

	b.Remove(5)
	b.Remove(33)

	b.Inorder()
	b.PreOrder()
	b.PostOrder()
}
