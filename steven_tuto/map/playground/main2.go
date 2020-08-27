package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors["white"] = "#fff"
	colors["red"] = "#ff0000"
	colors["black"] = "#000"

	//Delete key value of map
	delete(colors, "black")

	fmt.Println(colors)
}
