package main

func main() {
	i:=0
	HERE:
		print(i)
		i++
		if i==5 {
			return
		}
		goto HERE // 倒序的 goto 是不被鼓励的
}
