package main

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

var testStr string
var initBufSize = 30000

// 使用 pool 来节省内存
var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, initBufSize))
	},
}

func main() {
	for i := 0; i < 10000; i++ {
		testStr += "a"
	}
	fmt.Println("bench", testing.Benchmark(bench).String(), testing.Benchmark(bench).MemString())
	fmt.Println("bench buf", testing.Benchmark(benchBuf).String(), testing.Benchmark(benchBuf).MemString())
	fmt.Println("bench go", testing.Benchmark(benchGo).String(), testing.Benchmark(benchGo).MemString())
	fmt.Println("bench buf go", testing.Benchmark(benchBufGo).String(), testing.Benchmark(benchBufGo).MemString())
}

func read() {
	buf := bytes.NewBufferString(testStr)
	buf.WriteString("aaa")
	buf.String()
	buf.Reset()
}

func readBuf() {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.WriteString("aaa")
	buf.String()
	buf.Reset()
}

func bench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read()
	}
}

func benchBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readBuf()
	}
}

func benchGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go read()
	}
}

func benchBufGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go readBuf()
	}
}
