package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

type Person struct {
	mu     *sync.Mutex
	Name   string
	salary float64
}

func NewPerson(name string, salary float64) *Person {

	p := &Person{new(sync.Mutex), name, salary}
	return p
}

func (p *Person) SetSalary(sal float64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.salary = sal
}

func (p *Person) Salary() float64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.salary
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
