package main

import "fmt"

func main() {
    // 会在上层函数 (main) 返回后执行
    // 函数的参数会立刻生成 但是在上层函数返回前函数都不会被调用
  	defer fmt.Println("world")

  	fmt.Println("hello")
}
