package main

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1000000)
	},
}

/*
	查看 gc 信息
	GODEBUG=gctrace=1 go run test.go

	gc 信息解释：

	执行次数  总执行时间   gc占比    STW 清扫时间, 并发标记和扫描的时间，STW标记的时间
 gc 412      @8.140s     0%:       0+0.97+0 ms clock,

	垃圾回收占用cpu时间  堆的大小，gc后堆的大小，存活堆的大小   整体堆的大小    使用的处理器数量
	0+0/0/0+0 ms cpu,   4->4->0 MB,                         5 MB goal,    12 P

	系统内存回收信息解释：

       使用多少M内存  剩下要清除的内存   系统映射的内存   释放的系统内存    申请的系统内存
scvg0: inuse: 426,   idle: 0,          sys: 427,      released: 0,     consumed: 427 (MB)
*/

/*
	pprof
	http://127.0.0.1:8080/debug/pprof/ 查看调试信息
	http://127.0.0.1:8080/debug/pprof/heap 查看堆分配

	go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap > heap.svg 生成堆分配图

*/

/*
	逃逸分析 go build --gcflags=-m [filename]
*/
func main() {

	http.HandleFunc("/test", test)
	// 开启 pprof
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	// a := make([]byte, 1000000)
	a := pool.Get().([]byte) // 用 sync pool 对象池优化下，gc 明显变少了很多，heap 分配也没那么大了
	defer pool.Put(a)
	for i := 0; i < 1000000; i++ {
		if len(a) > 1000000 { // pool 重用要控制下大小防止 pool 没被回收时无限增长
			break
		}
		a = append(a, byte(i))
	}
	w.Write(a)
}
