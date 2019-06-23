package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	input := "<Person type=\"gif\"><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	// Token() 方法返回输入流里的下一个 XML token
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
		
	}
}