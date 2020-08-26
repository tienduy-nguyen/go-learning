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

	for _, link := range links {
		go checkLink(link)
	}
	time.Sleep(time.Second)
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		print(link)
		color.Red(" might be down!")
	} else {
		print(link)
		color.Green(" is up!")
	}
}
