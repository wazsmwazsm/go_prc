package main

import (
	"fmt"
)

type Any interface{}

// EvalFunc 接收一个值, 返回当前值和满足规律的下一个值
type EvalFunc func(state Any) (Any, Any)

func main() {
	// fibonacci 生成的函数
	fibFunc := func(state Any) (Any, Any) {
		now := state.([]uint64)
		next := []uint64{now[1], now[0] + now[1]}
		return now, next
	}
	// 获取结果接收函数
	fib := BuildLazyUInt64Evaluator(fibFunc, []uint64{0, 1})

	// 生成前 10 个 fibonacci 数
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fib: %v\n", i, fib()[0])
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

// BuildLazyIntEvaluator 包装 BuildLazyEvaluator，返回一个返回 uint64 结果的结果接收函数
func BuildLazyUInt64Evaluator(evalFunc EvalFunc, initState Any) func() []uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() []uint64 {
		return ef().([]uint64)
	}
}
