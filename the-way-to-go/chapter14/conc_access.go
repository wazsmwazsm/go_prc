package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend() // 扔出一个协程，监听放入协程的函数，执行函数
	return p
}

// 监听放入协程的函数，执行函数，函数序列化顺序执行，保证了安全
// 和加锁相似，并发执行 SetSalary、Salary 时，顺序执行（相当于依次加锁解锁）
func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// SetSalary 扔一个设置 salary 的闭包到通道，等待被 backend 消费
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// Salary 扔一个读取 salary 的闭包到通道，等待被 backend 消费
func (p *Person) Salary() float64 {
	chS := make(chan float64) // 通过通道和闭包通信
	p.chF <- func() { chS <- p.salary }
	return <-chS
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("New York", 5200)
	fmt.Println(bs)
	bs.SetSalary(7000)
	fmt.Println("Salary changed:")
	fmt.Println(bs)

	fmt.Println("SetSalary benchmark: ", testing.Benchmark(TestSetSalary).String())
	fmt.Println("Salary benchmark: ", testing.Benchmark(TestSalary).String())
}

func TestSetSalary(b *testing.B) {
	bs := NewPerson("New York", 5200)
	for i := 0; i < b.N; i++ {
		bs.SetSalary(float64(i))
	}
}

func TestSalary(b *testing.B) {
	bs := NewPerson("New York", 5200)
	for i := 0; i < b.N; i++ {
		bs.Salary()
	}
}
