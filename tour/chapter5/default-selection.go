package main

import (
    "fmt"
    "time"
)

func main()  {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)

    for {
        select {
            case <- tick:
                fmt.Println("tick.")
            case <- boom:
                // 到达定时
                fmt.Println("BOOM!")
                return
            default:
                // 没有到时, sleep 50 ms
          			fmt.Println("    .")
          			time.Sleep(50 * time.Millisecond)
        }
    }
}
