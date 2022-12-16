package main

import (
	"fmt"
)

type Node struct {
	_element int
	_left    *Node
	_right   *Node
}

type binaryTree struct {
	_root *Node
}

func (t *binaryTree) maketree(e int, left, right *binaryTree) {
	t._root = &Node{e, left._root, right._root}
	//fmt.Println(t._root)
}

func (t *binaryTree) inorder(troot *Node) {
	if troot != nil {
		t.inorder(troot._left)
		fmt.Print(troot._element, " ")
		t.inorder(troot._right)
	}
}

func (t *binaryTree) preorder(troot *Node) {
	if troot != nil {
		fmt.Print(troot._element, " ")
		t.preorder(troot._left)
		t.preorder(troot._right)
	}
}

func (t *binaryTree) postorder(troot *Node) {
	if troot != nil {
		t.postorder(troot._left)
		t.postorder(troot._right)
		fmt.Print(troot._element, " ")
	}
}

func (t *binaryTree) levelorder(troot *Node) {
	Q := []*Node{}
	a := troot
	fmt.Print(troot._element, " ")
	Q = append(Q, a)
	for len(Q) > 0 {
		a = Q[0]
		Q = Q[1:]
		if a._left != nil {
			fmt.Print(a._left._element, " ")
			Q = append(Q, a._left)
		}
		if a._right != nil {
			fmt.Print(a._right._element, " ")
			Q = append(Q, a._right)
		}
	}
}

func (t *binaryTree) count(troot *Node) int {
	if troot != nil {
		x := t.count(troot._left)
		y := t.count(troot._right)
		return x + y + 1
	}
	return 0
}

func (t *binaryTree) height(troot *Node) int {
	if troot != nil {
		x := t.height(troot._left)
		y := t.height(troot._right)
		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return 0
}

func main() {
	x := binaryTree{}
	y := binaryTree{}
	z := binaryTree{}
	r := binaryTree{}
	s := binaryTree{}
	t := binaryTree{}
	a := binaryTree{}
	x.maketree(40, &a, &a)
	y.maketree(60, &a, &a)
	z.maketree(20, &x, &a)
	r.maketree(50, &a, &y)
	s.maketree(30, &r, &a)
	t.maketree(10, &z, &s)
	fmt.Println("Inorder Traversal")
	t.inorder(t._root)
	fmt.Println()
	fmt.Println("Preorder Traversal")
	t.preorder(t._root)
	fmt.Println()
	fmt.Println("Postorder Traversal")
	t.postorder(t._root)
	fmt.Println()
	fmt.Println("Levelorder Traversal")
	t.levelorder(t._root)

	
	fmt.Println()
	fmt.Println("Number of Nodes :", t.count(t._root))
	fmt.Println("Height of the Tree :", t.height(t._root)-1)
	t.inorder(t._root)
	fmt.Println()

}