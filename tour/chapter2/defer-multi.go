package main

import "fmt"

func main() {
	 fmt.Println("counting")

   // 延迟的函数调用被压入一个栈中。
   // 当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
   for i := 0; i < 10; i++ {
      defer fmt.Println(i)
   }

   fmt.Println("mian done")
}
