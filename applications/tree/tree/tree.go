package tree

import (
	"dirtree/stack"
)

// Node tree node
type Node struct {
	ID       int
	Pid      int
	Title    string
	IsLeaf   bool
	Children []*Node
}

// Nodes tree nodes
type Nodes []*Node

// GenTree generate tree
func GenTree(nodes Nodes) Nodes {
	tree := Nodes{}
	nodeMap := make(map[int]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}
	for key, node := range nodeMap {
		if pnode, ok := nodeMap[node.Pid]; ok {
			pnode.Children = append(pnode.Children, node)
			pnode.IsLeaf = true
		} else {
			tree = append(tree, nodeMap[key])
		}
	}

	return tree
}

// FindNode find node by ID
func FindNode(tree Nodes, id int) *Node {
	stk := stack.NewStack()
	for _, node := range tree {
		stk.Push(node)
	}

	for stk.GetLen() != 0 {
		tmpNode := stk.Pop().(*Node)

		if tmpNode.ID == id {
			return tmpNode
		}

		if tmpNode.Children != nil {
			for _, child := range tmpNode.Children {
				stk.Push(child)
			}
		}
	}

	return new(Node)
}
