package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Second * 1)

}
