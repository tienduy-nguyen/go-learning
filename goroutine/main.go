package main

import (
	"fmt"
	"time"
)

func main() {
	go g1()
	go g2()
	time.Sleep(time.Second)
}

func g1() {
	for i := 1; i <= 10; i++ {
		go fmt.Println(i)
	}
}

func g2() {
	for i := 1; i <= 10; i++ {
		go fmt.Println(i)
	}
}
