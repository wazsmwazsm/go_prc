package main

func main() {
	for i := 0; i < 10; i++ {
		// 忽略剩余的循环体而直接进入下一次循环的过程
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
}
