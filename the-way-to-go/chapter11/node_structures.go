package main 

import "fmt"

type Node struct {
	le *Node
	ri *Node
	data interface{}
}

func (n *Node) String() string {
	return fmt.Sprintf("[%T] %v", n.data, n.data)
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func NewNode(left, right *Node) *Node {
	return &Node{left, right, nil}
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("Root Node")
	// make child (leaf) nodes:
	a := NewNode(nil, nil)
	b := NewNode(nil, nil)
	a.SetData("left Node")
	b.SetData(123)
	root.le = a
	root.ri = b
	fmt.Printf("root: %v\nroot.left: %v\nroot.right: %v\n", root, root.le, root.ri)
}
