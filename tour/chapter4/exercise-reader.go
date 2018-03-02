package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// 实现 Reader 接口, 重写 Read 方法, 返回无限流 'A'
func (r MyReader) Read(b []byte) (int, error) {
    for i := range b {
        b[i] = 'A'
    }
  	return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
