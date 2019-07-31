package main

import "fmt"

type Log struct {
	msg string
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

type Customer struct {
	Name string
	Log  // 嵌入 Log 结构
}

func (c *Customer) DLog() string {
	return c.String()
}

func main() {
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c.DLog())
}
