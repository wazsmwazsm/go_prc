package trans

import "math"

var Pi float64

// 用来初始化

// 不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
// 每一个源文件都可以包含一个或多个 init 函数
// 初始化总是以单线程执行，并且按照包的依赖关系顺序执行
func init() {
	Pi = 4 * math.Atan(1)
}
