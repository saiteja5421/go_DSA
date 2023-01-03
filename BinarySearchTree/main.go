package main

import (
	"fmt"
)

type Node struct {
	_element int
	_left    *Node
	_right   *Node
}

type binarysearchtree struct {
	_root *Node
}

func (t *binarysearchtree) search(key int) bool {
	troot := t._root
	for troot != nil {
		if troot._element == key {
			return true
		} else if key < troot._element {
			troot = troot._left
		} else if key > troot._element {
			troot = troot._right
		}
	}
	return false
}

func (t *binarysearchtree) insert(troot *Node, e int) {
	var temp *Node
	for troot != nil {
		temp = troot
		if e == troot._element {
			return
		} else if e < troot._element {
			troot = troot._left
		} else if e > troot._element {
			troot = troot._right
		}
	}
	n := &Node{e, nil, nil}
	if t._root != nil {
		if e < temp._element {
			temp._left = n
		} else {
			temp._right = n
		}
	} else {
		t._root = n
	}
}

func (t *binarysearchtree) delete(e int) {
	var pp *Node
	p := t._root
	for p != nil && p._element != e {
		pp = p
		if e < p._element {
			p = p._left
		} else {
			p = p._right
		}
	}

	if p == nil {
		return
	}

	if p._left != nil && p._right != nil {
		s := p._left
		ps := p
		for s._right != nil {
			ps = s
			s = s._right
		}
		p._element = s._element
		p = s
		pp = ps
	}

	var c *Node
	if p._left != nil {
		c = p._left
	} else {
		c = p._right
	}

	if p == t._root {
		t._root = nil
	} else {
		if p == pp._left {
			pp._left = c
		} else {
			pp._right = c
		}
	}

}

func (t *binarysearchtree) inorder(troot *Node) {
	if troot != nil {
		t.inorder(troot._left)
		fmt.Print(troot._element, " ")
		t.inorder(troot._right)
	}
}
func (t *binarysearchtree) preorder(troot *Node) {
	if troot != nil {
		fmt.Print(troot._element, " ")
		t.preorder(troot._left)
		t.preorder(troot._right)
	}
}

func (t *binarysearchtree) postorder(troot *Node) {
	if troot != nil {
		t.postorder(troot._left)
		t.postorder(troot._right)
		fmt.Print(troot._element, " ")
	}
}


func (t *binarysearchtree)count(troot *Node) int {
      if troot != nil{
		  x := t.count(troot._left)
		  y := t.count(troot._right)
		  return x+y+1
	  }
	  return 0
}



func main() {
	x := binarysearchtree{}
	x.insert(x._root, 50)
	x.insert(x._root, 40)
	x.insert(x._root, 60)
	x.insert(x._root, 20)
	x.insert(x._root, 10)
	//x.delete(60)
	x.inorder(x._root)
	fmt.Println()
	x.preorder(x._root)
	fmt.Println()
	a := x.search(40)
	fmt.Println(a)
	x.postorder(x._root)
	fmt.Println()
	b := x.count(x._root)
	fmt.Println(b)
	
}
