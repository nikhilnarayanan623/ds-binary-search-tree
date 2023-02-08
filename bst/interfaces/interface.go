package interfaces

type BinarySearchTreeInterface interface {
	Insert(data int)        // to insert a new value to the binary search tree
	Remove(data int)        // to remove a value from binary search tree
	Contains(data int) bool // to check a value contain in binary search tree if it contain return true otherwise return false
	Inorder()               // print all value in the bst as order ( left -> root -> right  )
	PreOrder()              // print all values in the bst as order ( root -> left -> right)
	PostOrder()             // print all values int the bst as order ( left -> right ->root )
}
