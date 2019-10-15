package main

import (
	"dirtree/tree"
	"encoding/json"
	"fmt"
)

func main() {
	nodes := []*tree.Node{
		{ID: 1, Pid: 0, Title: "root"},
		{ID: 2, Pid: 1, Title: "aa"},
		{ID: 3, Pid: 1, Title: "bb"},
		{ID: 4, Pid: 2, Title: "cc"},
		{ID: 5, Pid: 2, Title: "dd"},
		{ID: 6, Pid: 2, Title: "ee"},
		{ID: 7, Pid: 3, Title: "ff"},
	}

	t := tree.NewTree(nodes)
	root := t.GetRoot()
	treeJSON, _ := json.Marshal(root)
	fmt.Printf("%s\n", treeJSON)

	node := t.FindNode(7)
	nodeJSON, _ := json.Marshal(node)
	fmt.Printf("%s\n", nodeJSON)

	path := t.GetPathBetweenNode(1, 3)
	pathJSON, _ := json.Marshal(path)
	fmt.Printf("%s\n", pathJSON)
}
