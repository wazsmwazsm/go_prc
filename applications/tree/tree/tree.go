package tree

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
