package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}

func (this *Page) Save() error {
	return ioutil.WriteFile(this.Title, this.Body, 0666) // 不用 buf
}

func (this *Page) Load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title) // 不用 buf
	return
}

func main() {
	page := &Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."), // 强转类型
	}
	page.Save()

	// laod
	var new_page Page
	new_page.Load("Page.md")
	fmt.Println(string(new_page.Body))
}
