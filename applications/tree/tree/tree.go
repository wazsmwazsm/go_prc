package tree

// Node tree node
type Node struct {
	ID       int
	Pid      int
	Title    string
	parent   *Node
	Children []*Node
}

// Tree multi tree
type Tree struct {
	nodes   []*Node
	nodeMap map[int]*Node
	root    *Node
}

// NewTree create tree
func NewTree(nodes []*Node) *Tree {
	var root *Node
	nodeMap := make(map[int]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	for key, node := range nodeMap {
		if pnode, ok := nodeMap[node.Pid]; ok {
			pnode.Children = append(pnode.Children, node)
			node.parent = pnode
		} else {
			root = nodeMap[key]
		}
	}

	return &Tree{
		nodes:   nodes,
		nodeMap: nodeMap,
		root:    root,
	}
}

// GetRoot get root
func (t *Tree) GetRoot() *Node {
	return t.root
}

// FindNode find node by ID (bfs)
func (t *Tree) FindNode(id int) *Node {
	queue := []*Node{}
	queue = append(queue, t.root)

	for len(queue) != 0 {
		tmpNode := queue[0]
		queue = queue[1:]

		if tmpNode.ID == id {
			return tmpNode
		}

		if tmpNode.Children != nil {
			for _, child := range tmpNode.Children {
				queue = append(queue, child)
			}
		}
	}

	return nil
}

// FindPNode find parent node by ID (bfs)
func (t *Tree) FindPNode(id int) *Node {

	if node := t.FindNode(id); node != nil {
		return node.parent
	}

	return nil
}

// func (t *Tree) GetPath(id int) (root *Node) {

// }

func (t *Tree) GetPathBetweenNode(fromID, toID int) Node {

	toNode, ok := t.nodeMap[toID]
	if !ok {
		return Node{}
	}

	pid := toNode.Pid
	currNode := toNode
	for pid != fromID {

		pnode, ok := t.nodeMap[pid]
		if !ok {
			return *pnode
		}

		pnode.Children = append(pnode.Children, currNode)
		currNode = pnode
	}

	return *currNode
}
