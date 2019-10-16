package main

import (
	"dirtree/tree"
	"encoding/json"
	"fmt"
)

func main() {
	nodes := []*tree.Node{
		{ID: 1, Pid: 0, Title: "root", Children: []*tree.Node{}},
		{ID: 2, Pid: 1, Title: "aa", Children: []*tree.Node{}},
		{ID: 3, Pid: 1, Title: "bb", Children: []*tree.Node{}},
		{ID: 4, Pid: 2, Title: "cc", Children: []*tree.Node{}},
		{ID: 5, Pid: 2, Title: "dd", Children: []*tree.Node{}},
		{ID: 6, Pid: 2, Title: "ee", Children: []*tree.Node{}},
		{ID: 7, Pid: 3, Title: "ff", Children: []*tree.Node{}},
		{ID: 8, Pid: 1, Title: "ff", Children: []*tree.Node{}},
	}

	t := tree.NewTree(nodes)
	root := t.GetRoot()
	treeJSON, _ := json.Marshal(root)
	fmt.Printf("tree:\n %s\n", treeJSON)

	node := t.Find(2)
	nodeJSON, _ := json.Marshal(node)
	fmt.Printf("node:\n %s\n", nodeJSON)

	pnode := t.FindParent(2)
	pnodeJSON, _ := json.Marshal(pnode)
	fmt.Printf("pnode:\n %s\n", pnodeJSON)

	chain := t.GetNodeChain(1, 6)
	chainJSON, _ := json.Marshal(chain)
	fmt.Printf("chain:\n %s\n", chainJSON)

	childrens := t.FindChildrens(2)
	childrensJSON, _ := json.Marshal(childrens)
	fmt.Printf("childrens:\n %s\n", childrensJSON)

	subs := t.GetSubNodes(2)
	subsJSON, _ := json.Marshal(subs)
	fmt.Printf("subs:\n %s\n", subsJSON)

	/* ---- other tree ---- */
	subs = []*tree.Node{}
	subs3 := t.GetSubNodes(3)
	subs4 := t.GetSubNodes(4)
	subs6 := t.GetSubNodes(6)
	chain3 := t.GetNodeChain(0, 3)
	chain4 := t.GetNodeChain(0, 4)
	chain6 := t.GetNodeChain(0, 6)
	nodeMap := make(map[int]*tree.Node)
	apd := func(elem []*tree.Node) []*tree.Node {
		for _, sub := range elem {
			if _, ok := nodeMap[sub.ID]; !ok {
				nodeMap[sub.ID] = sub
				subs = append(subs, sub)
			}
		}
		return subs
	}
	apd(subs3)
	apd(subs4)
	apd(subs6)
	apd(chain3)
	apd(chain4)
	apd(chain6)

	t2 := tree.NewTree(subs)
	root2 := t2.GetRoot()
	tree2JSON, _ := json.Marshal(root2)
	fmt.Printf("tree2:\n %s\n", tree2JSON)
}
