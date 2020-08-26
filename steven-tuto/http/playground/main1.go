package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))
}
