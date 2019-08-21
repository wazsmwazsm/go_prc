package main

import (
	"fmt"
)

type Any interface{}

// EvalFunc 接收一个值, 返回当前值和满足规律的下一个值
type EvalFunc func(state Any) (Any, Any)

func main() {
	// 偶数生成的函数
	evenFunc := func(state Any) (Any, Any) {
		now := state.(int)
		next := now + 2
		return now, next
	}
	// 获取结果接收函数
	even := BuildLazyIntEvaluator(evenFunc, 0)

	// 生成前 10 个偶数
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

// BuildLazyEvaluator 惰性生成器，接收一个可以生成一系列数的函数 evalFunc
// 返回结果接收函数
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	// 生成
	go func() {
		var actState = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState) // 根据初始值生成当前值和下一个值
			retValChan <- retVal
		}
	}()
	// 返回接收函数
	return func() Any {
		return <-retValChan
	}
}

// BuildLazyIntEvaluator 包装 BuildLazyEvaluator，返回一个返回 int 结果的结果接收函数
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
