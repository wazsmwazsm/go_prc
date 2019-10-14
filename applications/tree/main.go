package main

import (
	"dirtree/tree"
	"encoding/json"
	"fmt"
)

func main() {
	nodes := tree.Nodes{
		{ID: 1, Pid: 0, Title: "root"},
		{ID: 2, Pid: 1, Title: "aa"},
		{ID: 3, Pid: 1, Title: "bb"},
		{ID: 4, Pid: 2, Title: "cc"},
		{ID: 5, Pid: 2, Title: "dd"},
		{ID: 6, Pid: 2, Title: "ee"},
		{ID: 7, Pid: 3, Title: "ff"},
	}

	t := tree.GenTree(nodes)

	treeJson, _ := json.Marshal(t)
	fmt.Printf("%s\n", treeJson)

	node := tree.FindNode(t, 4)
	nodeJson, _ := json.Marshal(node)
	fmt.Printf("%s\n", nodeJson)
}
