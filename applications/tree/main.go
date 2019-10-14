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

	tree := tree.GenTree(nodes)

	treeJson, _ := json.Marshal(tree)
	fmt.Printf("%s", treeJson)
}
