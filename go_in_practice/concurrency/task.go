package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func mainTask() {
	go echo(os.Stdin, os.Stdout)

	time.Sleep(30 * time.Second)
	fmt.Println("Oops! Time out. You tapped nothing after 30 seconds.")
	os.Exit(0)

}

func echo(in io.Reader, out io.Writer) {
	fmt.Print("Enter a input text: ")
	io.Copy(out, in)

}
