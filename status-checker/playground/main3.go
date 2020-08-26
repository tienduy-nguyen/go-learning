package main

import (
	"github.com/fatih/color"
	"net/http"
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
	for {
		go checkLink(<-c, c)
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
