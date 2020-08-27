package main

import (
	"github.com/fatih/color"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://godlang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			go checkLink(link, c)
		}(l)
	}

	// time.Sleep(time.Second)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		print(link)
		color.Red(" might be down!")
		c <- link
		return
	}
	print(link)
	color.Green(" is up!")
	c <- link
}
