package main
func main() {
	var p *int = nil
	// 对一个空指针的反向引用是不合法的
	*p = 0
}
