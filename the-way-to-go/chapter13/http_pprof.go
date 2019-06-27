package main

import (
	"fmt"
	"math/rand"
    "time"
	"net/http"
	_ "net/http/pprof" // 调用 pprof 包中的 init
)

const (
    col = 10000
    row = 10000
)

func main() {
	http.HandleFunc("/loop", loop)
	http.ListenAndServe("0.0.0.0:6060", nil)
}

func loop(w http.ResponseWriter, r *http.Request) {
	x := [row][col]int{}
    s := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < row; i++{
        for j := 0; j < col; j++ {
            x[i][j] = s.Intn(100000)
        }
    }


    for i := 0; i < row; i++{
        tmp := 0
        for j := 0; j < col; j++ {
            tmp += x[i][j]
        }
	}
	
	fmt.Fprintf(w, "Successed! count: %v", len(x))
}
