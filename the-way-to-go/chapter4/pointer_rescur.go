package main

// 递归中传指针的赋值的问题
// 这里拿结构体指针做例子，slice 等同理
import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	t := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
				Right: &Node{
					Val: 5,
				},
			},
		},
		Right: &Node{
			Val: 3,
		},
	}

	p := &Node{}

	helper(t, p)
	fmt.Println(p)
}

func helper(root *Node, p *Node) {
	if root == nil {
		return
	}

	if root.Val == 4 {
		// p = root // p 是个形参，赋值的话每次递归调用都会被刷新
		*p = *root // 这是正确的做法
	}

	helper(root.Left, p)
	helper(root.Right, p)
}
