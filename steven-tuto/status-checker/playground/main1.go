package main

import (
	"fmt"
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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		color.Red(link, "might be down!")
	} else {
		fmt.Println(link, "is up!")
	}
}
