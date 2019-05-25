package main

import (
	"fmt"
	"code.google.com/p/go-tour/tree"
)

/*
	
	程序功能 : 对比两个结构不同的二叉树是否存储了相同的序列

	tree.New(k) 构造了一个随机结构的二叉树, 保存了值 k ， 2k ， 3k ， ... ， 10k
	而且保证了嵌套遍历后的顺序一样 (不然就没法判断了)

*/

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// 嵌套遍历二叉树的各个节点
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// 关闭 channel
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	// 循环读取 w1 w2 channel 的输出
	for {
		v1, ok1 := <- w1
		v2, ok2 := <- w2

		// 如果 channel 读完, 返回成功结果
		if ! ok1 || ! ok2 {
			return ok1 == ok2
		}
		// 如果不相等, 返回错误结果
		// 这里直接返回, Walk goroutine 还在阻塞中 (没有遍历完, 阻塞到 ch <- t.Value 上)
		// 会造成 goroutine 的泄漏, 需要正确的退出 goroutine 才行
		// 见 quit 版本
		if v1 != v2 {
			return false
		}
	}

}


func main() {

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if ! Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

}
