package main

import (
	"fmt"
)

func main() {
	g1()
	g2()
}

func g1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func g2() {
	for i := 11; i <= 20; i++ {
		fmt.Println(i)
	}
}
