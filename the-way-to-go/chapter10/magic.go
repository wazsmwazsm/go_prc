package main

import (
	"fmt"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	// 虽然 MoreMagic 被 Voodoo 的实例调用，但是这里的接收者依然指的是 Base
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic() // voodoo magic
	v.MoreMagic() // base magic\n base magic
}
