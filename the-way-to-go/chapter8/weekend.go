package main

import "fmt"

func main() {
	weekend := map[string]string{
		"Monday": "星期一",
		"Tuesday": "星期二",
		"Wednesday": "星期三",
		"Thursday": "星期四",
		"Friday": "星期五",
		"Saturday": "星期六",
		"Sunday": "星期七",
	}

	for k, v := range weekend {
		fmt.Printf("%s : %s\n", k, v)
	}	

	value, ok := weekend["Thursday"]
	if ok {
		fmt.Println("存在 Thursday", value)
	} else {
		fmt.Println("不存在 Thursday")
	}

	value, ok = weekend["Hollyday"]
	if ok {
		fmt.Println("存在 Hollyday", value)
	} else {
		fmt.Println("不存在 Hollyday")
	}
}
