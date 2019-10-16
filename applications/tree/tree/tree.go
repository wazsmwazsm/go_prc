package tree

// import (
// 	"fmt"
// )

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
	nodeMap map[int]*Node
	root    *Node
}

// NewTree create tree
func NewTree(nodes []*Node) *Tree {
	root := &Node{ID: 0, Pid: -1, Title: "root"}
	nodes = append(nodes, root)
	nodeMap := make(map[int]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	for _, node := range nodeMap {
		if pnode, ok := nodeMap[node.Pid]; ok {
			pnode.Children = append(pnode.Children, node)
			node.parent = pnode
		}
	}
	return &Tree{
		nodeMap: nodeMap,
		root:    root,
	}
}

// GetRoot get root
func (t *Tree) GetRoot() *Node {
	return t.root
}

// Find find node by ID
func (t *Tree) Find(id int) *Node {
	if node, ok := t.nodeMap[id]; ok {
		return node
	}

	return nil
}

// FindParent find parent node by ID (bfs)
func (t *Tree) FindParent(id int) *Node {

	if node := t.Find(id); node != nil {
		return node.parent
	}

	return nil
}

// FindChildrens find children nodes by ID
func (t *Tree) FindChildrens(id int) []*Node {

	if node := t.Find(id); node != nil {
		return node.Children
	}

	return nil
}

// GetSubNodes get sub nodes (contain itself) flat by ID (bfs)
func (t *Tree) GetSubNodes(id int) []*Node {
	subs := []*Node{}
	node := t.Find(id)
	if node == nil {
		return subs
	}

	queue := []*Node{}
	queue = append(queue, node)

	for len(queue) != 0 {
		tmpNode := queue[0]
		queue = queue[1:]
		subs = append(subs, &Node{
			ID:       tmpNode.ID,
			Pid:      tmpNode.Pid,
			Title:    tmpNode.Title,
			Children: []*Node{},
		})
		if tmpNode.Children != nil {
			for _, child := range tmpNode.Children {
				queue = append(queue, child)
			}
		}
	}

	return subs
}

// GetNodeChain get nodes chain flat between node1 to node2
func (t *Tree) GetNodeChain(fromID, toID int) []*Node {

	chain := []*Node{}
	toNode := t.Find(toID)

	currNode := toNode
	for currNode != nil {
		chain = append(chain, &Node{
			ID:       currNode.ID,
			Pid:      currNode.Pid,
			Title:    currNode.Title,
			Children: []*Node{},
		})

		if currNode.ID == fromID {
			return chain // end
		}
		pnode := currNode.parent
		currNode = pnode
	}

	return []*Node{} // path doesn't exist
}
