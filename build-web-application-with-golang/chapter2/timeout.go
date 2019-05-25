package main

import "time"

func main() {
	c := make(chan int)	
	timeout := make(chan bool)

	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second): // After 方法返回一个 channel
				println("timeout")
				timeout <- true
				break
			}
		}
	}()
	if <- timeout {
		println("timeout quit")
	}
}
