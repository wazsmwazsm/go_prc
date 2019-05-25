package main

func main() {
	for i:=0; i<3; i++ {
		for j:=0; j<10; j++ {
			if j>5 {
			    break   // 只跳出了最内层循环, 想要跳出外层用编号指定跳出位置
			}
			print(j)
		}
		print("  ")
	}
}
