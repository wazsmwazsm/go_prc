package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1e9) // 超时会关闭 done channel

	ch := make(chan string)

	go request(ctx, ch)

	fmt.Println(<-ch)
}

func request(ctx context.Context, ch chan string) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", "https://github.com/", strings.NewReader(""))
	if err != nil {
		log.Panic(err)
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done(): // 超时, panic
			log.Panic(ctx.Err())
		}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	ch <- string(body)
}
