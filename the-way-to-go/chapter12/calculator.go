package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"./packages/stack"
)

func main() {
	inputBuffer := bufio.NewReader(os.Stdin)
	st := stack.NewStack()
	fmt.Println("Give two number, an operator (+, -, *, /), or q to stop:")
	for {
		input, err := inputBuffer.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}

		input = strings.TrimRight(input, "\r\n")
		switch {
		case input == "q":
			fmt.Println("Calculator stopped")
			return
		case input >= "0" && input <= "999999":
			i, _ := strconv.Atoi(input)
			st.Push(i)
		case input == "+":
			if st.Len() < 2 {
				fmt.Println("Not enough numbers")
				continue
			}
			i1, _ := st.Pop()
			i2, _ := st.Pop()
			st.Clear()
			fmt.Printf("The result of %d %s %d = %d\n", i2, input, i1, i2.(int) + i1.(int))
		case input == "-":
			if st.Len() < 2 {
				fmt.Println("Not enough numbers")
				continue
			}
			i1, _ := st.Pop()
			i2, _ := st.Pop()
			st.Clear()
			fmt.Printf("The result of %d %s %d = %d\n", i2, input, i1, i2.(int) - i1.(int))
		case input == "*":
			if st.Len() < 2 {
				fmt.Println("Not enough numbers")
				continue
			}
			i1, _ := st.Pop()
			i2, _ := st.Pop()
			st.Clear()
			fmt.Printf("The result of %d %s %d = %d\n", i2, input, i1, i2.(int) * i1.(int))
		case input == "/":
			if st.Len() < 2 {
				fmt.Println("Not enough numbers")
				continue
			}
			i1, _ := st.Pop()
			i2, _ := st.Pop()
			st.Clear()
			fmt.Printf("The result of %d %s %d = %d\n", i2, input, i1, i2.(int) / i1.(int))
		default:
			fmt.Println("No valid input")
		}
	}
}