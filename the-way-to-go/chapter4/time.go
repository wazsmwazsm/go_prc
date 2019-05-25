package main
import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t) // 2018-03-28 02:22:15.077610044 +0000 UTC m=+0.001839496
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 28.03.2018
	
	t = time.Now().UTC()
	fmt.Println(t) // 2018-03-28 02:22:15.078882756 +0000 UTC
	fmt.Println(time.Now()) // 2018-03-28 02:22:15.079087395 +0000 UTC m=+0.003316856
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // 2018-04-04 02:22:15.078882756 +0000 UTC
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 28 Mar 18 02:22 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Mar 28 02:22:15 2018
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 28 Mar 2018 02:22
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// 2018-03-28 02:22:15.078882756 +0000 UTC => 20180328
}
