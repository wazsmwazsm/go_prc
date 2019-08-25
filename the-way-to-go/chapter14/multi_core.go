package main

import (
	"runtime"
)

const NCPU = 4

func main() {
	runtime.GOMAXPROCS(NCPU)
	DoAll()
}

func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}

	for i := 0; i < NCPU; i++ {
		<-sem
	}
}

func DoPart(sem chan int) {
	sem <- 1
}
