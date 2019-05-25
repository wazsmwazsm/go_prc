	package main

	import (
		"fmt"
		"sort"
	)

	var drinks = map[string]string{
		"sprite": "雪碧",
		"orangeJuice": "橘汁",
		"cola": "可乐",
		"tea": "茶",
	}

	func main() {
		for k, v := range drinks {
			fmt.Printf("%s %s\n", k, v)
		}

		keys := make([]string, len(drinks))
		i := 0
		for k, _ := range drinks {
			keys[i] = k
			i++
		}
		sort.Strings(keys)

		for _, v := range keys {
			fmt.Printf("%s %s\n", v, drinks[v])
		}

	}