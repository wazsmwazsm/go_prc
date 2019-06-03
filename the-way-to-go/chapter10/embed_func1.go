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
	log *Log // 组合 Log 结构
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"

	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}
